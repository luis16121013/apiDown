package Infrastructure

import (
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func(s *MysqlRepository) ValidateLogin(u Domain.User) (Domain.User,error){
	sqlDB := "SELECT id,username,pwd,email FROM UserTable WHERE username=? AND pwd=?"
	rows,err := s.Query(sqlDB,u.Username,u.Password)
	if err != nil {
		return u,err
	}
	if rows.Next() {
		rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email)
	}
	return u,nil
}
