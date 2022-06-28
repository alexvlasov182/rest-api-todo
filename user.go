package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `josn:"username"`
	Password string `json:"password"`
}
