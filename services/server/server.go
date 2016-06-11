package server

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/evanh/fundmyworld/db"
	"github.com/evanh/fundmyworld/handlers"
	"github.com/evanh/fundmyworld/repositories"
	"github.com/evanh/fundmyworld/services"
)

type Server struct {
	HTTPServer  *http.Server
	FundService *services.FundService
}

var background = context.Background()

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateAndInitializeServer() *Server {
	db, err := db.NewDB()
	PanicOnErr(err)

	fundRepo, err := repositories.NewFundRepository(db)
	PanicOnErr(err)

	fundService, err := services.NewFundService(fundRepo)
	PanicOnErr(err)

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		HTTPServer:  server,
		FundService: fundService,
	}
}

func (s *Server) AddHandlers() {
	http.HandleFunc("/", func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(background, "test", "test")
			handlers.Root(ctx, w, r, s.FundService)
		}
	}())
}
