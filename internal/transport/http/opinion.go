package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/untaldouglas/go-rest-api/internal/opinion"
)

type OpinionService interface {
	PostOpinion(context.Context, opinion.Opinion) (opinion.Opinion, error)
	GetOpinion(ctx context.Context, ID string) (opinion.Opinion, error)
	UpdateOpinion(ctx context.Context, ID string, newOpi opinion.Opinion) (opinion.Opinion, error)
	DeleteOpinion(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

// Creamos structs para cada petición/request a procesar
// en donde describimos atributos de validación en json
type PostOpinionRequest struct {
	Asunto    string `json:"asunto" validate:"required"`
	Contenido string `json:"contenido" validate:"required"`
	Autor     string `json:"autor" validate:"required"`
}

// convertPostOpinionRequestToOpinion - convierte struct validado a opinion a procesar en servicio y db
func convertPostOpinionRequestToOpinion(o PostOpinionRequest) opinion.Opinion {
	return opinion.Opinion{
		Asunto:    o.Asunto,
		Autor:     o.Autor,
		Contenido: o.Contenido,
	}
}

func (h *Handler) PostOpinion(w http.ResponseWriter, r *http.Request) {
	// Al agregar los structs por petición y sus validaciones json
	// modificamos entonces el type que procesamos en las funciones
	// var opi opinion.Opinion
	var opi PostOpinionRequest
	if err := json.NewDecoder(r.Body).Decode(&opi); err != nil {
		return
	}

	validar := validator.New()
	err := validar.Struct(opi)
	if err != nil {
		http.Error(w, "datos invalidos en opinion", http.StatusBadRequest)
		return
	}

	convertedOpinion := convertPostOpinionRequestToOpinion(opi)
	postedOpinion, err := h.Service.PostOpinion(r.Context(), convertedOpinion)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(postedOpinion); err != nil {
		panic(err)
	}
}

func (h *Handler) GetOpinion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	opi, err := h.Service.GetOpinion(r.Context(), id)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(opi); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateOpinion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var opi opinion.Opinion
	if err := json.NewDecoder(r.Body).Decode(&opi); err != nil {
		return
	}

	opi, err := h.Service.UpdateOpinion(r.Context(), id, opi)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(opi); err != nil {
		panic(err)
	}

}

func (h *Handler) DeleteOpinion(w http.ResponseWriter, r *http.Request) {
	// Get the id from the request
	vars := mux.Vars(r)
	commentID := vars["id"]
	if commentID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.Service.DeleteOpinion(r.Context(), commentID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "se borró opinión con éxito"}); err != nil {
		panic(err)
	}

}
