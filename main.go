package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Define un puerto por defecto si PORT no está configurado
        log.Printf("Defaulting to port %s", port)
    }

    // Manejador para servir el sitemap.xml
    http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "sitemap.xml")
    })

    // Manejador para servir otros archivos estáticos desde el directorio ./src
    http.Handle("/", http.FileServer(http.Dir("./src")))

    // Iniciar el servidor HTTP en el puerto especificado por la variable de entorno PORT
    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
