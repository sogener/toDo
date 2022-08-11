package toDo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"requried"`
	Username string `json:"username" binding:"requried"`
	Password string `json:"password" binding:"requried"`
}
