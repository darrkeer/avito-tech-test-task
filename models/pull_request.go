package models

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusMerged Status = "MERGED"
)

type PullRequest struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Status Status `json:"pr_status"`
}

func NewPullRequest(id int, user_id int, title string, status Status) *PullRequest {
	return &PullRequest{
		Id:     id,
		UserId: user_id,
		Title:  title,
		Status: status,
	}
}
