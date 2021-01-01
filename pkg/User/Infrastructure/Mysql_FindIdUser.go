package Infrastructure

import (
	"errors"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func (s *MysqlRepository) FindIdUser(id string) (Domain.User, error) {
	user := Domain.User{}
	sqlDB := "SELECT id,username,pwd,role,email FROM UserTable WHERE id=?"
	rows, err := s.Query(sqlDB,id)
	if err != nil {
		return user, errors.New("User not found!")
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Email)
	}
	return user, nil
}