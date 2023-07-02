package data

import (
	"database/sql"
	"log"
	"time"
)

// Movie type whose fields describe the movie.
// Note that the Runtime type uses a custom Runtime type instead of int32. Furthermore, the omitempty
// directive on the Runtime type will still work on this: if the Runtime field has the underlying
// value 0, then it will be considered empty and omitted -- and the MarshalJSON() method won't
// be called.
type Email struct {
	ID        int64     `json:"id"` // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`  // Use the - directive to never export in JSON output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Movie release year0
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // The version number starts at 1 and is incremented each
	// time the movie information is updated.
}

// MovieModel struct wraps a sql.DB connection pool and allows us to work with Movie struct type
// and the movies table in our database.
type EmailModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
