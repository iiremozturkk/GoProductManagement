package config

import (
	"database/sql" 
	"fmt" 
	_ "github.com/denisenkom/go-mssqldb"
)

type DBConfig struct {
	Server   string 
	Port     int 
	User     string 
	Password string 
	Database string 
}

func GetConnectionString(cfg DBConfig) string { 
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", cfg.Server, cfg.User, cfg.Password, cfg.Port, cfg.Database)
}

func ConnectDB(cfg DBConfig) (*sql.DB, error) {
	connString := GetConnectionString(cfg) 
	db, err := sql.Open("sqlserver", connString) 
	if err != nil {
		return nil, err 
	}
	if err = db.Ping(); err != nil { 
		return nil, err 
	}
	return db, nil
}             