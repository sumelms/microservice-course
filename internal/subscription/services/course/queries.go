package course

const (
	existCourse = "exist course"
)

type Query string

func queries() map[string]Query {
	return map[string]Query{
		existCourse: "SELECT id, uuid FROM courses WHERE uuid = $1",
	}
}
