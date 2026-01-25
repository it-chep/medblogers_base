package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/action/mm/action/create_mm/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mastermind/v1"
)

func (i *Implementation) CreateMM(ctx context.Context, req *desc.CreateMMRequest) (resp *desc.CreateMMResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/mm/create", func(ctx context.Context) error {
		resp = &desc.CreateMMResponse{}

		err := i.admin.Actions.MMModule.CreateMM.Do(ctx, dto.CreateMMRequest{
			//MMDatetime: req.GetMmDatetime(), // todo надо распарсить дату
			Name:   req.GetName(),
			MMLink: req.GetMmLink(),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
