package domain

type Repository interface {
	Create(course *Course) (Course, error)
	//Update()
	//Get()
	//List()
	//Delete()
}
