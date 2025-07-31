package helper

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
	"time"
)

var AddNewDoctor = func(ctx context.Context, pool postgres.PoolWrapper, doc *doctor.Doctor) *doctor.Doctor {
	By("Создание нового доктора")

	newDoctor := dao.DoctorDAO{}

	sql := `insert into docstar_site_doctor (
		                                 name, slug, email, inst_url, vk_url, dzen_url, tg_url, 
		                                 main_blog_theme, prodoctorov, city_id, speciallity_id, youtube_url, 
		                                 is_active, date_created, birth_date, tg_channel_url, tiktok_url
		                                 )
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, true, now(), $13, $14, $15)`

	args := []any{
		doc.GetName(), doc.GetSlug(),
		"doctor@mail.com",
		doc.GetInstURL(),
		doc.GetVkURL(),
		doc.GetDzenURL(),
		doc.GetTgURL(),
		doc.GetMainBlogTheme(),
		doc.GetSiteLink(),
		doc.GetMainCityID(),
		doc.GetMainSpecialityID(),
		doc.GetYoutubeURL(),
		time.Now(),
		doc.GetTgChannelURL(),
		doc.GetTiktokURL(),
	}

	rows, err := pool.Exec(
		ctx,
		sql,
		args...,
	)

	Expect(rows.RowsAffected()).To(Equal(int64(1)))
	Expect(err).NotTo(HaveOccurred())
	return newDoctor.ToDomain()
}
