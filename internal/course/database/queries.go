package database

const (
	createCourse = "create course"
	deleteCourse = "delete course by uuid"
	getCourse    = "get course by uuid"
	listCourse   = "list course"
	updateCourse = "update course by uuid"
)

type Query string

func queries() map[string]Query {
	return map[string]Query{
		createCourse: Query("INSERT INTO courses (title, subtitle, excerpt, description) VALUES ($1, $2, $3, $4) RETURNING *"),
		deleteCourse: Query("UPDATE courses SET deleted_at = NOW() WHERE uuid = $1"),
		getCourse:    Query("SELECT * FROM courses WHERE uuid = $1"),
		listCourse:   Query("SELECT * FROM courses"),
		updateCourse: Query("UPDATE courses SET title = $1, subtitle = $2, excerpt = $3, description = $4 WHERE uuid = $5 RETURNING *"),
	}
}
