package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/modules/blogs/domain/blog_photo"
	"medblogers_base/internal/modules/blogs/domain/doctor_author"
	"time"

	"github.com/google/uuid"
)

type BlogDAO struct {
	ID    uuid.UUID `db:"id"`
	Title string    `db:"name"`
	Slug  string    `db:"slug"`

	Body              sql.NullString `db:"body"`
	PreviewText       sql.NullString `db:"preview_text"`
	SocietyPreview    sql.NullString `db:"society_preview"`
	AdditionalSEOText sql.NullString `db:"additional_seo_text"`
	CreatedAt         time.Time      `db:"created_at"`
	OrderingNumber    sql.NullInt64  `db:"ordering_number"`
	DoctorID          sql.NullInt64  `db:"doctor_id"`
}

func (b *BlogDAO) ToDomain() *blog.Blog {
	return blog.New(
		blog.WithID(b.ID),
		blog.WithTitle(b.Title),
		blog.WithSlug(b.Slug),
		blog.WithBody(b.Body.String),
		blog.WithPreviewText(b.PreviewText.String),
		blog.WithSocietyPreviewText(b.SocietyPreview.String),
		blog.WithAdditionalSEOText(b.AdditionalSEOText.String),
		blog.WithCreatedAt(b.CreatedAt),
		blog.WithOrderingNumber(b.OrderingNumber.Int64),
		blog.WithDoctorID(b.DoctorID.Int64),
	)
}

type DoctorAuthorDAO struct {
	Name           string `db:"name"`
	Slug           string `db:"slug"`
	S3Key          string `db:"s3_key"`
	SpecialityName string `db:"speciality_name"`
}

func (d *DoctorAuthorDAO) ToDomain() *doctor_author.Doctor {
	return doctor_author.NewDoctor(
		d.Name, d.Slug, d.S3Key, d.SpecialityName,
	)
}

type BlogMiniatureDAO struct {
	ID    uuid.UUID `db:"id"`
	Title string    `db:"name"`
	Slug  string    `db:"slug"`

	PreviewText    sql.NullString `db:"preview_text"`
	CreatedAt      time.Time      `db:"created_at"`
	OrderingNumber sql.NullInt64  `db:"ordering_number"`
}

type BlogMiniatureDAOs []*BlogMiniatureDAO

func (b *BlogMiniatureDAO) ToDomain() *blog.Blog {
	return blog.New(
		blog.WithID(b.ID),
		blog.WithTitle(b.Title),
		blog.WithSlug(b.Slug),
		blog.WithPreviewText(b.PreviewText.String),
		blog.WithCreatedAt(b.CreatedAt),
		blog.WithOrderingNumber(b.OrderingNumber.Int64),
	)
}

func (b BlogMiniatureDAOs) ToDomain() blog.Blogs {
	domains := make(blog.Blogs, 0, len(b))
	for _, dao := range b {
		domains = append(domains, dao.ToDomain())
	}
	return domains
}

type PrimaryPhotoDAO struct {
	ID        uuid.UUID      `db:"id"`
	BlogID    uuid.UUID      `db:"blog_id"`
	FileType  sql.NullString `db:"file_type"`
	IsPrimary bool           `db:"is_primary"`
}

type PrimaryPhotoDAOs []*PrimaryPhotoDAO

// ToDomain .
func (d PrimaryPhotoDAO) ToDomain() *blog_photo.BlogPhoto {
	return blog_photo.New(
		blog_photo.WithBlogID(d.BlogID),
		blog_photo.WithIsPrimary(d.IsPrimary),
		blog_photo.WithFileType(d.FileType.String),
		blog_photo.WithPhotoID(d.ID),
	)
}
