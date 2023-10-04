package dto

import "github.com/dgnydn/binance-order/internal/app/order/domain"

type CreateOrderRequest struct {
	Order domain.Order `json:"order"`
}

type CreateOrderResponse struct {
	Status bool         `json:"status"`
	Order  domain.Order `json:"order"`
}

type GetOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

type GetOrderResponse struct {
	Order domain.Order `json:"order"`
}

type GetAllOrdersResponse struct {
	Orders []domain.Order `json:"orders"`
}

type UpdateOrderRequest struct {
	OrderId int64        `json:"orderId"`
	Order   domain.Order `json:"order"`
}

type UpdateOrderResponse struct {
	Status bool `json:"status"`
}

type DeleteOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

type DeleteOrderResponse struct {
	Status bool `json:"status"`
}
