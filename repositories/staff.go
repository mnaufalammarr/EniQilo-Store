package repositories

import (
	"EniQilo/entities"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRepository interface {
	Create(staff entities.StaffRequast) (entities.Staff, error)
	FindById(id int) (entities.Staff, error)
	FindByUserId(id int) (entities.Staff, error)
}

type staffRepository struct {
	db *pgxpool.Pool
}

func NewStaffRepository(db *pgxpool.Pool) *staffRepository {
	return &staffRepository{db}
}

func (r *staffRepository) Create(staff entities.StaffRequast) (entities.Staff, error) {
	fmt.Print("stafId", staff.UserId)
	_, err := r.db.Exec(context.Background(), "INSERT INTO auths (user_id, password) VALUES ($1, $2)", staff.UserId, staff.Password)
	if err != nil {
		return entities.Staff{}, err
	}

	return entities.Staff{}, nil
}

func (r *staffRepository) FindById(id int) (entities.Staff, error) {
	var staff entities.Staff
	err := r.db.QueryRow(context.Background(), "SELECT id,pasword FROM auths WHERE id = $1", id).Scan(&staff.Id, &staff.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Staff{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return entities.Staff{}, err // Error lainnya
	}

	return staff, nil
}
func (r *staffRepository) FindByUserId(id int) (entities.Staff, error) {
	var staff entities.Staff
	fmt.Println("id", id)
	err := r.db.QueryRow(context.Background(), "SELECT id,password FROM auths WHERE user_id = $1", id).Scan(&staff.Id, &staff.Password)

	fmt.Println("staff", staff)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Staff{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return entities.Staff{}, err // Error lainnya
	}

	return staff, nil
}
