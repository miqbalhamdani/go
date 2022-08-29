package services

type User struct {
	Nama string
}

type userService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) *userService {
	return &userService{
		db: db,
	}
}

func (u *userService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " success to register"
}

func (u *userService) GetUser() []*User {
	return u.db
}
