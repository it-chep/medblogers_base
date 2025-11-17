package user

type User struct {
	id       int64
	email    string
	password string
	fullName string
	isActive bool
	roleID   int64
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetRoleID() int64 {
	return u.roleID
}

func New(options ...Option) *User {
	d := &User{}
	for _, option := range options {
		option(d)
	}
	return d
}
