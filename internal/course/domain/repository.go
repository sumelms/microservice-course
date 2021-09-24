package domain

type Repository interface {
	Create(*Course) (Course, error)
	Find(string) (Course, error)
	Update(*Course) (Course, error)
	Delete(string) error
	List() ([]Course, error)
}
