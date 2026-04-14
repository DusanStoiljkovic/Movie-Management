package handler

import (
	"encoding/json"
	dto "movie-management/internal/dto/user"
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
	var req dto.RequestRegisterUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := handler.service.Register(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestLoginUser

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

func (h *UserHandler) AddFavouriteGenres(w http.ResponseWriter, r *http.Request) {
	return
}
