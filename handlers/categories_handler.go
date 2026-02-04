package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tugas/models"
	"tugas/services"
)

type CategoriesHandler struct {
	service *services.CategoriesService
}

func NewCategoriesHandler(service *services.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{service:service}
}

func (h *CategoriesHandler) HandleCategories(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoriesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error (w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoriesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var categories models.Barang
	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		http.Error (w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.Create(&categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoriesHandler) HandleCategoriesByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoriesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Categories ID tidak valid", http.StatusBadRequest)
		return
	}

	categories, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoriesHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Categories ID tidak valid", http.StatusBadRequest)
		return
	}

	var categories models.Barang
	err = json.NewDecoder(r.Body).Decode(&categories)
	if err != nil{
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	categories.ID = id
	err = h.service.Update(&categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoriesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Categories ID tidak valid", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kategori berhasil dihapus",
	})
}