package main

import (
	"go-in-5/episode01/handlers"
	"go-in-5/episode01/storage"
	"log"
	"net/http"
)

func main() {
	db := storage.NewInMemDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/get", handlers.GetKey(db))
	mux.HandleFunc("/set", handlers.PutKey(db))
	log.Printf("Starting server on 8081...")
	err := http.ListenAndServe("localhost:8081", mux)
	log.Fatal(err)
}
