<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import type { Map as LeafletMap, GeoJSON as LeafletGeoJSON } from "leaflet";
    import * as L from "leaflet";

    interface DistrictFeature {
        type: "Feature";
        properties: {
            kode: string;
            nama: string;
            centroid_lat: number;
            centroid_lng: number;
        };
        geometry: {
            type: "MultiPolygon";
            coordinates: number[][][][];
        };
    }

    interface DistrictGeoJSON {
        type: "FeatureCollection";
        features: DistrictFeature[];
    }

    interface Props {
        geojsonUrl?: string;
        volunteerData?: Record<string, number>; // district name → volunteer count
        onDistrictClick?: (name: string, count: number) => void;
    }

    let {
        geojsonUrl = "/dist/tanahbumbu-kecamatan.geojson",
        volunteerData = {},
        onDistrictClick = () => {},
    }: Props = $props();

    let mapContainer: HTMLDivElement;
    let map: LeafletMap | null = null;
    let geoLayer: LeafletGeoJSON | null = null;
    let error: string | null = $state(null);

    // Linear interpolation between two colors
    function interpolateColor(value: number, min: number, max: number): string {
        // 0 → very light orange/yellow, 1 → RENJANA orange (#f97316)
        const ratio = max === min ? 0.5 : Math.max(0, Math.min(1, (value - min) / (max - min)));
        // From light cream (255, 245, 230) to RENJANA orange (249, 115, 22)
        const r = Math.round(255 + (249 - 255) * ratio);
        const g = Math.round(245 + (115 - 245) * ratio);
        const b = Math.round(230 + (22 - 230) * ratio);
        return `rgb(${r}, ${g}, ${b})`;
    }

    function getColor(name: string): string {
        const counts = Object.values(volunteerData);
        const min = Math.min(...counts, 0);
        const max = Math.max(...counts, 1);
        return interpolateColor(volunteerData[name] ?? 0, min, max);
    }

    onMount(async () => {
        try {
            // Dynamic import to avoid SSR issues with leaflet (uses window/document)
            const L = await import("leaflet");

            // Fix marker icon paths (Leaflet default uses bundler-incompatible paths)
            const iconUrl = "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png";
            const iconRetinaUrl = "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png";
            const shadowUrl = "https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png";
            // @ts-expect-error - internal prop
            delete L.Icon.Default.prototype._getIconUrl;
            L.Icon.Default.mergeOptions({ iconUrl, iconRetinaUrl, shadowUrl });

            // Initialize map langsung menampilkan Tanah Bumbu tanpa loading overlay.
            // maxBounds menjaga viewport tetap di area Tanah Bumbu.
            map = L.map(mapContainer, {
                zoomControl: true,
                scrollWheelZoom: true,
                attributionControl: true,
                maxBounds: L.latLngBounds([-3.8, 115.2], [-3.0, 116.2]),
                maxBoundsViscosity: 1.0,
            }).setView([-3.37, 115.95], 10);

            // OSM base tiles
            L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
                maxZoom: 19,
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> | Kemendagri 2024',
            }).addTo(map);

            // Fetch GeoJSON
            const resp = await fetch(geojsonUrl);
            if (!resp.ok) throw new Error(`Failed to load GeoJSON: ${resp.status}`);
            const geo: DistrictGeoJSON = await resp.json();

            // Add choropleth layer
            geoLayer = L.geoJSON(geo as unknown as GeoJSON.FeatureCollection, {
                style: (feature) => {
                    const nama = feature?.properties?.nama ?? "";
                    return {
                        fillColor: getColor(nama),
                        weight: 1.5,
                        opacity: 1,
                        color: "#9a3412", // darker orange border
                        fillOpacity: 0.65,
                    };
                },
                onEachFeature: (feature, layer) => {
                    const nama = feature.properties?.nama ?? "";
                    const count = volunteerData[nama] ?? 0;
                    // Tooltip on hover
                    layer.bindTooltip(`<strong>${nama}</strong><br/>${count} volunteer`, {
                        sticky: true,
                        className: "peta-tooltip",
                        direction: "top",
                    });

                    // Highlight on hover — layer is a Polygon path
                    const path = layer as unknown as L.Polygon;
                    layer.on({
                        mouseover: () => {
                            path.setStyle({ fillOpacity: 0.85, weight: 2.5 });
                        },
                        mouseout: () => {
                            geoLayer?.resetStyle(path);
                        },
                        click: () => {
                            if (path.getBounds) {
                                map?.fitBounds(path.getBounds(), { padding: [40, 40], maxZoom: 11 });
                            }
                            onDistrictClick(nama, count);
                        },
                    });
                },
            }).addTo(map);

            // Add centroid markers sized by volunteer count
            geo.features.forEach((f) => {
                const nama = f.properties.nama;
                const count = volunteerData[nama] ?? 0;
                const radius = Math.max(6, Math.min(20, 6 + count * 0.07));
                L.circleMarker([f.properties.centroid_lat, f.properties.centroid_lng], {
                    radius,
                    fillColor: "#ea580c",
                    color: "#fff",
                    weight: 2,
                    opacity: 1,
                    fillOpacity: 0.95,
                })
                    .bindTooltip(`<strong>${nama}</strong><br/>${count} volunteer`, { permanent: false, direction: "top" })
                    .addTo(map!);
            });

            // Fit bounds to polygons
            if (geoLayer.getBounds().isValid()) {
                map.fitBounds(geoLayer.getBounds(), { padding: [20, 20] });
            }

        } catch (err) {
            console.error("Failed to initialize map:", err);
            error = err instanceof Error ? err.message : "Failed to load map";
        }
    });

    onDestroy(() => {
        map?.remove();
        map = null;
    });

    // Reactive: when volunteerData changes, restyle polygons
    $effect(() => {
        void volunteerData; // track dep
        if (!geoLayer) return;
        const layer = geoLayer as unknown as { eachLayer: (cb: (l: unknown) => void) => void };
        layer.eachLayer((l) => {
            const featLayer = l as { feature?: { properties?: { nama?: string } }; setStyle: (s: Record<string, unknown>) => void };
            const nama = featLayer.feature?.properties?.nama;
            if (nama) {
                featLayer.setStyle({ fillColor: getColor(nama) });
            }
        });
    });
