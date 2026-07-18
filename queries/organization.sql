-- name: GetOrganization :one
SELECT id, vision, mission, history, structure, contact_email, contact_phone,
       address, social_instagram, social_tiktok, social_youtube,
       social_instagram_url, social_instagram_name,
       social_tiktok_url, social_tiktok_name,
       social_youtube_url, social_youtube_name,
       updated_at
FROM renjana_organization
WHERE id = 1;

-- name: UpsertOrganization :execrows
INSERT INTO renjana_organization (id, vision, mission, history, structure,
    contact_email, contact_phone, address, social_instagram, social_tiktok, social_youtube,
    social_instagram_url, social_instagram_name,
    social_tiktok_url, social_tiktok_name,
    social_youtube_url, social_youtube_name,
    updated_at)
VALUES (1, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
    ?, ?,
    ?, ?,
    ?, ?,
    CURRENT_TIMESTAMP)
ON CONFLICT(id) DO UPDATE SET
    vision = excluded.vision,
    mission = excluded.mission,
    history = excluded.history,
    structure = excluded.structure,
    contact_email = excluded.contact_email,
    contact_phone = excluded.contact_phone,
    address = excluded.address,
    social_instagram = excluded.social_instagram,
    social_tiktok = excluded.social_tiktok,
    social_youtube = excluded.social_youtube,
    social_instagram_url = excluded.social_instagram_url,
    social_instagram_name = excluded.social_instagram_name,
    social_tiktok_url = excluded.social_tiktok_url,
    social_tiktok_name = excluded.social_tiktok_name,
    social_youtube_url = excluded.social_youtube_url,
    social_youtube_name = excluded.social_youtube_name,
    updated_at = CURRENT_TIMESTAMP;