package domain

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Golden, error)
	GetByID(ctx context.Context, id string) (*Golden, error)
	Create(ctx context.Context, doc *Golden) error
	Update(ctx context.Context, doc *Golden) error
	Delete(ctx context.Context, id string) error
}
