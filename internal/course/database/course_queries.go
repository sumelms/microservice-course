package database

const (
	createCourse       = "create course"
	deleteCourse       = "delete course by uuid"
	getCourse          = "get course by uuid"
	listCourse         = "list course"
	updateCourseByID   = "update course by uuid"
	updateCourseByCode = "update course by code"
)

func queriesCourse() map[string]string {
	return map[string]string{
		createCourse: `INSERT INTO 
    		courses (code, name, underline, image, image_cover, excerpt, description)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		deleteCourse: "UPDATE courses SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL",
		getCourse:    "SELECT * FROM courses WHERE uuid = $1 AND deleted_at IS NULL",
		listCourse:   "SELECT * FROM courses WHERE deleted_at IS NULL",
		updateCourseByID: `UPDATE courses 
			SET code = $1, name = $2, underline = $3, image = $4, image_cover = $5, excerpt = $6, description = $7 
			WHERE uuid = $8 AND deleted_at IS NULL RETURNING *`,
		updateCourseByCode: `UPDATE courses
			SET name = $1, underline = $2, image = $3, image_cover = $4, excerpt = $5, description = $6
			WHERE code = $7 AND deleted_at IS NULL RETURNING  *`,
	}
}
