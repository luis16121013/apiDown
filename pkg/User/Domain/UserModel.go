package Domain

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	Email    string `json:"email"`
}
