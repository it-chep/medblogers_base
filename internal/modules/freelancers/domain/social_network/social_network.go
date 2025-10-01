package social_network

// SocialNetwork - справочик соц сетей
type SocialNetwork struct {
	id               int64
	name             string
	slug             string
	freelancersCount int64
}

type Networks []*SocialNetwork

// BuildSocialNetwork создать соц сеть
func BuildSocialNetwork(options ...Option) *SocialNetwork {
	e := &SocialNetwork{}
	for _, option := range options {
		option(e)
	}
	return e
}

func (c *SocialNetwork) ID() int64 {
	return c.id
}

func (c *SocialNetwork) Name() string {
	return c.name
}

func (c *SocialNetwork) Slug() string {
	return c.slug
}

func (c *SocialNetwork) FreelancersCount() int64 {
	return c.freelancersCount
}
