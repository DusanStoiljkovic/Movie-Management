package handler

import (
	"encoding/json"
	"movie-management/internal/models"
	"movie-management/internal/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Register(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string
		Password string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := handler.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	req.Password = ""

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("Role")

	users, err := h.service.GetAllUsers(r.Context(), role)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	json.NewEncoder(w).Encode(users)
}
