package opinion

import (
	"context"
	"errors"
	"fmt"
)

// Al utilizar constantes en el manejo de errores nos aseguramos
// no compartir información sobre nuestra implementación (DB metadata por ej)
// al exterior del module.
var (
	ErrFetchingOpinion = errors.New("error en la búsqueda de la opinión por ID")
	ErrNoImplementado  = errors.New("funcionalidad no implementada")
)

// Opinion - representación o estructura de la opinión
// es la data structure del servicio
type Opinion struct {
	ID        string
	Asunto    string
	Contenido string
	Autor     string
}

// Store - this interface defines all the methods
// that our service needs in order to operate
type Store interface {
	GetOpinion(context.Context, string) (Opinion, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetOpinion - es el nivel de inteligencia de negocio
// s.Store.GetOpinion es el respository, donde la data se encuentra DB
func (s *Service) GetOpinion(ctx context.Context, id string) (Opinion, error) {
	fmt.Println("Obtener una opinión")
	opi, err := s.Store.GetOpinion(ctx, id)
	if err != nil {
		fmt.Println(err)                     // este log es para uso interno de monitoreo, no se expone fuera de module
		return Opinion{}, ErrFetchingOpinion //regresamos error como constante declarada
	}

	return opi, nil
}

func (s *Service) ActualizarOpinion(ctx context.Context, opi Opinion) error {
	return ErrNoImplementado
}

func (s *Service) BorrarOpinion(ctx context.Context, id string) error {
	return ErrNoImplementado
}

func (s *Service) CrearOpinion(ctx context.Context, opi Opinion) (Opinion, error) {
	return Opinion{}, ErrNoImplementado
}
