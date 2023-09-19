package database

import (
	"database/sql"

	"github.com/TiagoSilvaLourenco/challenge-clean-arch-go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var id string
		var price, tax, finalPrice float64
		if err := rows.Scan(&id, &price, &tax, &finalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, entity.Order{
			ID:         id,
			Price:      price,
			Tax:        tax,
			FinalPrice: finalPrice,
		})
	}
	return orders, nil
}

// func (r *OrderRepository) ListOrders(page, limit int, sort string) ([]entity.Order, error) {

// 	if page < 1 {
// 		page = 1
// 	}
// 	if limit < 1 {
// 		limit = 10
// 	}

// 	query := "Select * from orders"
// 	if sort != "" {
// 		query += " ORDER BY " + sort
// 	}
// 	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, (page-1)*limit)

// 	rows, err := r.Db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	orders := []entity.Order{}

// 	for rows.Next() {
// 		var id string
// 		var price, tax, final_price float64
// 		if err := rows.Scan(&id, &price, &tax, &final_price); err != nil {
// 			return nil, err
// 		}
// 		orders = append(orders, entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: final_price})
// 	}
// 	return orders, nil

// }
