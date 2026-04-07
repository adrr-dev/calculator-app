package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/adrr-dev/calculator-app/internal/handlers"
	"github.com/adrr-dev/calculator-app/internal/repository"
	"github.com/adrr-dev/calculator-app/internal/service"
)

func main() {
	dataFile := "data.json"
	myRepo := &repository.Repo{DataFile: dataFile}
	myService := service.NewService(myRepo)

	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fragments, err := template.ParseGlob("templates/fragments/*.html")
	if err != nil {
		log.Fatal(err)
	}
	myHandling := handlers.NewHandling(tmpls, fragments, myService)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /{$}", myHandling.RootHandler)
	mux.HandleFunc("POST /key", myHandling.KeyHandler)
	mux.HandleFunc("POST /enter", myHandling.EnterHandler)
	mux.HandleFunc("POST /clear", myHandling.ClearHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
