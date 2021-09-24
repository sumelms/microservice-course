package domain

type Repository interface {
	Create(*Matrix) (Matrix, error)
	Find(string) (Matrix, error)
	Update(*Matrix) (Matrix, error)
	Delete(string) error
	List() ([]Matrix, error)
}
