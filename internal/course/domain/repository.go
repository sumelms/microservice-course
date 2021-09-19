package domain

type Repository interface {
	Create(course *Course) (Course, error)
	Find(course string) (Course, error)
	Update(course *Course) (Course, error)
	//List()
	//Delete()
}
