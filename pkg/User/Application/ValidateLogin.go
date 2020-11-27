package Application

import (
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func ValidateLogin(s Domain.UserService, u *Domain.User) (Domain.User,error){	

	user,err := s.ValidateLogin(u)
	if err != nil {
		return user,err
	}
	return user,nil
}
