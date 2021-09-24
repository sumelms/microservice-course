package domain

type Repository interface {
	Create()
	Find()
	Update()
	Delete()
	List()
}
