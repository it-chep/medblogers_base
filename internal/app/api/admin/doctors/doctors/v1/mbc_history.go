package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
	"medblogers_base/internal/pkg/formatters"
)

func (i *Implementation) DoctorMBCHistory(ctx context.Context, req *desc.DoctorMBCHistoryRequest) (resp *desc.DoctorMBCHistoryResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{doctor_id}/mbc_history", func(ctx context.Context) error {
		items, err := i.admin.Actions.DoctorModule.DoctorAgg.MBCHistory.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		historyItems := make([]*desc.DoctorMBCHistoryResponse_MBCHistoryItem, 0, len(items))
		for _, item := range items {
			historyItems = append(historyItems, &desc.DoctorMBCHistoryResponse_MBCHistoryItem{
				MbcCount:   item.MBCCount,
				OccurredAt: formatters.TimeRuFormat(item.OccurredAt),
			})
		}

		resp = &desc.DoctorMBCHistoryResponse{
			Items: historyItems,
		}
		return nil
	})
}
