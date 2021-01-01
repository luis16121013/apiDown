package Domain

type UserService interface {
	FindAllUser() ([]User, error)
	FindIdUser(string) (User,error)
	ValidateLogin(*User) (User,error)
}

type service struct {
	ur UserRepository
}

//CREATE NEW SERVICE
func NewService(r UserRepository) UserService {
	return &service{ur: r}
}

//IMPLEMENT INTERFACE SERVICE
func (s *service) FindAllUser() ([]User, error) {
	return s.ur.FindAllUser()
}

func (s *service) FindIdUser(id string) (User,error){
	return s.ur.FindIdUser(id)
}

func (s *service) ValidateLogin(u *User) (User,error){
	return s.ur.ValidateLogin(u)
}

