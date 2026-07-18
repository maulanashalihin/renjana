-- +goose Up
CREATE TABLE renjana_survey_skm (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    age INTEGER NOT NULL,
    gender TEXT NOT NULL,
    education TEXT NOT NULL,
    occupation TEXT NOT NULL,
    year INTEGER NOT NULL DEFAULT 2026,
    q1 INTEGER NOT NULL CHECK(q1 >= 1 AND q1 <= 4),
    q2 INTEGER NOT NULL CHECK(q2 >= 1 AND q2 <= 4),
    q3 INTEGER NOT NULL CHECK(q3 >= 1 AND q3 <= 4),
    q4 INTEGER NOT NULL CHECK(q4 >= 1 AND q4 <= 4),
    q5 INTEGER NOT NULL CHECK(q5 >= 1 AND q5 <= 4),
    q6 INTEGER NOT NULL CHECK(q6 >= 1 AND q6 <= 4),
    q7 INTEGER NOT NULL CHECK(q7 >= 1 AND q7 <= 4),
    q8 INTEGER NOT NULL CHECK(q8 >= 1 AND q8 <= 4),
    q9 INTEGER NOT NULL CHECK(q9 >= 1 AND q9 <= 4),
    feedback TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS renjana_survey_skm;
