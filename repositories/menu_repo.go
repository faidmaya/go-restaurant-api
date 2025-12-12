package repositories

import (
	"database/sql"
	"restaurant-api/models"
)

type MenuRepo struct {
	DB *sql.DB
}

func NewMenuRepo(db *sql.DB) *MenuRepo { return &MenuRepo{DB: db} }

func (r *MenuRepo) Create(m *models.Menu) error {
	q := `INSERT INTO menus (name,description,price,category_id) VALUES ($1,$2,$3,$4) RETURNING id, created_at`
	return r.DB.QueryRow(q, m.Name, m.Description, m.Price, m.CategoryID).Scan(&m.ID, &m.CreatedAt)
}

func (r *MenuRepo) GetAll() ([]models.Menu, error) {
	rows, err := r.DB.Query(`SELECT id,name,description,price,category_id,created_at FROM menus`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Menu
	for rows.Next() {
		var m models.Menu
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.Price, &m.CategoryID, &m.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, nil
}
