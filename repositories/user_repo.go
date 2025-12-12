package repositories

import (
	"database/sql"
	"restaurant-api/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	query := `
        SELECT id, name, email, password, role, created_at
        FROM users
        WHERE email = $1
    `

	var user models.User
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) CreateUser(u *models.User) (*models.User, error) {
	query := `
        INSERT INTO users (name, email, password, role)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, email, role, created_at
    `

	var created models.User
	err := r.DB.QueryRow(query, u.Name, u.Email, u.Password, u.Role).Scan(
		&created.ID,
		&created.Name,
		&created.Email,
		&created.Role,
		&created.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}
