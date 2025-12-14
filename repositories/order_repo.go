package repositories

import (
	"database/sql"
	"restaurant-api/models"
)

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo { return &OrderRepo{DB: db} }

func (r *OrderRepo) Create(o *models.Order) error {
	q := `INSERT INTO orders (user_id,total,status) VALUES ($1,$2,$3) RETURNING id, created_at`
	return r.DB.QueryRow(q, o.UserID, o.Total, o.Status).Scan(&o.ID, &o.CreatedAt)
}

func (r *OrderRepo) AddItem(it *models.OrderItem) error {
	q := `INSERT INTO order_items (order_id,menu_id,quantity,price) VALUES ($1,$2,$3,$4) RETURNING id`
	return r.DB.QueryRow(q, it.OrderID, it.MenuID, it.Quantity, it.Price).Scan(&it.ID)
}

func (or *OrderRepo) FindByUserID(userID int) ([]models.Order, error) {
	rows, err := or.DB.Query(`
		SELECT id, user_id, total, status, created_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Total, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
