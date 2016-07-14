package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/nhanphan/fireside-server/post"
	pb "github.com/nhanphan/fireside-server/proto"
)

var (
	db     *sqlx.DB
	server *post.Server
)

func init() {
	flag.Parse()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"))

	glog.Infof("Connecting to DB [%s@%s]", os.Getenv("DB_HOST"), os.Getenv("DB_DBNAME"))

	var err error
	db, err = sqlx.Connect("postgres", connectionString)

	if err != nil {
		glog.Fatalf("Error connecting to database [%s]: %v", connectionString, err)
		panic(err)
	}

	server = post.NewServer(db)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx)

	if err != nil {
		return err
	}

	mux.Handle("/", gw)
	mux.HandleFunc("/swagger/", serveSwagger)

	http.ListenAndServe(os.Getenv("REST_PORT"), allowCORS(mux))

	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	lis, _ := net.Listen("tcp", os.Getenv("PORT"))

	s := grpc.NewServer()
	pb.RegisterPostSvcServer(s, server)

	go s.Serve(lis)

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("proto", p)
	http.ServeFile(w, r, p)
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterPostSvcHandlerFromEndpoint(ctx, mux, os.Getenv("Post_ENDPOINT"), dialOpts)

	if err != nil {
		return nil, err
	}

	return mux, nil
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	glog.Infof("preflight request for %s", r.URL.Path)
	return
}
