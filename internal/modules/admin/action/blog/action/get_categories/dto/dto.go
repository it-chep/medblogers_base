package dto

type BlogCategory struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Categories []BlogCategory
