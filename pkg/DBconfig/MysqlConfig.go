package DBconfig

import "fmt"

var (
	Username string = "usnbuus7qz634mfn"
	Password string = "eGPmxLkhfNLevfwyn6hW"
	Host     string = "bsifbrjg2n6llr3biozg-mysql.services.clever-cloud.com"
	Port     int    = 3306
	Database string = "bsifbrjg2n6llr3biozg"
)

func Url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", Username, Password, Host, Port, Database)
}
