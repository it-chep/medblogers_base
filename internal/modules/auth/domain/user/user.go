package user

type User struct {
	email    string
	password string
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPassword() string {
	return u.password
}
