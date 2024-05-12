package repositories

import (
	"EniQilo/entities"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	Create(order entities.Order) (string, error)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order entities.Order) (string, error) {
	query := `INSERT INTO orders (id, customer_id, cashier_id, paid, change) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var totalAmount int = 0

	tx, err := r.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	fmt.Println(query)
	_, erro := tx.Exec(context.Background(), query, order.Id, order.CustomerID, order.CashierID, order.Paid, order.Change)
	if erro != nil {
		return "", fmt.Errorf("failed to insert order: %w", erro)
	}

	for _, element := range order.ProductDetails {
		var enough bool
		var price int
		//check stock
		if err = tx.QueryRow(context.Background(), "SELECT (stock >= $1), price FROM products WHERE id = $2", element.Quantity, element.ProductId).Scan(&enough, &price); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return "", fmt.Errorf("NO SUCH PRODUCT SELECTED")
			}
			fmt.Printf("ERROR SELECT product: %s", err.Error())
			return "", err
		}

		if !enough {
			return "", fmt.Errorf("PRODUCT SELECTED WITH ID: %s QUANTITY IS NOT ENOUGH", element.ProductId)
		}

		// Update the album inventory to remove the quantity in the order.
		_, err = tx.Exec(context.Background(), "UPDATE products SET stock = stock - $1 WHERE id = $2",
			element.Quantity, element.ProductId)
		if err != nil {
			fmt.Printf("ERROR UPDATE: %s", err.Error())
			return "", fmt.Errorf(err.Error())
		}
		totalAmount += price * element.Quantity
		// Update the album inventory to remove the quantity in the order.
		_, err = tx.Exec(context.Background(), "INSERT INTO order_details (product_id, order_id, quantity) VALUES ($1, $2, $3)", element.ProductId, order.Id, element.Quantity)

		if err != nil {
			fmt.Printf("ERROR INSERT: %s", err.Error())
			return "", fmt.Errorf(err.Error())
		}
	}

	if order.Paid < totalAmount {
		return "", fmt.Errorf("CUSTOMER PAID IS NOT ENOUGH")
	}

	if erro := tx.Commit(context.Background()); erro != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", erro)
	}

	return order.Id, nil
}
