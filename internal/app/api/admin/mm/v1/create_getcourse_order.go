package v1

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/mm/action/create_getcourse_order/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mastermind/v1"
)

func (i *Implementation) CreateGetCourseOrder(ctx context.Context, req *desc.CreateGetCourseOrderRequest) (*desc.CreateGetCourseOrderResponse, error) {
	err := i.admin.Actions.MMModule.CreateGetcourceOrder.Do(ctx, dto.CreateOrderRequest{
		OrderID:  req.GetOrderId(),
		GkID:     req.GetGetcourceId(),
		Position: req.GetPosition(),
		UserName: req.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.CreateGetCourseOrderResponse{}, nil
}
