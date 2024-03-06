package database

import (
	"fmt"
)

const (
	returning_columns = "uuid, code, name, underline, image, image_cover, excerpt, description, created_at, updated_at"
	// CREATE.
	createCourse = "create course"
	// READ.
	listCourse = "list course"
	getCourse  = "get course by uuid"
	// UPDATE.
	updateCourseByUUID = "update course by uuid"
	// DELETE.
	deleteCourse = "delete course by uuid"
)

func queriesCourse() map[string]string {
	return map[string]string{
		// CREATE.
		createCourse: fmt.Sprintf(`INSERT INTO courses (
				code, name, underline, image, image_cover, excerpt, description
			) VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING %s`, returning_columns),
		// READ.
		listCourse: fmt.Sprintf("SELECT %s FROM courses WHERE deleted_at IS NULL", returning_columns),
		getCourse:  fmt.Sprintf("SELECT %s FROM courses WHERE uuid = $1 AND deleted_at IS NULL", returning_columns),
		// UPDATE.
		updateCourseByUUID: fmt.Sprintf(`UPDATE courses SET
				code = $1, name = $2, underline = $3, image = $4, image_cover = $5, excerpt = $6, description = $7 
			WHERE uuid = $8
				AND deleted_at IS NULL
			RETURNING %s`, returning_columns),
		// DELETE.
		deleteCourse: "UPDATE courses SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL",
	}
}
