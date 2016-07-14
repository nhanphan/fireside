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
)

// PostField is a field within the Post struct
// that is able to be filtered on, sorted on, or returned.
type PostField int

// Enum'd for helpfulness
const (
	PostID PostField = iota
	PostUUID
	PostCreatorUUID
	PostType
)

// MarshalText implements https://golang.org/pkg/encoding/#TextMarshaler
func (s PostField) MarshalText() ([]byte, error) {
	var data string

	switch s {
	case PostID:
		data = "id"
	case PostUUID:
		data = "uuid"
	case PostCreatorUUID:
		data = "creatorUUID"
	case PostType:
		data = "type"

	default:
		return nil, fmt.Errorf("Unable to marshal `%v` in to bytes", s)
	}
	return []byte(data), nil
}

// UnmarshalText implements https://golang.org/pkg/encoding/#TextUnmarshaler
func (s *PostField) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "id":
		*s = PostID
	case "uuid":
		*s = PostUUID
	case "creatorUUID":
		*s = PostCreatorUUID
	case "type":
		*s = PostType

	default:
		return fmt.Errorf("Unable to marshal '%s' into type PostField", str)
	}
	return nil
}

// DbFieldName returns the name of the field to use in the SQL query
func (s PostField) DbFieldName() string {
	switch s {
	case PostID:
		return "id"
	case PostUUID:
		return "uuid"
	case PostCreatorUUID:
		return "creator_uuid"
	case PostType:
		return "type"

	}
	return ""
}

/*
PostSearchRequest defines a set of parameters for
searching for Post.  It can be serialized and passed
between services as JSON, or used to generate a SQL statement.
*/
type PostSearchRequest struct {
	searcher.SearchRequestFields
	Filters []PostSearchFilter `json:"filters"`
	OrderBy PostOrderBy        `json:"orderBy"`
	Fields  []PostField        `json:"fields"`
}

/*
PostSearchFilter is a filter specific to Post
*/
type PostSearchFilter struct {
	Field     PostField                `json:"field"`
	Value     interface{}              `json:"value"`
	Operator  searcher.FilterOperator  `json:"operator"`
	Condition searcher.FilterCondition `json:"condition"`
}

/*
PostOrderBy is a sort directive that is specific to Post
*/
type PostOrderBy struct {
	Field      PostField `json:"field"`
	Descending bool      `json:"desc"`
}

// Implement searcher.SearchRequest interface

// GetTableName returns the database table name
func (sr *PostSearchRequest) GetTableName() string {
	return "posts"
}

// GetFilters returns the SQL WHERE clauses
func (sr *PostSearchRequest) GetFilters() []searcher.Filter {
	filters := []searcher.Filter{}

	for _, f := range sr.Filters {
		filter := searcher.Filter{
			Field:     f.Field,
			Value:     f.Value,
			Operator:  f.Operator,
			Condition: f.Condition,
		}
		filters = append(filters, filter)
	}

	return filters
}

// GetOrderBy returns the SQL ORDER BY clauses
func (sr *PostSearchRequest) GetOrderBy() searcher.OrderBy {
	return searcher.OrderBy{
		Field:      sr.OrderBy.Field,
		Descending: sr.OrderBy.Descending,
	}
}

// GetLimit returns the SQL LIMIT clause
func (sr *PostSearchRequest) GetLimit() int {
	return sr.Limit
}

// GetOffset returns the SQL OFFSET clause
func (sr *PostSearchRequest) GetOffset() int {
	return sr.Offset
}

// AddFilter adds a WHERE clause
func (sr *PostSearchRequest) AddFilter(field PostField, value interface{}, operator searcher.FilterOperator, condition searcher.FilterCondition) {
	f := PostSearchFilter{
		Field:     field,
		Value:     value,
		Operator:  operator,
		Condition: condition,
	}
	sr.Filters = append(sr.Filters, f)
}

// SetOrderBy sets the ORDER BY clause
func (sr *PostSearchRequest) SetOrderBy(field PostField, isDescending bool) {
	sr.OrderBy = PostOrderBy{
		Field:      field,
		Descending: isDescending,
	}
}

// GetFields returns the SQL SELECT fields
func (sr *PostSearchRequest) GetFields() []string {
	fields := []string{}

	for _, f := range sr.Fields {
		fields = append(fields, f.DbFieldName())
	}

	return fields
}
