package freelancer

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/activate_freelancer"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_additional_city"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_additional_speciality"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_network"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_price_list"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/add_recommendation"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/deactivate_freelancer"
	delete_freelancer "medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_additional_city"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_additional_speciality"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_network"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_price_list"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/delete_recommendation"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_additional_cities"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_additional_specialities"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_by_id"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_price_list"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_recommendations"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_social_networks"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/save_freelancer_photo"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/search"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update"
	"medblogers_base/internal/pkg/postgres"
)

type FreelancerAggregator struct {
	Activate   *activate_freelancer.Action
	Deactivate *deactivate_freelancer.Action

	AddAdditionalCity    *add_additional_city.Action
	DeleteAdditionalCity *delete_additional_city.Action

	AddAdditionalSpeciality    *add_additional_speciality.Action
	DeleteAdditionalSpeciality *delete_additional_speciality.Action

	AddNetwork    *add_network.Action
	DeleteNetwork *delete_network.Action

	AddPriceList    *add_price_list.Action
	DeletePriceList *delete_price_list.Action

	AddRecommendation    *add_recommendation.Action
	DeleteRecommendation *delete_recommendation.Action

	GetFreelancers      *get.Action
	GetFreelancerByID   *get_by_id.Action
	SaveFreelancerPhoto *save_freelancer_photo.Action
	SearchFreelancers   *search.Action
	UpdateFreelancer    *update.Action
	DeleteFreelancer    *delete_freelancer.Action

	GetPriceList              *get_price_list.Action
	GetRecommendations        *get_recommendations.Action
	GetAdditionalCities       *get_additional_cities.Action
	GetNetworks               *get_social_networks.Action
	GetAdditionalSpecialities *get_additional_specialities.Action
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *FreelancerAggregator {
	return &FreelancerAggregator{
		Activate:   activate_freelancer.New(pool),
		Deactivate: deactivate_freelancer.New(pool),

		AddAdditionalCity:    add_additional_city.New(pool),
		DeleteAdditionalCity: delete_additional_city.New(pool),

		AddAdditionalSpeciality:    add_additional_speciality.New(pool),
		DeleteAdditionalSpeciality: delete_additional_speciality.New(pool),

		AddNetwork:    add_network.New(pool),
		DeleteNetwork: delete_network.New(pool),

		AddPriceList:    add_price_list.New(pool),
		DeletePriceList: delete_price_list.New(pool),

		AddRecommendation:    add_recommendation.New(pool),
		DeleteRecommendation: delete_recommendation.New(pool),

		GetFreelancers:      get.New(pool),
		GetFreelancerByID:   get_by_id.New(clients, pool),
		SaveFreelancerPhoto: save_freelancer_photo.New(clients, pool),
		SearchFreelancers:   search.New(pool),
		UpdateFreelancer:    update.New(),
		DeleteFreelancer:    delete_freelancer.New(pool),

		GetPriceList:              get_price_list.New(pool),
		GetRecommendations:        get_recommendations.New(pool),
		GetAdditionalCities:       get_additional_cities.New(pool),
		GetNetworks:               get_social_networks.New(pool),
		GetAdditionalSpecialities: get_additional_specialities.New(pool),
	}
}
