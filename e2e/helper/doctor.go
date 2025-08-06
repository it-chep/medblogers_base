package helper

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
	"time"
)

var AddNewDoctor = func(ctx context.Context, pool postgres.PoolWrapper, doc *doctor.Doctor) int64 {
	By("Создание нового доктора")

	sql := `insert into docstar_site_doctor (
		                                 name, slug, email, inst_url, vk_url, dzen_url, tg_url, 
		                                 main_blog_theme, prodoctorov, city_id, speciallity_id, youtube_url, 
		                                 is_active, date_created, birth_date, tg_channel_url, tiktok_url
		                                 )
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, true, now(), $13, $14, $15) returning id;`

	args := []any{
		doc.GetName(), doc.GetSlug(),
		fmt.Sprintf("doctor%d@mail.com", rand.Intn(10000)),
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

	var id int64
	err := pool.QueryRow(
		ctx,
		sql,
		args...,
	).Scan(&id)

	Expect(err).NotTo(HaveOccurred())
	return id
}
