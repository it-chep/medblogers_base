package helper

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"medblogers_base/internal/pkg/postgres"
)

var AddCity = func(ctx context.Context, pool postgres.PoolWrapper, cityName string) {
	By("Создание нового города")

	sql := `insert into docstar_site_city (name) values ($1)`

	args := []any{
		cityName,
	}

	_, err := pool.Exec(
		ctx,
		sql,
		args...,
	)

	Expect(err).NotTo(HaveOccurred())
}
