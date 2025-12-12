package repositories

import (
	"database/sql"
	"restaurant-api/models"
)

type CategoryRepo struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo { return &CategoryRepo{DB: db} }

func (r *CategoryRepo) Create(c *models.Category) error {
	q := `INSERT INTO categories (name, description) VALUES ($1,$2) RETURNING id`
	return r.DB.QueryRow(q, c.Name, c.Description).Scan(&c.ID)
}

func (r *CategoryRepo) GetAll() ([]models.Category, error) {
	rows, err := r.DB.Query(`SELECT id,name,description FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, nil
}
