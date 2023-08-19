package models

import (
	"database/sql"
	"time"

	"github.com/farmani/sharebuy/pkg/logger"
)

type Email struct {
	ID        int64     `json:"id"` // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`  // Use the - directive to never export in JSON output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Movie release year0
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // The version number starts at 1 and is incremented each
	// time the movie information is updated.
}

type EmailModel struct {
	DB     *sql.DB
	Logger *logger.Logger
}
