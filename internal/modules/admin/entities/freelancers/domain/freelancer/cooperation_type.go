package freelancer

type CooperationType struct {
	id   int64
	name string
}

func NewCooperationType(id int64, name string) *CooperationType {
	return &CooperationType{
		id:   id,
		name: name,
	}
}

func (r *CooperationType) ID() int64 {
	return r.id
}

func (r *CooperationType) Name() string {
	return r.name
}
