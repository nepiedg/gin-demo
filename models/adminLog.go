package models

type AdminLog struct {
	Id         int
	AdminId    int
	Url        string
	Title      string
	Ip         string
	Content    string `gorm:"longtext"`
	Createtime int
}

func (AdminLog) TableName() string {
	return "fa_admin_log"
}
