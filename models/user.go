package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func NewUser(id int, name string, isActive bool) *User {
	return &User{
		Id:       id,
		Name:     name,
		IsActive: isActive,
	}
}
