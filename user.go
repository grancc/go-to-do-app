package gotodo

type User struct {
	Id       int    `json: "-"`
	Name     string `json: "name" binding:"required"`
	UserName string `json: "username" binding:"required"`
	Password string `json: "password_hash" binding:"required"`
}
