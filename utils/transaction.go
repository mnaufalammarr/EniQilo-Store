package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func withTransaction(pool *pgxpool.Pool, fn func(*pgx.Tx) error) error {
	tx, err := pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rollbackErr) // Log or handle rollback error (optional)
			}
			return
		}
		if err := tx.Commit(context.Background()); err != nil {
			fmt.Printf("failed to commit transaction: %v\n", err) // Log or handle commit error (optional)
		}
	}()

	err = fn(&tx)
	return err
}

//func createUserWithStaff(pool *pgxpool.Pool, user entities.UserRequest, staff entities.SignUpRequest) error {
//
//}
