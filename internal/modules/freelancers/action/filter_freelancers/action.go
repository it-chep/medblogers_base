package filter_freelancers

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
)

type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context, filter freelancer.Filter) ([]*freelancer.Freelancer, error) {

}
