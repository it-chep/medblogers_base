package v1

import (
	"context"
	renewalDto "medblogers_base/internal/modules/admin/entities/mm/action/getcourse_subscription_renewal/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mastermind/v1"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetCourseSubscriptionRenewal(ctx context.Context, req *desc.GetCourseSubscriptionRenewalRequest) (*desc.GetCourseSubscriptionRenewalResponse, error) {
	if req.GetApiKey() != os.Getenv("GK_SUBS_API_KEY") {
		return nil, status.Error(codes.PermissionDenied, "invalid api key")
	}

	err := i.admin.Actions.MMModule.GetCourseSubscriptionRenewal.Do(ctx, renewalDto.Request{
		GkID:      req.GetGkId(),
		DaysCount: req.GetDaysCount(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.GetCourseSubscriptionRenewalResponse{}, nil
}
