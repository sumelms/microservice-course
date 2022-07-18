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
		createCourse: Query("INSERT INTO courses (code, name, underline, image, image_cover, excerpt, description) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"),
		deleteCourse: Query("UPDATE courses SET deleted_at = NOW() WHERE uuid = $1"),
		getCourse:    Query("SELECT * FROM courses WHERE uuid = $1"),
		listCourse:   Query("SELECT * FROM courses"),
		updateCourse: Query("UPDATE courses SET name = $1, underline = $2, image = $3, image_cover = $4, excerpt = $5, description = $6 WHERE uuid = $7 RETURNING *"),
	}
}
