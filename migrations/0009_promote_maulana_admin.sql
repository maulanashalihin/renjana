-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Set existing users as admin (preserve their original password)
-- ============================================================================

-- Jadikan maulanashalihin@gmail.com sebagai admin (jika sudah ada)
UPDATE users
SET role = 'admin', is_active = 1, updated_at = CURRENT_TIMESTAMP
WHERE email = 'maulanashalihin@gmail.com';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
UPDATE users
SET role = 'relawan', updated_at = CURRENT_TIMESTAMP
WHERE email = 'maulanashalihin@gmail.com';
-- +goose StatementEnd