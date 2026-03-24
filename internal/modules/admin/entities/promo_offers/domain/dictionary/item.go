package dictionary

type NamedItem struct {
	id   int64
	name string
}

type NamedItems []*NamedItem

func NewNamedItem(options ...NamedItemOption) *NamedItem {
	item := &NamedItem{}
	for _, option := range options {
		option(item)
	}

	return item
}

func (n *NamedItem) ID() int64 {
	return n.id
}

func (n *NamedItem) Name() string {
	return n.name
}

type SocialNetwork struct {
	id   int64
	name string
	slug string
}

type SocialNetworks []*SocialNetwork

func NewSocialNetwork(options ...SocialNetworkOption) *SocialNetwork {
	item := &SocialNetwork{}
	for _, option := range options {
		option(item)
	}

	return item
}

func (s *SocialNetwork) ID() int64 {
	return s.id
}

func (s *SocialNetwork) Name() string {
	return s.name
}

func (s *SocialNetwork) Slug() string {
	return s.slug
}

type BrandSocialNetwork struct {
	socialNetworkID int64
	name            string
	slug            string
	link            string
}

type BrandSocialNetworks []*BrandSocialNetwork

func NewBrandSocialNetwork(options ...BrandSocialNetworkOption) *BrandSocialNetwork {
	item := &BrandSocialNetwork{}
	for _, option := range options {
		option(item)
	}

	return item
}

func (b *BrandSocialNetwork) SocialNetworkID() int64 {
	return b.socialNetworkID
}

func (b *BrandSocialNetwork) Name() string {
	return b.name
}

func (b *BrandSocialNetwork) Slug() string {
	return b.slug
}

func (b *BrandSocialNetwork) Link() string {
	return b.link
}
