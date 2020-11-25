package Domain

type UserRepository interface {
	FindAllUser() ([]User, error)
	ValitadeLogin(User) (User,error)
}
