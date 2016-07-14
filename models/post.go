package models

import (
	// "database/sql"
	"github.com/nhanphan/fireside-server/proto"
)

//go:generate searcher -struct=Post -table=posts -impls=sql
//go:generate creator -struct=Post -table=posts -impls=sql

type Post struct {
	ID          int64  `beget:"search" json:"id" db:"id"`
	UUID        string `beget:"search,create" json:"uuid" db:"uuid"`
	CreatorUUID string `beget:"search,create" json:"creatorUUID" db:"creator_uuid"`
	Type        string `beget:"search,create" json:"type" db:"type"`
}

func (p *Post) ToProto() *post.Post {
	post := post.Post{
		Uuid:        p.UUID,
		CreatorUuid: p.CreatorUUID,
		Type:        p.Type,
	}

	return &post
}
