package Domain

type UserRepository interface {
	FindAllUser() ([]User, error)
	FindIdUser(string) (User,error)
	ValidateLogin(*User) (User,error)
}
