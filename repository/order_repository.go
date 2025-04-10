package repository

import (
	"github.com/jmoiron/sqlx"
	"order-service/db"
	"order-service/domain"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() *Repository {
	return &Repository{db: db.GetConnection()}
}

func (r *Repository) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order

	query := `SELECT id AS order_id, product_id, quantities FROM orders`

	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) GetOrderByID(orderID uint) (*domain.Order, error) {
	query := `SELECT id AS order_id, product_id, quantities FROM orders WHERE id = $1`

	var order domain.Order
	err := r.db.Get(&order, query, orderID)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repository) CreateOrder(order domain.Order) (*domain.Order, error) {
	query := `
		INSERT INTO orders (product_id, quantities)
		VALUES (:product_id, :quantities)
		RETURNING id AS order_id, product_id, quantities
	`

	rows, err := r.db.NamedQuery(query, &order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&order)
		if err != nil {
			return nil, err
		}
	}

	return &order, nil
}

func (r *Repository) UpdateOrder(order domain.Order) (*domain.Order, error) {
	query := `
		UPDATE orders
		SET product_id = :product_id, quantities = :quantities
		WHERE id = :order_id
	`

	_, err := r.db.NamedExec(query, &order)
	if err != nil {
		return nil, err
	}

	return r.GetOrderByID(order.OrdetID)
}

func (r *Repository) DeleteOrder(orderID uint) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.Exec(query, orderID)
	return err
}
