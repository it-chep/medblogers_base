package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/network/v1"
)

func (i *Implementation) GetNetworks(ctx context.Context, req *desc.GetNetworksRequest) (resp *desc.GetNetworksResponse, err error) {
	return &desc.GetNetworksResponse{}, nil
}
