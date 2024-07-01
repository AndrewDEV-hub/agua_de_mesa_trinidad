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

    // Ruta para servir el sitemap.xml
    http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "sitemap.xml")
    })

    // Ruta por defecto para servir otros archivos o contenido
    http.Handle("/", http.FileServer(http.Dir("./src")))

    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
