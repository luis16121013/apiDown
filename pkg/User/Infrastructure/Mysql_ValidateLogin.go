package Infrastructure

import (
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

func(s *MysqlRepository) ValidateLogin(u *Domain.User) (Domain.User,error){
	user := Domain.User{}
	sqlDB := "SELECT id,username,pwd,role,email FROM UserTable WHERE username=? AND pwd=?"
	rows,err := s.Query(sqlDB,u.Username,u.Password)
	if err != nil {
		return user,err
	}
	if rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Email)
	}
	return user,nil
}
