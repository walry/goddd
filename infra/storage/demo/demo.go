package demo

import (
	"errors"
	"gorm.io/gorm"
	"youras/domain/aggregate/demo"
	"youras/infra/storage/model"
)

type Demo struct {
	db *gorm.DB
}

func NewDemoStorage(db *gorm.DB) *Demo {
	return &Demo{db: db}
}

func (d *Demo) Find(id uint) (bool, demo.Demo, error) {
	var data model.Demo
	err := d.db.Table("demo").Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, demo.Demo{}, nil
		}
		return false, demo.Demo{}, err
	}
	return true, demo.Demo{Id: data.Id, Name: data.Name}, nil
}

func (d *Demo) Save(demo demo.Demo) error {
	data := &model.Demo{Id: demo.Id, Name: demo.Name}
	err := d.db.Table("demo").FirstOrCreate(&data).Error
	return err
}
