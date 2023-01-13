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
	Service OpinionService
	Server  *http.Server
}

func NewHandler(service OpinionService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is alive !!")
	})
	// Aseguramos el handler PostOpinion con
	// middleware de autenticación con JWT
	h.Router.HandleFunc("/api/v1/opinion", JWTAuth(h.PostOpinion)).Methods("POST")

	h.Router.HandleFunc("/api/v1/opinion/{id}", h.GetOpinion).Methods("GET")
	h.Router.HandleFunc("/api/v1/opinion/{id}", h.UpdateOpinion).Methods("PUT")
	h.Router.HandleFunc("/api/v1/opinion/{id}", h.DeleteOpinion).Methods("DELETE")

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

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}
