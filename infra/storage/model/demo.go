package model

type Demo struct {
	Id   uint `gorm:"primary_key"`
	Name string
}

func (Demo) TableName() string {
	return "demo"
}
