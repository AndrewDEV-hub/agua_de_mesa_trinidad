package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	// Establece el directorio desde el cual se sirven los archivos estáticos
	fs := http.FileServer(http.Dir("./src"))

	// Manejador para la raíz
	http.Handle("/", fs)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
