package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/lcastelobranco/front/handler"
	"log"
	"net/http"
	"os"
)

//go:embed public
var FS embed.FS

func main() {

	err := initEverything()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	//slog.Info("application started", "port", port)
	port := os.Getenv("HTTP_LISTEN_ADDR")
	log.Println("application started", "port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initEverything() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil

}
