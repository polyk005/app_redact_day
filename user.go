package todo

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
