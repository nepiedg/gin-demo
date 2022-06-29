package models

type User struct {
	ID       int16
	GroupId  int
	Username string
	Nickname string
	Password string
}

func (User) TableName() string {
	return "fa_user"
}
