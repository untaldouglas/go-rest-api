package grpc

import (
	"context"
	"log"
	"net"

	"github.com/untaldouglas/go-rest-api/internal/opinion"
	opi "github.com/untaldouglas/go-rest-api/repositorio-protos/opinion/v1"
	"google.golang.org/grpc"
)

// OpinionService - define la interface que controla
// la implementación de los servicios a servir.
type OpinionService interface {
	GetOpinion(ctx context.Context, ID string) (opinion.Opinion, error)
	PostOpinion(context.Context, opinion.Opinion) (opinion.Opinion, error)
	DeleteOpinion(ctx context.Context, ID string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	OpinionService OpinionService
}

// New - returns a new gRPC handler
func New(opiService OpinionService) Handler {
	return Handler{
		OpinionService: opiService,
	}
}

// TODO: HABILITAR PUERTO 50051 EN DOCKER
func (h Handler) Serve() error {
	listento, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	opi.RegisterOpinionServicioServer(grpcServer, &h)

	if err := grpcServer.Serve(listento); err != nil {
		log.Printf("failed to serve grpc: %s\n", err)
		return err
	}

	return nil
}

func (h Handler) GetOpinion(ctx context.Context, req *opi.GetOpinionRequest) (*opi.GetOpinionResponse, error) {
	return &opi.GetOpinionResponse{}, nil
}

func (h Handler) PostOpinion(ctx context.Context, req *opi.PostOpinionRequest) (*opi.PostOpinionResponse, error) {
	return &opi.PostOpinionResponse{}, nil
}

func (h Handler) DeleteOpinion(ctx context.Context, req *opi.DeleteOpinionRequest) (*opi.DeleteOpinionResponse, error) {
	return &opi.DeleteOpinionResponse{}, nil
}
