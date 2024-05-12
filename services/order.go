package services

import (
	"EniQilo/entities"
	"EniQilo/repositories"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type OrderService interface {
	Create(orderRequest entities.OrderRequest, cashierId int) (string, error)
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) *orderService {
	return &orderService{orderRepository}
}

func (s *orderService) Create(orderRequest entities.OrderRequest, cashierId int) (string, error) {
	orderReq := entities.Order{
		Id:             strconv.Itoa(uuid.New().ClockSequence()),
		CustomerID:     orderRequest.CustomerID,
		ProductDetails: orderRequest.ProductDetails,
		Paid:           orderRequest.Paid,
		Change:         orderRequest.Change,
		CashierID:      strconv.Itoa(cashierId),
	}

	orderId, err := s.orderRepository.Create(orderReq)

	fmt.Println("newOrder", orderId)
	fmt.Println("newOrderErr", err)

	return orderId, err
}
