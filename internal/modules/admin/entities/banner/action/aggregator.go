package action

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/banner/action/activate"
	"medblogers_base/internal/modules/admin/entities/banner/action/create"
	"medblogers_base/internal/modules/admin/entities/banner/action/deactivate"
	"medblogers_base/internal/modules/admin/entities/banner/action/get"
	"medblogers_base/internal/modules/admin/entities/banner/action/get_by_id"
	"medblogers_base/internal/modules/admin/entities/banner/action/save_desktop_image"
	"medblogers_base/internal/modules/admin/entities/banner/action/save_mobile_image"
	"medblogers_base/internal/modules/admin/entities/banner/action/update"
	"medblogers_base/internal/pkg/postgres"
)

type Aggregator struct {
	Get              *get.Action
	GetByID          *get_by_id.Action
	Create           *create.Action
	Update           *update.Action
	Activate         *activate.Action
	Deactivate       *deactivate.Action
	SaveDesktopImage *save_desktop_image.Action
	SaveMobileImage  *save_mobile_image.Action
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		Get:              get.New(pool),
		GetByID:          get_by_id.New(pool),
		Create:           create.New(pool),
		Update:           update.New(pool),
		Activate:         activate.New(pool),
		Deactivate:       deactivate.New(pool),
		SaveDesktopImage: save_desktop_image.New(clients, pool),
		SaveMobileImage:  save_mobile_image.New(clients, pool),
	}
}
