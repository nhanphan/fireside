package post

import (
	// "bytes"
	// "fmt"

	"github.com/brianstarke/go-beget/searcher"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"

	"github.com/nhanphan/fireside-server/search"

	pb "github.com/nhanphan/fireside-server/proto"
	"golang.org/x/net/context"
)

// Server has method receivers to handle RPC requests to the
// service
type Server struct {
	db *sqlx.DB
}

// NewServer returns a handler
func NewServer(db *sqlx.DB) *Server {
	return &Server{
		db: db,
	}
}

func (h *Server) GetByUUID(ctx context.Context, r *pb.GetByUUIDRequest) (*pb.Post, error) {
	post, err := search.NewSQLPostSearcher(h.db).GetByField(search.PostUUID, r.Uuid)

	if err != nil {
		return nil, err
	}

	return post.ToProto(), nil
}

func (h *Server) Search(ctx context.Context, r *pb.SearchRequest) (*pb.Posts, error) {
	var sr search.PostSearchRequest

	if r.Filters != nil {
		for _, f := range r.Filters {
			var v interface{}

			switch f.Value.(type) {
			case *pb.Filter_StringVal:
				v = f.GetStringVal()
			case *pb.Filter_IntVal:
				v = f.GetIntVal()
			default:
				v = nil
			}

			sr.AddFilter(
				search.PostField(f.Field),
				v,
				searcher.FilterOperator(f.Operator.String()),
				searcher.FilterCondition(f.Condition.String()))
		}
	}

	if len(r.Fields) > 0 {
		for _, f := range r.Fields {
			sr.Fields = append(sr.Fields, search.PostField(f))
		}
	}

	sr.Limit = int(r.Limit)
	sr.Offset = int(r.Offset)

	if r.OrderBy != nil {
		sr.SetOrderBy(search.PostField(r.OrderBy.Field), r.OrderBy.IsDescending)
	}

	refs, err := search.NewSQLPostSearcher(h.db).Search(sr)

	if err != nil {
		glog.Errorf("Error searching Posts: %v", err)
		return nil, err
	}

	var pbRefs []*pb.Post

	for _, result := range refs {
		pbRefs = append(pbRefs, result.ToProto())
	}

	return &pb.Posts{Posts: pbRefs}, nil
}
