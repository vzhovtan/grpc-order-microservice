package grpc

import (
	"context"
	"microservices/order/internal/application/core/domain"
	"microservices/order/proto"
)

func (a Adapter) Create(ctx context.Context, request *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &proto.CreateOrderResponse{OrderId: result.ID}, nil
}
