package demo

type Repository interface {
	Find(id uint) (bool, Demo, error)
	Save(demo Demo) error
}
