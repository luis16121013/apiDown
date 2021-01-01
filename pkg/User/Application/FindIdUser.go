package Application

import (
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func FindIdUser(s Domain.UserService, id string) (Domain.User,error){	

	user,err := s.FindIdUser(id)
	if err != nil {
		return user,err
	}
	return user,nil
}