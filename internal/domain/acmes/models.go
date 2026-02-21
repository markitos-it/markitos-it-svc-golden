package acmes

import (
	"time"
)

type Acme struct {
	ID          string
	Title       string
	Description string
	Category    string
	Tags        []string
	UpdatedAt   time.Time
	ContentB64  string
	CoverImage  string
}
