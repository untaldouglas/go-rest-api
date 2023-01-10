package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type OpinionService interface{}

type Handler struct {
	Router  *mux.Router
	Service OpinionService
	Server  *http.Server
}

func NewHandler(service OpinionService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
 
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello mundo")
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
