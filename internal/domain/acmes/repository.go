package acmes

import (
	"context"
)

// Repository define la interfaz para el repositorio de acmeos
type Repository interface {
	// GetAll retorna todos los acmeos
	GetAll(ctx context.Context) ([]Acme, error)

	// GetByID retorna un acmeo por su ID
	GetByID(ctx context.Context, id string) (*Acme, error)

	// Create crea un nuevo acmeo
	Create(ctx context.Context, doc *Acme) error

	// Update actualiza un acmeo existente
	Update(ctx context.Context, doc *Acme) error

	// Delete elimina un acmeo por su ID
	Delete(ctx context.Context, id string) error
}
