package helper

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"medblogers_base/internal/pkg/postgres"
)

var AddSpeciality = func(ctx context.Context, pool postgres.PoolWrapper, specialityName string) {
	By("Создание новой специальности")

	sql := `insert into docstar_site_speciallity (name) values ($1)`

	args := []any{
		specialityName,
	}

	_, err := pool.Exec(
		ctx,
		sql,
		args...,
	)

	Expect(err).NotTo(HaveOccurred())
}
