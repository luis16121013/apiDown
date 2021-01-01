package Infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luis16121013/apiDown/pkg/DBconfig"
	"log"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewRepository() (*MysqlRepository, error) {
	db, err := sql.Open("mysql", DBconfig.Url())
	if err != nil {
		return nil, err
	}
	return &MysqlRepository{db}, nil
}

//------------------------------------QUERY MYSQL
func (s *MysqlRepository) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

//------------------------------------EXEC MYSQL
func (s *MysqlRepository) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := s.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
