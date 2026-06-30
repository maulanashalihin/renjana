package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePagination(t *testing.T) {
	tests := []struct {
		name             string
		page, perPage    int
		wantPage, offset int
	}{
		{"defaults", 0, 0, 1, 0},
		{"valid", 2, 10, 2, 10},
		{"capped perPage", 1, 200, 1, 0},
		{"negative page", -1, 10, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, perPage, offset := NormalizePagination(tt.page, tt.perPage)
			assert.Equal(t, tt.wantPage, page)
			if tt.perPage > 0 && tt.perPage <= MaxPerPage {
				assert.Equal(t, tt.perPage, perPage)
			} else if tt.perPage > MaxPerPage {
				assert.Equal(t, MaxPerPage, perPage)
			} else {
				assert.Equal(t, DefaultPerPage, perPage)
			}
			assert.Equal(t, tt.offset, offset)
		})
	}
}

func TestBuildPagination(t *testing.T) {
	data := []string{"a", "b", "c"}
	result := BuildPagination(data, 1, 10, 25)

	assert.Equal(t, data, result.Data)
	assert.Equal(t, 1, result.CurrentPage)
	assert.Equal(t, 10, result.PerPage)
	assert.Equal(t, int64(25), result.TotalItems)
	assert.Equal(t, 3, result.TotalPages)
	assert.False(t, result.HasPrev)
	assert.True(t, result.HasNext)

	// Last page
	last := BuildPagination(data, 3, 10, 25)
	assert.True(t, last.HasPrev)
	assert.False(t, last.HasNext)
}

func TestBuildPaginationZeroItems(t *testing.T) {
	result := BuildPagination([]string{}, 1, 10, 0)
	assert.Equal(t, 1, result.TotalPages)
	assert.False(t, result.HasPrev)
	assert.False(t, result.HasNext)
}

func TestEscapeLike(t *testing.T) {
	assert.Equal(t, `hello`, EscapeLike("hello"))
	assert.Equal(t, `100\%`, EscapeLike("100%"))
	assert.Equal(t, `test\_name`, EscapeLike("test_name"))
	assert.Equal(t, `foo\\bar`, EscapeLike(`foo\bar`))
}

func TestSortDirSafe(t *testing.T) {
	assert.Equal(t, "asc", SortDirSafe("asc"))
	assert.Equal(t, "desc", SortDirSafe("desc"))
	assert.Equal(t, "asc", SortDirSafe("ASC"))
	assert.Equal(t, "asc", SortDirSafe("invalid"))
	assert.Equal(t, "asc", SortDirSafe(""))
}
