//go:build e2e
// +build e2e

package doctor

import (
	"context"
	"medblogers_base/e2e/helper"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/e2e/module/doctor/fixture"
	"medblogers_base/e2e/shared_behavior"
	respDTO "medblogers_base/internal/modules/doctors/action/settings/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/dto"
	"medblogers_base/internal/modules/doctors/domain/doctor"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("Пользователь запрашивает список доступных фильтров", Label("rega", "get_all_cities"), func() {
	var (
		module         *app.TestableModule
		additionalCity int64
		doctorID       int64
	)

	BeforeEach(func(ctx context.Context) {
		module = fixture.SetupModule(ctx, conn)
		cityID := helper.AddCity(ctx, conn, "Москва")
		specialityID := helper.AddSpeciality(ctx, conn, "Хирург")

		doctorID = helper.AddNewDoctor(ctx, module.PoolWrapper, doctor.New(
			doctor.WithName("Мага"),
			doctor.WithSlug("maga"),
			doctor.WithMainCityID(cityID),
			doctor.WithMainSpecialityID(specialityID),
		))
		additionalCity = helper.AddCity(ctx, conn, "Санкт-петербург")
		// Основной и дополнительные города лежат в таблице additional_cities
		helper.AddAdditionalCity(ctx, conn, additionalCity, doctorID)
		helper.AddAdditionalCity(ctx, conn, cityID, doctorID)

		helper.AddAdditionalSpeciality(ctx, conn, specialityID, additionalCity)
	})

	It("Успешное получение всех фильтров", func(ctx context.Context) {
		shared_behavior.ExpectGetFilterInfoSuccess(module.Http, dto.FilterInfoResponse{
			Messengers: []dto.FilterInfo{
				{Name: "Telegram", Slug: "tg"},
			},
		})
		settingsInfo, err := module.Module.Actions.Settings.Do(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(settingsInfo.Cities)).To(BeNumerically(">", 0))
		Expect(len(settingsInfo.Specialities)).To(BeNumerically(">", 0))
		Expect(settingsInfo.FilterInfo).Should(Equal([]respDTO.FilterItem{
			{Name: "Telegram", Slug: "tg"},
		}))
	})

	It("Ошибка subscribers нет фильтров по подписчикам", func(ctx context.Context) {
		shared_behavior.ExpectGetFilterInfoError(module.Http, errors.New("e2e test err"))
		settingsInfo, err := module.Module.Actions.Settings.Do(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(settingsInfo.Cities)).To(BeNumerically(">", 0))
		Expect(len(settingsInfo.Specialities)).To(BeNumerically(">", 0))
		Expect(len(settingsInfo.FilterInfo)).To(Equal(0))
	})

	It("Успешное получение всех фильтров. Есть доп города", func(ctx context.Context) {
		shared_behavior.ExpectGetFilterInfoSuccess(module.Http, dto.FilterInfoResponse{
			Messengers: []dto.FilterInfo{
				{Name: "Telegram", Slug: "tg"},
			},
		})
		settingsInfo, err := module.Module.Actions.Settings.Do(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(settingsInfo.Cities)).To(BeNumerically(">", 0))
		Expect(len(settingsInfo.Specialities)).To(BeNumerically(">", 0))
		Expect(settingsInfo.FilterInfo).Should(Equal([]respDTO.FilterItem{
			{Name: "Telegram", Slug: "tg"},
		}))
	})
})
