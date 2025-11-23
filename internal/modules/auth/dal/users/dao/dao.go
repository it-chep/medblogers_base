package dao

import "medblogers_base/internal/modules/auth/domain/user"

type UserDAO struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
	IsActive bool   `db:"is_active"`
	RoleID   int64  `db:"role_id"`
}

func (u UserDAO) ToDomain() *user.User {
	return user.New(
		user.WithID(u.ID),
		user.WithEmail(u.Email),
		user.WithPassword(u.Password),
		user.WithRoleID(u.RoleID),
		user.WithName(u.FullName),
	)
}
