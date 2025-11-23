package models

// type ReadMember = User

type ReadMember struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type ReadTeam struct {
	TeamName string       `json:"team_name"`
	Members  []ReadMember `json:"members"`
}
