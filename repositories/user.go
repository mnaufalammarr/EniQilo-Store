package repositories

import (
	"EniQilo/entities"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	FindAll(filterParams entities.UserQueryParams, isCustomer bool) ([]entities.UserResponse, error)
	Create(user entities.User) (entities.User, error)
	FindByPhone(phone string) (entities.User, error)
	PhoneIsExist(phone string) bool
	FindById(id int) (entities.User, error)
	CreateUser(user entities.User) (int, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user entities.User) (entities.User, error) {
	_, err := r.db.Exec(context.Background(), "INSERT INTO users ( name, phone_number , role ) VALUES ($1, $2, $3 ) RETURNING id", user.Name, user.Phone, user.Role)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user entities.User) (int, error) {
	var userId int

	// Replace "users" with your actual table name
	query := `INSERT INTO users (name, phone_number, role) VALUES ($1, $2, $3) RETURNING id`

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil { // Rollback on error
			if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rollbackErr)
			}
		}
	}()

	row := tx.QueryRow(context.Background(), query, user.Name, user.Phone, user.Role)
	err = row.Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user or scan ID: %w", err)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return userId, nil
}
func (r *userRepository) FindByPhone(phone string) (entities.User, error) {
	var user entities.User
	err := r.db.QueryRow(context.Background(), "SELECT id, phone_number, name , role  FROM users WHERE phone_number = $1", phone).Scan(&user.Id, &user.Phone, &user.Name, &user.Role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.User{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return entities.User{}, err // Error lainnya
	}
	return user, nil
}

func (r *userRepository) PhoneIsExist(phone string) bool {
	var exist string
	err := r.db.QueryRow(context.Background(), "SELECT phone_number FROM users WHERE phone_number = $1 LIMIT 1", phone).Scan(&exist)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
	}
	return true
}

func (r *userRepository) FindById(id int) (entities.User, error) {
	var user entities.User
	err := r.db.QueryRow(context.Background(), "SELECT id, phone_number, name , role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Phone, &user.Name, &user.Role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.User{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return entities.User{}, err // Error lainnya
	}

	return user, nil
}
func (r *userRepository) FindAll(filterParams entities.UserQueryParams, isCustomer bool) ([]entities.UserResponse, error) {
	query := "SELECT id, name, phone_number FROM users WHERE 1=1"
	num := 1

	if isCustomer {
		query += " AND role = false "
	}

	// Filter by ID
	if filterParams.ID != "" {
		query += "AND id = '" + filterParams.ID + "' "
	}

	if filterParams.Name != "" {
		query += " AND LOWER(name) LIKE '%" + strings.ToLower(filterParams.Name) + "%' "
	}

	if filterParams.PhoneNumber != "" {
		query += " AND phone_number LIKE '+" + filterParams.PhoneNumber + "%'"
		num++
	}

	// query += (" ORDER BY id DESC ")

	if filterParams.Limit > 0 {
		query += " LIMIT " + strconv.Itoa(filterParams.Limit)
	}

	if filterParams.Offset >= 0 {
		query += " OFFSET  " + strconv.Itoa(filterParams.Offset)
	}
	fmt.Println(query)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.UserResponse
	for rows.Next() {
		var user entities.UserResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Phone)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// createdAtISO8601 := cat.CreatedAt.Format(time.RFC3339)
		userResponse := entities.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
		}
		users = append(users, userResponse)
	}
	if users == nil {
		users = make([]entities.UserResponse, 0)
	}
	return users, nil
}
