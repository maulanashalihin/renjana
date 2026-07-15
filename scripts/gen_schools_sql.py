#!/usr/bin/env python3
"""
Generate SQL INSERT for renjana_schools from datapendidikan.com JSON.
Usage: python3 scripts/gen_schools_sql.py > insert.sql
"""
import json, re, sys, urllib.request

URL = "https://datapendidikan.com/json/kabupaten/151100.json"

# Map nama kecamatan to our standard
KECAMATAN_MAP = {
    "Kec. Kusan Hilir": "Kusan Hilir",
    "Kec. Sungai Loban": "Sungai Loban",
    "Kec. Satui": "Satui",
    "Kec. Kusan Hulu": "Kusan Hulu",
    "Kec. Batu Licin": "Batu Licin",
    "Kec. Simpang Empat": "Simpang Empat",
    "Kec. Karang Bintang": "Karang Bintang",
    "Kec. Mantewe": "Mantewe",
    "Kec. Angsana": "Angsana",
    "Kec. Kuranji": "Kuranji",
    "Kec. Teluk Kepayang": "Teluk Kepayang",
    "Kec. Kusan Tengah": "Kusan Tengah",
}

def detect_level(nama, nama_kec):
    up = nama.upper()
    # Must check specific prefixes in order
    if up.startswith("SD ") or up.startswith("SDN ") or up.startswith("SDIT ") or up.startswith("SDTQ ") or up.startswith("SD "):
        return "SD"
    if up.startswith("MI ") or up.startswith("MIN ") or up.startswith("MIS "):
        return "MI"
    if up.startswith("SMP ") or up.startswith("SMPN ") or "SMP " in up[:10] or up.startswith("SMP IT") or up.startswith("SMP ISLAM"):
        return "SMP"
    if up.startswith("MTs") or up.startswith("MTSS ") or up.startswith("MTsN "):
        return "MTs"
    if up.startswith("SMA ") or up.startswith("SMAN ") or up.startswith("SMAIT ") or up.startswith("SMAS "):
        return "SMA"
    if up.startswith("MA ") or up.startswith("MAN ") or up.startswith("MAS "):
        return "MA"
    if up.startswith("SMK ") or up.startswith("SMKN ") or up.startswith("SMKS "):
        return "SMK"
    return None

def main():
    print("Fetching data...", file=sys.stderr)
    req = urllib.request.Request(URL, headers={"User-Agent": "Mozilla/5.0"})
    data = json.loads(urllib.request.urlopen(req, timeout=30).read())

    schools = []
    skipped = []
    for s in data:
        nama = s["nama"].strip()
        nama_kec = s.get("nama_kec", "").strip()
        status_raw = s.get("status", "").strip().upper()
        
        kec = KECAMATAN_MAP.get(nama_kec, nama_kec.replace("Kec. ", ""))
        status = "Negeri" if status_raw == "NEGERI" else "Swasta" if status_raw == "SWASTA" else status_raw
        level = detect_level(nama, nama_kec)
        
        if not level:
            skipped.append(nama)
            continue
        
        # Deduplicate by (name, kecamatan)
        schools.append((nama, level, status, kec))
    
    # Deduplicate preserving order
    seen = set()
    unique = []
    for s in schools:
        key = (s[0].upper(), s[3])
        if key not in seen:
            seen.add(key)
            unique.append(s)
    
    unique.sort(key=lambda x: (x[1], x[0]))
    
    print(f"Total: {len(unique)} schools (skipped {len(skipped)})", file=sys.stderr)
    if skipped:
        print(f"Skipped: {skipped}", file=sys.stderr)
    
    print("INSERT INTO renjana_schools (name, level, status, kecamatan) VALUES")
    rows = []
    for nama, level, status, kec in unique:
        escaped = nama.replace("'", "''")
        rows.append(f"    ('{escaped}', '{level}', '{status}', '{kec}')")
    print(",\n".join(rows) + ";")

if __name__ == "__main__":
    main()
