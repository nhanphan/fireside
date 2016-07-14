package search

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Post
	using http://github.com/brianstarke/go-beget/generator/searcher
*/

import (
	"fmt"
	"strings"

	"github.com/brianstarke/go-beget/searcher"
	"github.com/jmoiron/sqlx"

	types "github.com/nhanphan/fireside-server/models"
)

// PostSearcher is the interface
type PostSearcher interface {
	Search(searchRequest PostSearchRequest) ([]types.Post, error)
	Count(searchRequest PostSearchRequest) (int32, error)
	GetByField(field PostField, value interface{}) (*types.Post, error)
}

// SQLPostSearcher implements a SQL based searcher
type SQLPostSearcher struct {
	db *sqlx.DB
}

// NewSQLPostSearcher returns a configured repo for you
func NewSQLPostSearcher(db *sqlx.DB) PostSearcher {
	return &SQLPostSearcher{db: db}
}

// Search converts a PostSearchRequest in to SQL and executes it
func (r *SQLPostSearcher) Search(searchRequest PostSearchRequest) ([]types.Post, error) {
	results := []types.Post{}

	sqlStr, values, err := searcher.GenerateSelectSQL(&searchRequest, false)

	if err != nil {
		return nil, fmt.Errorf("Error generating SQL for PostSearchRequest : %s", err.Error())
	}

	err = r.db.Select(&results, sqlStr, values.([]interface{})...)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return results, nil
		}
	}

	return results, err
}

// Count converts a PostSearchRequest in to SQL and executes it
func (r *SQLPostSearcher) Count(searchRequest PostSearchRequest) (int32, error) {
	var results []struct {
		Count int32 `db:"cnt"`
	}

	sqlStr, values, err := searcher.GenerateSelectSQL(&searchRequest, true)

	if err != nil {
		return 0, fmt.Errorf("Error generating SQL for PostSearchRequest : %s", err.Error())
	}

	err = r.db.Select(&results, sqlStr, values.([]interface{})...)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return 0, fmt.Errorf("No results returned")
		}
	}

	return results[0].Count, err
}

// GetByField is a convenience method to return just one result by matching on a single field
func (r *SQLPostSearcher) GetByField(field PostField, value interface{}) (*types.Post, error) {
	var searchRequest PostSearchRequest

	searchRequest.AddFilter(
		field,
		value,
		searcher.EQ,
		searcher.AND,
	)
	searchRequest.Limit = 1

	sqlStr, values, err := searcher.GenerateSelectSQL(&searchRequest, false)

	if err != nil {
		return nil, fmt.Errorf("Error generating GetByField SQL for Post [db field:%s]: %s", field.DbFieldName(), err.Error())
	}

	var result types.Post
	err = r.db.Get(&result, sqlStr, values.([]interface{})...)

	if err != nil {
		return nil, err
	}

	return &result, err
}
