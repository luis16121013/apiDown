package Domain

type UserService interface {
	FindAllUser() ([]User, error)
	ValidateLogin(User) (User,error)
}

type service struct {
	ur UserRepository
}

func NewService(r UserRepository) UserService {
	return &service{ur: r}
}

//IMPLEMENT INTERFACE SERVICE
func (s *service) FindAllUser() ([]User, error) {
	return s.ur.FindAllUser()
}

func (s *service) ValidateLogin(u User) (User,error){
	return s.ur.ValidateLogin(u)
}
