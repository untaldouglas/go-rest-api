package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/untaldouglas/go-rest-api/internal/opinion"
)

type OpinionService interface {
	PostOpinion(context.Context, opinion.Opinion) (opinion.Opinion, error)
	GetOpinion(ctx context.Context, ID string) (opinion.Opinion, error)
	UpdateOpinion(ctx context.Context, ID string, newOpi opinion.Opinion) (opinion.Opinion, error)
	DeleteOpinion(ctx context.Context, ID string) error
}

func (h *Handler) PostOpinion(w http.ResponseWriter, r *http.Request) {
	var opi opinion.Opinion
	if err := json.NewDecoder(r.Body).Decode(&opi); err != nil {
		return
	}

	opi, err := h.Service.PostOpinion(r.Context(), opi)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(opi); err != nil {
		panic(err)
	}
}

func (h *Handler) GetOpinion(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateOpinion(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteOpinion(w http.ResponseWriter, r *http.Request) {

}
