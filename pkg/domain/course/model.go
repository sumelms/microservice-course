package course

type Course struct {
	ID          uint
	UUID        string
	Name        string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   *int64
	PublishedAt *int64
}
