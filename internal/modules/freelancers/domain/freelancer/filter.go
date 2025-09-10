package freelancer

type Filter struct {
	Cities                []int64
	Specialities          []int64
	SocialNetworks        []int64
	PriceCategory         []int64
	ExperienceWithDoctors *bool

	Page int64
}
