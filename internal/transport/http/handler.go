package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Router:  mux.NewRouter(),
		Service: service,
	}
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)
	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}

	fmt.Println("server created")
	return h
}
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !!!")
	})

	h.Router.HandleFunc("/api/v1/comments/{id}", JWTAuth(h.GetComment)).Methods("GET")

}

func (h *Handler) Serve() error {

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down")
	return nil
}
