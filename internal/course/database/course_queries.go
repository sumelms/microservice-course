package database

const (
	createCourse = "create course"
	deleteCourse = "delete course by uuid"
	getCourse    = "get course by uuid"
	listCourse   = "list course"
	updateCourse = "update course by uuid"
)

func queriesCourse() map[string]string {
	return map[string]string{
		createCourse: `INSERT INTO 
    		courses (code, name, underline, image, image_cover, excerpt, description)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		deleteCourse: "UPDATE courses SET deleted_at = NOW() WHERE uuid = $1",
		getCourse:    "SELECT * FROM courses WHERE uuid = $1",
		listCourse:   "SELECT * FROM courses",
		updateCourse: `UPDATE courses 
			SET code = $1, name = $2, underline = $3, image = $4, image_cover = $5, excerpt = $6, description = $7 
			WHERE uuid = $8 RETURNING *`,
	}
}
