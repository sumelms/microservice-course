package database

const (
	createCourse = "create course"
	deleteCourse = "delete course by uuid"
	getCourse    = "get course by uuid"
	listCourse   = "list course"
	updateCourse = "update course by uuid"
)

func courseQueries() map[string]string {
	return map[string]string{
		createCourse: "INSERT INTO courses (code, name, underline, image, image_cover, excerpt, description) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *",
		deleteCourse: "UPDATE courses SET deleted_at = NOW() WHERE uuid = $1",
		getCourse:    "SELECT * FROM courses WHERE uuid = $1",
		listCourse:   "SELECT * FROM courses",
		updateCourse: "UPDATE courses SET name = $1, underline = $2, image = $3, image_cover = $4, excerpt = $5, description = $6 WHERE uuid = $7 RETURNING *",
	}
}
