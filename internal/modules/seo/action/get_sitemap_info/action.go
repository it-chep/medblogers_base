package get_sitemap_info

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/seo/action/get_sitemap_info/dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Repository interface {
	GetAllDoctorsSlugs(ctx context.Context) ([]string, error)
	GetAllFreelancersSlugs(ctx context.Context) ([]string, error)
	GetAllBlogsSlugs(ctx context.Context) ([]string, error)
}

type Action struct {
	repository Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repository: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]string, error) {
	sitemap := make([]string, 0, 1000)
	sitemap = append(
		sitemap, []string{
			"/welcome", "/blogs", "/helpers", "/helpers_welcome",
			"/new_club_participant", "/new_freelancer",
		}...,
	)

	sitemap = append(sitemap, a.getDoctorsURLs(ctx)...)
	sitemap = append(sitemap, a.getFreelancersURLs(ctx)...)
	sitemap = append(sitemap, a.getBlogsURLs(ctx)...)

	return sitemap, nil
}

func (a *Action) getDoctorsURLs(ctx context.Context) []string {
	doctorsSlugs, err := a.repository.GetAllDoctorsSlugs(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получения sitemap докторов", err)
	}
	return lo.Map(doctorsSlugs, func(item string, _ int) string {
		return "/doctors/" + item
	})
}

func (a *Action) getFreelancersURLs(ctx context.Context) []string {
	freelancersSlugs, err := a.repository.GetAllFreelancersSlugs(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получения sitemap докторов", err)
	}
	return lo.Map(freelancersSlugs, func(item string, _ int) string {
		return "/helpers/" + item
	})
}

func (a *Action) getBlogsURLs(ctx context.Context) []string {
	blogsSlugs, err := a.repository.GetAllBlogsSlugs(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получения sitemap докторов", err)
	}
	return lo.Map(blogsSlugs, func(item string, _ int) string {
		return "/blog/" + item
	})
}
