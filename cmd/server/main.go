package main

import (
	"context"
	"fmt"

	"github.com/untaldouglas/go-rest-api/internal/db"
	"github.com/untaldouglas/go-rest-api/internal/opinion"
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
	// para verificar función primera de paquete de database
	// GetOpinion.
	// buscar en web donde crear UUID a pasar a función,
	// crear antes registro con dicho UUID usando dBeaver por ej
	opiService := opinion.NewService(db)
	fmt.Println(opiService.GetOpinion(
		context.Background(),
		"96cab9f6-344a-44e8-8b4d-dda2a1033af5",
	))

	opiService = opinion.NewService(db)
	opiService.PostOpinion(
		context.Background(),
		opinion.Opinion{
			ID:        "3c4e8041-71aa-4a33-9661-17dea10dca87",
			Asunto:    "Testing desde main app",
			Autor:     "API-testing",
			Contenido: "Envio de post desde app como prueba dg",
		},
	)

	fmt.Println("*** Succesfully connected and pinged database !!")
	return nil
}

func main() {
	fmt.Println("Go REST API bienvenidos !")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
