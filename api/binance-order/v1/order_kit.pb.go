// Code generated by protoc-gen-go-kit. DO NOT EDIT.
// versions:
// - protoc-gen-go-kit
// - protoc            v3.21.12

package orderv1

import (
	context "context"
)

// OrderServiceHandler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
//
// This interface is based on github.com/go-kit/kit/transport/grpc.Handler.
type OrderServiceHandler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

// OrderServiceKitServer is the Go kit server implementation for OrderService service.
type OrderServiceKitServer struct {
	*UnimplementedOrderServiceServer

	CreateOrderHandler  OrderServiceHandler
	GetOrderHandler     OrderServiceHandler
	GetAllOrdersHandler OrderServiceHandler
	UpdateOrderHandler  OrderServiceHandler
	DeleteOrderHandler  OrderServiceHandler
}

func (s OrderServiceKitServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	_, resp, err := s.CreateOrderHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*CreateOrderResponse), nil
}

func (s OrderServiceKitServer) GetOrder(ctx context.Context, req *GetOrderRequest) (*GetOrderResponse, error) {
	_, resp, err := s.GetOrderHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*GetOrderResponse), nil
}

func (s OrderServiceKitServer) GetAllOrders(ctx context.Context, req *GetAllOrdersRequest) (*GetAllOrdersResponse, error) {
	_, resp, err := s.GetAllOrdersHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*GetAllOrdersResponse), nil
}

func (s OrderServiceKitServer) UpdateOrder(ctx context.Context, req *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	_, resp, err := s.UpdateOrderHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*UpdateOrderResponse), nil
}

func (s OrderServiceKitServer) DeleteOrder(ctx context.Context, req *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	_, resp, err := s.DeleteOrderHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*DeleteOrderResponse), nil
}
