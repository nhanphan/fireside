package create

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Post
	using http://github.com/brianstarke/go-beget/generator/creator
*/

import (
	"github.com/brianstarke/go-beget/creator"
	"github.com/jmoiron/sqlx"

	types "github.com/nhanphan/fireside-server/models"
)

// PostCreator defines the interface for this creator
type PostCreator interface {
	Create(c types.Post) (int64, error)
}

// SQLPostCreator implements a SQL creator
type SQLPostCreator struct {
	db *sqlx.DB
}

// NewSQLPostCreator returns a SQL based PostCreator
func NewSQLPostCreator(db *sqlx.DB) PostCreator {
	return &SQLPostCreator{db: db}
}

// Create inserts a Post record in to the database
func (r *SQLPostCreator) Create(c types.Post) (int64, error) {
	columns := []string{
		"uuid",
		"creator_uuid",
		"type",
	}

	return creator.Insert(r.db.DB, "posts", columns, c.UUID, c.CreatorUUID, c.Type)
}
