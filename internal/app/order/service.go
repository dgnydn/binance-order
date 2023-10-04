package order

import (
	"context"
	"emperror.dev/errors"
	"github.com/dgnydn/binance-order/internal/app/order/domain"
	"github.com/dgnydn/binance-order/internal/app/order/dto"
	"github.com/dgnydn/binance-order/internal/app/order/store"
	"github.com/dgnydn/binance-order/internal/common"
)

// +kit:endpoint:errorStrategy=service
// +kit:transport:protoFilePath=../../../api/binance-order/v1/order.proto

type Service interface {
	CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
	GetOrder(ctx context.Context, request dto.GetOrderRequest) (dto.GetOrderResponse, error)
	GetAllOrders(ctx context.Context) (dto.GetAllOrdersResponse, error)
	UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) (dto.UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, request dto.DeleteOrderRequest) (dto.DeleteOrderResponse, error)
}

type service struct {
	store store.Store
	l     common.Logger
}

func (s service) CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	if !request.Order.IsValid() {
		return dto.CreateOrderResponse{}, errors.New("Order is invalid")
	}

	order, err := s.store.CreateOrder(request.Order)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	return dto.CreateOrderResponse{
		Status: true,
		Order:  order,
	}, nil
}

func (s service) GetOrder(ctx context.Context, request dto.GetOrderRequest) (dto.GetOrderResponse, error) {
	if request.OrderId == 0 {
		return dto.GetOrderResponse{}, errors.New("Order ID is missing")
	}

	order, err := s.store.GetOrder(request.OrderId)
	if err != nil {
		return dto.GetOrderResponse{}, err
	}

	return dto.GetOrderResponse{Order: order}, nil
}

func (s service) GetAllOrders(ctx context.Context) (dto.GetAllOrdersResponse, error) {
	var orders []domain.Order
	var err error
	if orders, err = s.store.GetAllOrders(); err != nil {
		return dto.GetAllOrdersResponse{}, err
	}

	return dto.GetAllOrdersResponse{Orders: orders}, nil
}

func (s service) UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) (dto.UpdateOrderResponse, error) {
	if request.OrderId == 0 {
		return dto.UpdateOrderResponse{}, errors.New("Order ID is missing")
	}

	if !request.Order.IsValid() {
		return dto.UpdateOrderResponse{}, errors.New("Order is invalid")
	}

	if err := s.store.UpdateOrder(request.OrderId, request.Order); err != nil {
		return dto.UpdateOrderResponse{}, err
	}

	return dto.UpdateOrderResponse{
		Status: true,
	}, nil
}

func (s service) DeleteOrder(ctx context.Context, request dto.DeleteOrderRequest) (dto.DeleteOrderResponse, error) {
	if request.OrderId == 0 {
		return dto.DeleteOrderResponse{}, errors.New("Order ID is missing")
	}

	if err := s.store.DeleteOrder(request.OrderId); err != nil {
		return dto.DeleteOrderResponse{}, err
	}

	return dto.DeleteOrderResponse{
		Status: true,
	}, nil
}

func NewService(store store.Store, l common.Logger) Service {
	return &service{store: store, l: l}
}