</script>

<div class="relative w-full h-full min-h-[520px] rounded-xl overflow-hidden bg-neutral-100 dark:bg-neutral-800">
    <div bind:this={mapContainer} class="absolute inset-0" data-testid="peta-map"></div>

    {#if error}
        <div class="absolute inset-0 flex items-center justify-center bg-red-50 dark:bg-red-900/20 p-6">
            <div class="text-center">
                <p class="text-red-600 dark:text-red-400 font-medium">Gagal memuat peta</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400 mt-1">{error}</p>
            </div>
        </div>
    {/if}

    <!-- Color legend (bottom-left) -->
    {#if !error}
        <div class="absolute bottom-4 left-4 z-[400] bg-white/95 dark:bg-neutral-900/95 backdrop-blur rounded-lg p-3 shadow-lg border border-neutral-200 dark:border-neutral-800 text-xs">
            <p class="font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">Jumlah Volunteer</p>
            <div class="flex items-center gap-2">
                <span class="text-neutral-500">Sedikit</span>
                <div class="h-2 w-24 rounded-full" style="background: linear-gradient(to right, rgb(255,245,230), rgb(249,115,22));"></div>
                <span class="text-neutral-500">Banyak</span>
            </div>
        </div>
    {/if}
</div>

<style>
    :global(.peta-tooltip) {
        font-family: system-ui, sans-serif;
        font-size: 12px;
        padding: 4px 8px;
        background: rgba(15, 23, 42, 0.95);
        color: white;
        border-radius: 6px;
        border: none;
    }
    :global(.peta-tooltip::before) {
        display: none;
    }
</style>
