package Domain

type UserRepository interface {
	FindAllUser() ([]User, error)
	ValidateLogin(*User) (User,error)
}
