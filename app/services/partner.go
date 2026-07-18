package services

import (
	"context"
	"errors"

	"github.com/maulanashalihin/laju-go/app/queries"
)

var (
	ErrPartnerNotFound = errors.New("partner not found")
	ErrPartnerNameReq  = errors.New("partner name is required")
)

type PartnerService struct {
	querier *queries.Querier
}

func NewPartnerService(querier *queries.Querier) *PartnerService {
	return &PartnerService{querier: querier}
}

type PartnerItem struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LogoURL    string `json:"logo_url"`
	WebsiteURL string `json:"website_url"`
	SortOrder  int64  `json:"sort_order"`
}

type CreatePartnerInput struct {
	Name       string `json:"name"`
	LogoURL    string `json:"logo_url"`
	WebsiteURL string `json:"website_url"`
}

type UpdatePartnerInput struct {
	Name       string `json:"name"`
	LogoURL    string `json:"logo_url"`
	WebsiteURL string `json:"website_url"`
}

func (s *PartnerService) List(ctx context.Context) ([]PartnerItem, error) {
	rows, err := s.querier.ListPartners(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]PartnerItem, len(rows))
	for i, r := range rows {
		items[i] = PartnerItem{
			ID:         r.ID,
			Name:       r.Name,
			LogoURL:    r.LogoUrl,
			WebsiteURL: r.WebsiteUrl,
			SortOrder:  r.SortOrder,
		}
	}
	return items, nil
}

func (s *PartnerService) Create(ctx context.Context, input CreatePartnerInput) error {
	if input.Name == "" {
		return ErrPartnerNameReq
	}

	// Get max sort_order for appending at the end
	existing, err := s.querier.ListPartners(ctx)
	if err != nil {
		return err
	}
	maxOrder := int64(0)
	for _, p := range existing {
		if p.SortOrder > maxOrder {
			maxOrder = p.SortOrder
		}
	}

	return s.querier.CreatePartner(ctx, queries.CreatePartnerParams{
		Name:       input.Name,
		LogoUrl:    input.LogoURL,
		WebsiteUrl: input.WebsiteURL,
		SortOrder:  maxOrder + 1,
	})
}

func (s *PartnerService) Update(ctx context.Context, id int64, input UpdatePartnerInput) error {
	if input.Name == "" {
		return ErrPartnerNameReq
	}

	// Check exists
	_, err := s.querier.GetPartnerByID(ctx, id)
	if err != nil {
		return ErrPartnerNotFound
	}

	return s.querier.UpdatePartner(ctx, queries.UpdatePartnerParams{
		Name:       input.Name,
		LogoUrl:    input.LogoURL,
		WebsiteUrl: input.WebsiteURL,
		SortOrder:  0,
		ID:         id,
	})
}

func (s *PartnerService) Delete(ctx context.Context, id int64) error {
	// Check exists
	_, err := s.querier.GetPartnerByID(ctx, id)
	if err != nil {
		return ErrPartnerNotFound
	}
	return s.querier.DeletePartner(ctx, id)
}
