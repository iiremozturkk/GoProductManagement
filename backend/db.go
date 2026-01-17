package main

import (
	"GoProductManagement/config"
	"fmt"
	"log"
	"time"
)

func waitForDB(cfg config.DBConfig) error {
	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		log.Printf("Attempting to connect to SQL Server... (Tries %d/%d)", attempt, maxAttempts)
		db, err := config.ConnectDB(cfg)
		if err == nil {
			db.Close()
			return nil
		}
		if attempt < maxAttempts {
			time.Sleep(15 * time.Second)
		}
	}
	return fmt.Errorf("Failed to connect to SQL Server.")
}

