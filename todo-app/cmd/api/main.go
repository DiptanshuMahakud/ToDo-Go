package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// log.Println("Starting on port :8080")
	// err = http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	go func() {
		log.Println("Starting HTTP server on", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error : %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("shutdown signal recieved")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("server shutdown error:", err)

		dbpool.Close()

		log.Println("server closed cleanly")
	}

}
