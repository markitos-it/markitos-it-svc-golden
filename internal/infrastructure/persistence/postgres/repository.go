package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"markitos-it-svc-goldens/internal/domain"
	"time"

	"github.com/lib/pq"
)

type GoldenRepository struct {
	db *sql.DB
}

func NewGoldenRepository(db *sql.DB) *GoldenRepository {
	return &GoldenRepository{db: db}
}

func (r *GoldenRepository) InitSchema(ctx context.Context) error {
	schema := `
	CREATE TABLE IF NOT EXISTS goldens (
		id VARCHAR(255) PRIMARY KEY,
		title VARCHAR(500) NOT NULL,
		description TEXT,
		category VARCHAR(100),
		tags TEXT[],
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		content_b64 TEXT NOT NULL,
		cover_image VARCHAR(1000) NOT NULL
	);

	CREATE INDEX IF NOT EXISTS idx_goldens_category ON goldens(category);
	CREATE INDEX IF NOT EXISTS idx_goldens_updated_at ON goldens(updated_at DESC);
	`

	_, err := r.db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	return nil
}

func (r *GoldenRepository) SeedData(ctx context.Context) error {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM goldens").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing data: %w", err)
	}

	if count > 0 {
		return nil
	}

	docs := []domain.Golden{
		{
			ID:          "getting-started-keptn",
			Title:       "Getting Started with Keptn",
			Description: "A comprehensive guide to get started with Keptn for automated deployment and operations",
			Category:    "DevOps",
			Tags:        []string{"keptn", "ci-cd", "automation", "kubernetes"},
			UpdatedAt:   time.Now(),
			ContentB64:  "IyBHZXR0aW5nIFN0YXJ0ZWQgd2l0aCBLZXB0bg==",
			CoverImage:  "https://images.unsplash.com/photo-1667372393119-3d4c48d07fc9",
		},
		{
			ID:          "youtube-api-integration",
			Title:       "YouTube Data API v3 Integration",
			Description: "Complete guide to integrate YouTube Data API with practical examples",
			Category:    "APIs",
			Tags:        []string{"youtube", "api", "rest", "video"},
			UpdatedAt:   time.Now(),
			ContentB64:  "IyBZb3VUdWJlIERhdGEgQVBJIHYzIEludGVncmF0aW9u",
			CoverImage:  "https://images.unsplash.com/photo-1611162616475-46b635cb6868",
		},
	}

	for _, doc := range docs {
		err := r.Create(ctx, &doc)
		if err != nil {
			return fmt.Errorf("failed to seed golden %s: %w", doc.ID, err)
		}
	}

	return nil
}

func (r *GoldenRepository) GetAll(ctx context.Context) ([]domain.Golden, error) {
	query := `
		SELECT id, title, description, category, tags, updated_at, content_b64, cover_image
		FROM goldens
		ORDER BY updated_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query goldens: %w", err)
	}
	defer rows.Close()

	var docs []domain.Golden
	for rows.Next() {
		var doc domain.Golden
		var tags pq.StringArray

		err := rows.Scan(
			&doc.ID,
			&doc.Title,
			&doc.Description,
			&doc.Category,
			&tags,
			&doc.UpdatedAt,
			&doc.ContentB64,
			&doc.CoverImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan golden: %w", err)
		}

		doc.Tags = []string(tags)
		docs = append(docs, doc)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating goldens: %w", err)
	}

	return docs, nil
}

func (r *GoldenRepository) GetByID(ctx context.Context, id string) (*domain.Golden, error) {
	query := `
		SELECT id, title, description, category, tags, updated_at, content_b64, cover_image
		FROM goldens
		WHERE id = $1
	`

	var doc domain.Golden
	var tags pq.StringArray

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&doc.ID,
		&doc.Title,
		&doc.Description,
		&doc.Category,
		&tags,
		&doc.UpdatedAt,
		&doc.ContentB64,
		&doc.CoverImage,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("golden not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query golden: %w", err)
	}

	doc.Tags = []string(tags)
	return &doc, nil
}

func (r *GoldenRepository) Create(ctx context.Context, doc *domain.Golden) error {
	query := `
		INSERT INTO goldens (id, title, description, category, tags, updated_at, content_b64, cover_image)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		doc.ID,
		doc.Title,
		doc.Description,
		doc.Category,
		pq.Array(doc.Tags),
		doc.UpdatedAt,
		doc.ContentB64,
		doc.CoverImage,
	)

	if err != nil {
		return fmt.Errorf("failed to create golden: %w", err)
	}

	return nil
}

func (r *GoldenRepository) Update(ctx context.Context, doc *domain.Golden) error {
	query := `
		UPDATE goldens
		SET title = $2, description = $3, category = $4, tags = $5, updated_at = $6, content_b64 = $7, cover_image = $8
		WHERE id = $1
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		doc.ID,
		doc.Title,
		doc.Description,
		doc.Category,
		pq.Array(doc.Tags),
		doc.UpdatedAt,
		doc.ContentB64,
		doc.CoverImage,
	)

	if err != nil {
		return fmt.Errorf("failed to update golden: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("golden not found: %s", doc.ID)
	}

	return nil
}

func (r *GoldenRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM goldens WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete golden: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("golden not found: %s", id)
	}

	return nil
}
