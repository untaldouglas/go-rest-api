package main

import "fmt"

// Run - Responsable de inicializar el module
// tareas como iniciar conexión a DB van acá
func Run() error {
	fmt.Println("Iniciando la aplicación ...")
	return nil
}

func main() {
	fmt.Println("Go REST API bienvenidos !")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
