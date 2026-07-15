#!/usr/bin/env python3
"""
Fetch all schools in Kabupaten Tanah Bumbu from Dapodik API (Kemdikbud).
Output: SQL INSERT statements for renjana_schools table.

Mapping Dapodik bentuk_pendidikan -> our level:
  - SD        -> SD
  - MI        -> MI
  - SMP       -> SMP
  - MTs       -> MTs
  - SMA       -> SMA
  - MA        -> MA
  - SMK       -> SMK

Usage:
  python3 scripts/fetch_schools.py              # print SQL to stdout
  python3 scripts/fetch_schools.py --update     # update migration file directly
"""

import sys
import requests
import json

KODE_KABUPATEN = "150800"  # Tanah Bumbu
SEMESTER_ID = "20242"       # 2024/2025 semester genap (2)

HEADERS = {
    "User-Agent": "Mozilla/5.0 (compatible; RenjanaBot/1.0; +https://renjana.maulanabuilds.com)",
    "Accept": "application/json",
}

# Map Dapodik bentuk_pendidikan to our level constants
LEVEL_MAP = {
    "SD": "SD",
    "MI": "MI",
    "SMP": "SMP",
    "MTs": "MTs",
    "SMA": "SMA",
    "MA": "MA",
    "SMK": "SMK",
    "SMK": "SMK",
}

# Map Dapodik status_sekolah to our status
STATUS_MAP = {
    "NEGERI": "Negeri",
    "SWASTA": "Swasta",
}


def fetch_json(url):
    res = requests.get(url, headers=HEADERS, timeout=30)
    res.raise_for_status()
    return res.json()


def clean_name(name):
    """Clean school name: strip, title-case properly."""
    name = name.strip()
    # Keep common acronyms uppercase
    for acronym in ["SD", "MI", "SMP", "MTs", "SMA", "MA", "SMK", "SMAN", "SMKN",
                    "SDN", "MIN", "MTsN", "SMPN"]:
        name = name.replace(acronym, acronym)
    # Fix "Negeri" / "Swasta" suffix that sometimes appears in names
    name = name.replace(" NEGERI", "").replace(" SWASTA", "")
    return name.strip()


def fetch_all_schools():
    print("📡 Mengambil daftar kecamatan...", file=sys.stderr)
    url_kec = (
        f"https://dapo.kemdikbud.go.id/rekap/progresSP"
        f"?id_level_wilayah=2&kode_wilayah={KODE_KABUPATEN}&semester_id={SEMESTER_ID}"
    )
    kecamatan_list = fetch_json(url_kec)

    all_schools = []

    for kec in kecamatan_list:
        kode_kec = kec["kode_wilayah"]
        nama_kec = kec["nama"].strip()
        print(f"  📍 {nama_kec}...", file=sys.stderr)

        url_sekolah = (
            f"https://dapo.kemdikbud.go.id/rekap/dataSekolah"
            f"?id_level_wilayah=3&kode_wilayah={kode_kec}&semester_id={SEMESTER_ID}"
        )
        try:
            sekolah_list = fetch_json(url_sekolah)
        except Exception as e:
            print(f"  ⚠️  Gagal: {e}", file=sys.stderr)
            continue

        for s in sekolah_list:
            nama = clean_name(s.get("nama", ""))
            jenjang_raw = s.get("bentuk_pendidikan", "").strip().upper()
            status_raw = s.get("status_sekolah", "").strip().upper()

            level = LEVEL_MAP.get(jenjang_raw)
            status = STATUS_MAP.get(status_raw)

            if not level or not status or not nama:
                print(f"  ⚠️  Skipped: {nama} (jenjang={jenjang_raw}, status={status_raw})", file=sys.stderr)
                continue

            all_schools.append({
                "name": nama,
                "level": level,
                "status": status,
                "kecamatan": nama_kec,
            })

    # Deduplicate by (name, kecamatan)
    seen = set()
    unique = []
    for s in all_schools:
        key = (s["name"], s["kecamatan"])
        if key not in seen:
            seen.add(key)
            unique.append(s)

    return sorted(unique, key=lambda s: (s["level"], s["name"]))


def generate_sql(schools):
    lines = []
    for s in schools:
        name_escaped = s["name"].replace("'", "''")
        lines.append(
            f"('{name_escaped}', '{s['level']}', '{s['status']}', '{s['kecamatan']}')"
        )
    return lines


def update_migration(sql_lines):
    path = "migrations/0018_create_schools_table.sql"
    with open(path, "r") as f:
        content = f.read()

    # Find the INSERT line and replace everything after it until the semicolon
    marker = "-- Seed data from existing schools.ts"
    insert_start = "-- +goose StatementEnd\n\n-- +goose Down"
    
    new_insert = marker + "\nINSERT INTO renjana_schools (name, level, status, kecamatan) VALUES\n" + ",\n".join(sql_lines) + ";\n\n"
    
    before = content.split(marker)[0]
    after = content.split(insert_start)[1] if insert_start in content else ""
    
    new_content = before + new_insert + "-- +goose StatementEnd\n\n-- +goose Down" + after

    with open(path, "w") as f:
        f.write(new_content)
    
    print(f"\n✅ Migration updated: {path}", file=sys.stderr)


if __name__ == "__main__":
    import os
    os.chdir(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

    print("🏫 Mengambil data sekolah dari Dapodik...", file=sys.stderr)
    schools = fetch_all_schools()
    
    print(f"\n✅ Total {len(schools)} sekolah unik ditemukan.", file=sys.stderr)
    
    sql_lines = generate_sql(schools)
    
    if "--update" in sys.argv:
        update_migration(sql_lines)
    else:
        print("INSERT INTO renjana_schools (name, level, status, kecamatan) VALUES")
        print(",\n".join(sql_lines) + ";")
