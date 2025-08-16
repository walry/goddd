package service

import (
	"fmt"
	adto "youras/application/dto"
	"youras/domain/aggregate/demo"
	"youras/infra/status"
)

type DemoService struct {
	repo demo.Repository
}

func (d *DemoService) Get(id int) (demo.Demo, error) {
	has, data, err := d.repo.Find(uint(id))
	if err != nil {
		return demo.Demo{}, status.ErrDbOpt
	}
	if !has {
		return demo.Demo{}, nil
	}
	return data, nil
}

func NewDemoService(repo demo.Repository) *DemoService {
	return &DemoService{repo: repo}
}

func (d *DemoService) UpdateName(cmd adto.UpdateDemoCommand) error {
	has, data, err := d.repo.Find(cmd.Id)
	if err != nil {
		return fmt.Errorf("update err: %w", err)
	}
	if !has {
		return fmt.Errorf("id not found")
	}
	data.Name = cmd.Name
	err = d.repo.Save(data)
	if err != nil {
		return fmt.Errorf("update err: %w", err)
	}
	return nil
}
