package dictionary

type NamedItemOption func(*NamedItem)

func WithNamedItemID(id int64) NamedItemOption {
	return func(item *NamedItem) {
		item.id = id
	}
}

func WithNamedItemName(name string) NamedItemOption {
	return func(item *NamedItem) {
		item.name = name
	}
}

type SocialNetworkOption func(*SocialNetwork)

func WithSocialNetworkID(id int64) SocialNetworkOption {
	return func(item *SocialNetwork) {
		item.id = id
	}
}

func WithSocialNetworkName(name string) SocialNetworkOption {
	return func(item *SocialNetwork) {
		item.name = name
	}
}

func WithSocialNetworkSlug(slug string) SocialNetworkOption {
	return func(item *SocialNetwork) {
		item.slug = slug
	}
}

type BrandSocialNetworkOption func(*BrandSocialNetwork)

func WithBrandSocialNetworkID(id int64) BrandSocialNetworkOption {
	return func(item *BrandSocialNetwork) {
		item.socialNetworkID = id
	}
}

func WithBrandSocialNetworkName(name string) BrandSocialNetworkOption {
	return func(item *BrandSocialNetwork) {
		item.name = name
	}
}

func WithBrandSocialNetworkSlug(slug string) BrandSocialNetworkOption {
	return func(item *BrandSocialNetwork) {
		item.slug = slug
	}
}

func WithBrandSocialNetworkLink(link string) BrandSocialNetworkOption {
	return func(item *BrandSocialNetwork) {
		item.link = link
	}
}
