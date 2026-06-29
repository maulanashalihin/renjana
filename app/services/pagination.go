package services

import (
	"math"
	"strings"
)

// PaginationParams holds pagination and filter request.
type PaginationParams struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Search  string `json:"search"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"` // "asc" | "desc"
}

// PaginationResult is the generic shape for any paginated list endpoint.
type PaginationResult struct {
	Data        any   `json:"data"`
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	HasPrev     bool  `json:"has_prev"`
	HasNext     bool  `json:"has_next"`
}

// DefaultPage is the fallback when page param is invalid.
const DefaultPage = 1

// DefaultPerPage is the default items-per-page when not specified.
const DefaultPerPage = 20

// MaxPerPage is the hard cap (prevents memory blowup on huge pages).
const MaxPerPage = 100

// NormalizePagination normalizes page and perPage values with defaults & caps.
// Returns a tuple: (page, perPage, offset).
func NormalizePagination(page, perPage int) (int, int, int) {
	if page < 1 {
		page = DefaultPage
	}
	if perPage < 1 {
		perPage = DefaultPerPage
	}
	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}
	offset := (page - 1) * perPage
	return page, perPage, offset
}

// BuildPagination assembles the result envelope around a page of data.
func BuildPagination(data any, page, perPage int, totalItems int64) *PaginationResult {
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if totalPages < 1 {
		totalPages = 1
	}
	return &PaginationResult{
		Data:        data,
		CurrentPage: page,
		PerPage:     perPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		HasPrev:     page > 1,
		HasNext:     page < totalPages,
	}
}

// EscapeLike escapes SQL LIKE wildcards in user-provided search terms.
// Use with `LIKE '%' || @search || '%' ESCAPE '\'` in queries.
func EscapeLike(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		`%`, `\%`,
		`_`, `\_`,
	)
	return r.Replace(s)
}

// SortDirSafe returns "asc" or "desc" with safe fallback.
func SortDirSafe(dir string) string {
	lower := strings.ToLower(strings.TrimSpace(dir))
	if lower == "asc" || lower == "desc" {
		return lower
	}
	return "asc"
}
