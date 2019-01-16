package models

type UserInfo struct {
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	ProfilePic string `json:"profile_pic,omitempty"`
	Id         string `json:"id,omitempty"`
}
