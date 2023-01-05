package main

import (
	"context"
	"fmt"

	"github.com/untaldouglas/go-rest-api/internal/db"
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

	// Eliminar posteriormente
	if err := db.Ping(context.Background()); err != nil {
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
