package domain

type Repository interface {
	Create(course *Course) (Course, error)
	Find(course string) (Course, error)
	Update(course *Course) (Course, error)
	Delete(course string) error
	//List()
}
