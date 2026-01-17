package repository

import (
	"database/sql"
	"GoProductManagement/models"
)

type UserRepository interface {
	Create(user models.User) (int, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) (int, error) {
	query := "INSERT INTO Users (Name, Email, Password) OUTPUT INSERTED.ID VALUES (@p1, @p2, @p3)"
	var id int
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	query := "SELECT ID, Name, Email, Password FROM Users WHERE Email = @p1"
	var u models.User
	err := r.db.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	return u, err
} 