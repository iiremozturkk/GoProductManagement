package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"GoProductManagement/config"
	"GoProductManagement/handler"
	"GoProductManagement/repository"
	"GoProductManagement/service"

	"github.com/gorilla/mux"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Her isteği logla
        log.Printf("Gelen İstek: %s %s", r.Method, r.URL.Path)

        // CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.Header().Set("Content-Type", "application/json")

        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       
        if r.Method == "OPTIONS" {
            next.ServeHTTP(w, r)
            return
        }

        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

       
        if len(token) > 7 && token[0:7] == "Bearer " {
            token = token[7:]
        }

        
        next.ServeHTTP(w, r)
    })
}

func main() {
	
	dbHost := getEnv("DB_HOST", "IREM")
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "1433"))
	dbUser := getEnv("DB_USER", "sa")
	dbPassword := getEnv("DB_PASSWORD", "1234")
	dbName := getEnv("DB_NAME", "ProductDB")

	
	var db *sql.DB
	var err error
	maxRetries := 5
	retryDelay := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		cfg := config.DBConfig{
			Server:   dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Database: dbName,
		}

		db, err = config.ConnectDB(cfg)
		if err == nil {
			break
		}

		log.Printf("Attempts to connect to database %d failed: %v", i+1, err)
		if i < maxRetries-1 {
			log.Printf("Retrying after %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	
	
	r.HandleFunc("/register", userHandler.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", userHandler.Login).Methods("POST", "OPTIONS")

	
	protected := r.PathPrefix("").Subrouter()
	protected.Use(authMiddleware)
	handler.RegisterProductRoutes(protected, productHandler)

	
	handler := enableCORS(r)

	
	port := ":8080"
	fmt.Printf("Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
} 
