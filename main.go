package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/dist
var embeddedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embeddedFiles, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/url-decoder/", http.StripPrefix("/url-decoder/", http.FileServer(http.FS(fsys))))

	log.Println("listening on http://localhost:8123")
	err = http.ListenAndServe(":8123", nil)
	if err != nil {
		log.Fatal(err)
	}
}
