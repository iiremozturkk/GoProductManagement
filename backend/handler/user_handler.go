package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"GoProductManagement/models"
	"GoProductManagement/service"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("JSON decode error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Geçersiz istek formatı",
			"error": err.Error(),
		})
		return
	}

	// Validasyon
	if u.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Ad alanı zorunludur",
		})
		return
	}

	if u.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "E-posta alanı zorunludur",
		})
		return
	}

	if u.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Şifre alanı zorunludur",
		})
		return
	}

	_, err := h.Service.Register(u)
	if err != nil {
		log.Printf("Register error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errMsg := fmt.Sprintf("Kayıt sırasında hata: %v", err)
		if err.Error() == "UNIQUE constraint failed: users.email" {
			errMsg = "Bu e-posta adresi zaten kullanılıyor"
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": errMsg,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Kayıt başarılı",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON decode error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Geçersiz istek formatı",
		})
		return
	}

	// Validasyon
	if req.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "E-posta alanı zorunludur",
		})
		return
	}

	if req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Şifre alanı zorunludur",
		})
		return
	}

	// Veritabanından kullanıcıyı kontrol et
	user, err := h.Service.FindByEmail(req.Email)
	if err != nil {
		log.Printf("FindByEmail error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 401 yerine 200 kullanıyoruz
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "E-posta veya şifre hatalı",
		})
		return
	}

	// Şifre kontrolü
	if user.Password != req.Password { // Gerçek projede hash karşılaştırması yapılmalı
		log.Printf("Password mismatch for user: %s", req.Email)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 401 yerine 200 kullanıyoruz
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "E-posta veya şifre hatalı",
		})
		return
	}

	// Başarılı giriş
	log.Printf("Successful login for user: %s", req.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{
		Success: true,
		Message: "Giriş başarılı",
	})
} 