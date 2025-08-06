package helper

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"medblogers_base/internal/pkg/postgres"
)

var AddCity = func(ctx context.Context, pool postgres.PoolWrapper, cityName string) int64 {
	By("Создание нового города")

	sql := `insert into docstar_site_city (name) values ($1) returning id`

	args := []any{
		cityName,
	}

	var id int64
	err := pool.QueryRow(
		ctx,
		sql,
		args...,
	).Scan(&id)

	Expect(err).NotTo(HaveOccurred())
	return id
}

var AddAdditionalCity = func(ctx context.Context, pool postgres.PoolWrapper, cityID, doctorID int64) {
	By("Создание нового города")

	sql := `insert into docstar_site_doctor_additional_cities (doctor_id, city_id) values ($1, $2)`

	args := []any{
		doctorID,
		cityID,
	}

	_, err := pool.Exec(
		ctx,
		sql,
		args...,
	)

	Expect(err).NotTo(HaveOccurred())
}
