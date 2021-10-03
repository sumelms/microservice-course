package domain

type Repository interface {
	Create(*Matrix) (Matrix, error)
	Find(string) (Matrix, error)
	Update(*Matrix) (Matrix, error)
	Delete(string) error
	List(map[string]interface{}) ([]Matrix, error)
	FindBy(string, interface{}) ([]Matrix, error)
}
