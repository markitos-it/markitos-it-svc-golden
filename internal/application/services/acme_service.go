package services

import (
	"context"

	"markitos-it-svc-acmes/internal/domain/acmes"
)

type AcmeService struct {
	repo acmes.Repository
}

func NewAcmeService(repo acmes.Repository) *AcmeService {
	return &AcmeService{
		repo: repo,
	}
}

func (s *AcmeService) GetAllAcmes(ctx context.Context) ([]acmes.Acme, error) {
	return s.repo.GetAll(ctx)
}

func (s *AcmeService) GetAcmeByID(ctx context.Context, id string) (*acmes.Acme, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AcmeService) CreateAcme(ctx context.Context, doc *acmes.Acme) error {
	return s.repo.Create(ctx, doc)
}

func (s *AcmeService) UpdateAcme(ctx context.Context, doc *acmes.Acme) error {
	return s.repo.Update(ctx, doc)
}

func (s *AcmeService) DeleteAcme(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
