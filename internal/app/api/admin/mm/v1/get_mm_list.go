package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/mm/action/get_mm_list/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/mastermind/v1"
	"time"

	"github.com/samber/lo"
)

func (i *Implementation) GetMMList(ctx context.Context, req *desc.GetMMListRequest) (resp *desc.GetMMListResponse, err error) {

	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/mm", func(ctx context.Context) error {
		resp = &desc.GetMMListResponse{}

		res, err := i.admin.Actions.MMModule.GetMMList.Do(ctx)
		if err != nil {
			return err
		}

		resp.Mms = lo.Map(res, func(item dto.MM, index int) *desc.GetMMListResponse_Mm {
			return &desc.GetMMListResponse_Mm{
				MmId:       item.ID,
				MmDatetime: item.MMDatetime.Time.Format(time.DateTime),
				Name:       item.Name.String,
				CreatedAt:  item.CreatedAt.Format(time.DateTime),
				Status:     item.State.String,
				Activity:   item.IsActive.Bool,
				MmLink:     item.MMLink.String,
			}
		})

		return nil
	})
}
