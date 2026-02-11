package todo

import (
	"encoding/json"
	"net/http"
)

// this is the actual handler that will handle all the http requests

type Handler struct {
	service *Service
}

func Newhandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	} // this is the json struct of the request

	err := json.NewDecoder(r.Body).Decode(&req)
	// decode the json from the body and put it inside my req

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateTodo(r.Context(), req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.ListTodos(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
