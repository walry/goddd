package factory

import (
	"gorm.io/gorm"
	"youras/application/service"
	"youras/infra/storage/demo"
)

type Factory struct {
	db *gorm.DB
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{db: db}
}

func (f *Factory) CreateDemoService() *service.DemoService {
	return service.NewDemoService(demo.NewDemoStorage(f.db))
}
