package main

import (
	"fmt"

	"github.com/untaldouglas/go-rest-api/internal/db"
	"github.com/untaldouglas/go-rest-api/internal/opinion"
	"github.com/untaldouglas/go-rest-api/internal/transport/grpc"
	transportHttp "github.com/untaldouglas/go-rest-api/internal/transport/http"
)

// Run - Responsable de inicializar el module
// tareas como iniciar conexión a DB van acá
func Run() error {
	fmt.Println("Iniciando la aplicación ...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error al conectarse a la base de datos")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Error al migrar la database")
		return err
	}
	// opiService - Inicializa un servicio a ser llamado
	opiService := opinion.NewService(db)
	// http canal
	httpHandler := transportHttp.NewHandler(opiService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	// grpc canal
	grpcHandler := grpc.New(opiService)
	if err := grpcHandler.Serve(); err != nil {
		return err
	}

	fmt.Println("*** Succesfully connected and pinged database !!")
	return nil
}

func main() {
	fmt.Println("Go REST API bienvenidos !")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
