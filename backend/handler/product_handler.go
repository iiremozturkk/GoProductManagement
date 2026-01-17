package handler 

import (
	"database/sql"
	"encoding/json" 
	"net/http" 
	"strconv" 

	"GoProductManagement/models" 
	"GoProductManagement/service" 

	"github.com/gorilla/mux" 
)

type ProductHandler struct { 
	Service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler { 
	return &ProductHandler{Service: s}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product 
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) 
		return
	}
	id, err := h.Service.Create(p) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
		return
	}
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(map[string]int{"id": id}) 
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetAll() 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products) 
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, err := h.Service.GetByID(id)  
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product) 
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var p models.Product 
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = id 
	if err := h.Service.Update(p); err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK) 
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) 
}

func RegisterProductRoutes(r *mux.Router, h *ProductHandler) {
	r.HandleFunc("/products", h.CreateProduct).Methods("POST", "OPTIONS") 
	r.HandleFunc("/products", h.GetAllProducts).Methods("GET", "OPTIONS") 
	r.HandleFunc("/products/{id}", h.GetProductByID).Methods("GET", "OPTIONS") 
	r.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT", "OPTIONS") 
	r.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE", "OPTIONS") 
	r.HandleFunc("/products/{id}/stock", h.UpdateStock).Methods("PUT", "OPTIONS") 
}


type UpdateStockRequest struct {
	Quantity int `json:"quantity"`
}

func (h *ProductHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req UpdateStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		http.Error(w, "Quantity must be positive", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateStock(id, req.Quantity); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Insufficient stock", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
} 