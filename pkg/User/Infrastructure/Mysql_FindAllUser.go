package Infrastructure

import (
	"errors"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

type Usr []Domain.User

func (s *MysqlRepository) FindAllUser() ([]Domain.User, error) {
	users := Usr{}
	sqlDB := "SELECT id,username,pwd,role,email FROM UserTable"
	rows, err := s.Query(sqlDB)
	if err != nil {
		return users, errors.New("User no found!")
	}

	for rows.Next() {
		u := Domain.User{}
		rows.Scan(&u.Id, &u.Username, &u.Password, &u.Role, &u.Email)
		users = append(users, u)
	}
	return users, nil
}
