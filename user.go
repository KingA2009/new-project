package todo

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Address   string `json:"address" binding:"required"`
}
type GetUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
}
