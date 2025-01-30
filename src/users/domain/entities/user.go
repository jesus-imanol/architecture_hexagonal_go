package entities

type User struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Role int32 `json:"role"`
}

func NewUser(name string, lastName string, password string,role int32) *User {
	return &User{Name: name, LastName: lastName, Password: password, Role: role}
}