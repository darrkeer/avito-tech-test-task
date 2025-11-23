package models

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewTeam(id int, name string) *Team {
	return &Team{
		Id:   id,
		Name: name,
	}
}
