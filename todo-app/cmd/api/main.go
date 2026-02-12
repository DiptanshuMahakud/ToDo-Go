package main

import (
	"context"
	"log"
	"net/http"

	"github.com/DiptanshuMahakud/ToDo-Go/internal/config"
	"github.com/DiptanshuMahakud/ToDo-Go/internal/db"
	"github.com/DiptanshuMahakud/ToDo-Go/internal/todo"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", getHealth)

	cfg := config.Load()

	ctx := context.Background()
	dbpool, err := db.New(ctx, cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	repo := todo.NewPostgresRepo(dbpool)
	service := todo.NewService(repo)
	handler := todo.Newhandler(service)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTodo(w, r)
		case http.MethodGet:
			handler.ListTodos(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	})

	log.Println("Starting on port :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
