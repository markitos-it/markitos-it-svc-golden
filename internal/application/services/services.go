package services

import (
	"context"
	"markitos-it-svc-goldens/internal/domain"
)

type GoldenService struct {
	repo domain.Repository
}

func NewGoldenService(repo domain.Repository) *GoldenService {
	return &GoldenService{
		repo: repo,
	}
}

func (s *GoldenService) GetAllGoldens(ctx context.Context) ([]domain.Golden, error) {
	return s.repo.GetAll(ctx)
}

func (s *GoldenService) GetGoldenByID(ctx context.Context, id string) (*domain.Golden, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *GoldenService) CreateGolden(ctx context.Context, doc *domain.Golden) error {
	return s.repo.Create(ctx, doc)
}

func (s *GoldenService) UpdateGolden(ctx context.Context, doc *domain.Golden) error {
	return s.repo.Update(ctx, doc)
}

func (s *GoldenService) DeleteGolden(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
