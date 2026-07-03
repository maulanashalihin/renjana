package services

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupContactTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_contacts (id INTEGER PRIMARY KEY AUTOINCREMENT, district_id INTEGER REFERENCES renjana_districts(id) ON DELETE SET NULL, name TEXT NOT NULL, role TEXT NOT NULL DEFAULT 'Fasilitator', phone TEXT, email TEXT, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_contacts_district ON renjana_contacts(district_id);
		CREATE INDEX idx_renjana_contacts_active ON renjana_contacts(is_active);
	`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestContactServiceCreate(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	contact, err := svc.Create(context.Background(), CreateContactRequest{
		Name: "Test Contact", DistrictID: 1, Role: "Koordinator",
		Phone: "08123456789", Email: "test@example.com",
	})
	require.NoError(t, err)
	assert.NotZero(t, contact.ID)
	assert.Equal(t, "Simpang Empat", contact.DistrictName)
}

func TestContactServiceCreateValidation(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	tests := []struct {
		name string
		req  CreateContactRequest
	}{
		{"empty name", CreateContactRequest{DistrictID: 1}},
		{"zero district", CreateContactRequest{Name: "T"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Create(context.Background(), tt.req)
			assert.Error(t, err)
		})
	}
}

func TestContactServiceGet(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	created, err := svc.Create(context.Background(), CreateContactRequest{
		Name: "Get Test", DistrictID: 1,
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Get Test", got.Name)
	assert.Equal(t, "Fasilitator", got.Role) // default

	_, err = svc.Get(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrContactNotFound)
}

func TestContactServiceUpdate(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	created, err := svc.Create(context.Background(), CreateContactRequest{
		Name: "Original", DistrictID: 1,
	})
	require.NoError(t, err)

	err = svc.Update(context.Background(), created.ID, UpdateContactRequest{
		Name: "Updated", DistrictID: 2, Role: "Fasilitator", Phone: "080000000",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated", got.Name)
	assert.Equal(t, int64(2), got.DistrictID)
	assert.Equal(t, "Batulicin", got.DistrictName)

	// Not found
	err = svc.Update(context.Background(), 99999, UpdateContactRequest{
		Name: "Nope", DistrictID: 1,
	})
	assert.ErrorIs(t, err, ErrContactNotFound)
}

func TestContactServiceDelete(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	created, err := svc.Create(context.Background(), CreateContactRequest{
		Name: "To Delete", DistrictID: 1,
	})
	require.NoError(t, err)

	err = svc.Delete(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = svc.Get(context.Background(), created.ID)
	assert.ErrorIs(t, err, ErrContactNotFound)

	// Double delete
	err = svc.Delete(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrContactNotFound)
}

func TestContactServiceList(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	for i := 0; i < 4; i++ {
		did := int64(1)
		if i >= 2 {
			did = 2
		}
		_, err := svc.Create(context.Background(), CreateContactRequest{
			Name: "Contact", DistrictID: did,
		})
		require.NoError(t, err)
	}

	// All
	result, err := svc.List(context.Background(), "", 0, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 4, len(result.Data.([]ContactItem)))

	// By district
	result, err = svc.List(context.Background(), "", 2, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]ContactItem)))
}

func TestContactServiceListAll(t *testing.T) {
	q := setupContactTestDB(t)
	svc := NewContactService(q)

	_, err := svc.Create(context.Background(), CreateContactRequest{
		Name: "C1", DistrictID: 1,
	})
	require.NoError(t, err)

	items, err := svc.ListAll(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, len(items))
}
