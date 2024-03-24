package database

import "fmt"

const (
	returningMatrixColumns = `matrices.uuid, matrices.code, matrices.name, matrices.description,
		matrices.created_at, matrices.updated_at`
	// CREATE.
	createMatrix = "create matrix"
	addSubject   = "adds subject to matrix"
	// READ.
	getMatrix          = "get matrix by uuid"
	getCourseMatrix    = "get course and matrix by course_uuid"
	listMatrix         = "list matrices"
	listCourseMatrices = "list course matrices by course_uuid"
	// UPDATE.
	updateMatrix = "update matrix by uuid"
	// DELERE.
	deleteMatrix  = "delete matrix by uuid"
	removeSubject = "remove subject from matrix"
)

func queriesMatrix() map[string]string {
	return map[string]string{
		// CREATE.
		createMatrix: fmt.Sprintf(`INSERT INTO matrices (
				course_id, code, name, description
			) SELECT courses.id, $2, $3, $4
				FROM courses
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
			RETURNING %s`, returningMatrixColumns),
		addSubject: "INSERT INTO matrix_subjects (matrix_id, subject_id) VALUES($1, $2)",
		// READ.
		getMatrix: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				%s
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE matrices.uuid = $1 AND matrices.deleted_at IS NULL
				AND courses.deleted_at IS NULL`, returningMatrixColumns),
		getCourseMatrix: fmt.Sprintf(`SELECT
				%s
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
				AND matrices.uuid = $2 AND matrices.deleted_at IS NULL`, returningMatrixColumns),
		listMatrix: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				%s
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE matrices.deleted_at IS NULL
				AND courses.deleted_at IS NULL`, returningMatrixColumns),
		listCourseMatrices: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				%s
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
				AND matrices.deleted_at IS NULL`, returningMatrixColumns),
		// UPDATE.
		updateMatrix: fmt.Sprintf(`UPDATE matrices
			SET code = $2, name = $3, description = $4, updated_at = NOW()
			WHERE uuid = $1
			RETURNING %s`, returningMatrixColumns),
		// DELETE.
		deleteMatrix: fmt.Sprintf(`UPDATE matrices
			SET deleted_at = NOW()
			WHERE uuid = $1
			RETURNING %s`, returningMatrixColumns),
		removeSubject: `UPDATE matrix_subjects
			SET deleted_at = NOW()
			WHERE matrix_id = $1 AND subject_id = $2`,
	}
}
