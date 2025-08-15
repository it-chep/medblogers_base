package helper

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var AddSpeciality = func(ctx context.Context, pool postgres.PoolWrapper, specialityName string) int64 {
	By("Создание новой специальности")

	sql := `insert into docstar_site_speciallity (name) values ($1) returning id;`

	args := []any{
		specialityName,
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

var AddAdditionalSpeciality = func(ctx context.Context, pool postgres.PoolWrapper, specialityID, doctorID int64) {
	By("Создание дополнительной специальности врачу")

	sql := `insert into docstar_site_doctor_additional_specialties (doctor_id, speciallity_id) values ($1, $2)`

	args := []any{
		doctorID,
		specialityID,
	}

	_, err := pool.Exec(
		ctx,
		sql,
		args...,
	)

	Expect(err).NotTo(HaveOccurred())
}
