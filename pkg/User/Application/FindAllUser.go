package Application

import (
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func FindAllUser(s Domain.UserService) ([]Domain.User, error) {
	users, err := s.FindAllUser()
	if err != nil {
		return users, err
	}
	return users, nil
}
