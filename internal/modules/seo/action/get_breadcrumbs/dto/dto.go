package dto

type Breadcrumb struct {
	Name string `db:"name"`
	Path string `db:"path"`
}

type Breadcrumbs []Breadcrumb
