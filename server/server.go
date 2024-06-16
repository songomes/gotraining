package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request processada")
	select {
	case <-ctx.Done():
		log.Println("request cancelada")
	case <-time.After(5 * time.Second):
		log.Println("request concluída")
	}
}
