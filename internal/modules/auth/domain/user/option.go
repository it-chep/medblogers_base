package user

// Option .
type Option func(u *User)

// WithID .
func WithID(id int64) Option {
	return func(u *User) {
		u.id = id
	}
}

// WithEmail .
func WithEmail(email string) Option {
	return func(u *User) {
		u.email = email
	}
}

// WithPassword .
func WithPassword(pass string) Option {
	return func(u *User) {
		u.password = pass
	}
}

// WithRoleID .
func WithRoleID(roleID int64) Option {
	return func(u *User) {
		u.roleID = roleID
	}
}

// WithName .
func WithName(name string) Option {
	return func(u *User) {
		u.fullName = name
	}
}
