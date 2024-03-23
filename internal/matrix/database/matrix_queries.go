package database

const (
	// CREATE.
	createMatrix = "create matrix"
	addSubject   = "adds subject to matrix"
	// READ.
	getMatrix       = "get matrix by uuid"
	getCourseMatrix = "get course and matrix by uuid"
	listMatrix      = "list matrices"
	// UPDATE.
	updateMatrix = "update matrix by uuid"
	// DELERE
	deleteMatrix  = "delete matrix by uuid"
	removeSubject = "remove subject from matrix"
)

func queriesMatrix() map[string]string {
	return map[string]string{
		// CREATE.
		createMatrix: "INSERT INTO matrices (name, description) VALUES ($1, $2) RETURNING *",
		addSubject:   "INSERT INTO matrix_subjects (matrix_id, subject_id) VALUES($1, $2)",
		// READ.
		getMatrix: "SELECT * FROM matrices WHERE uuid = $1",
		getCourseMatrix: `SELECT *
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
				AND matrices.uuid = $2 AND matrices.deleted_at IS NULL`,
		listMatrix: "SELECT * FROM matrices",
		// UPDATE.
		updateMatrix: "UPDATE matrices SET name = $1, description = $2 WHERE uuid = $3 RETURNING *",
		// DELETE.
		deleteMatrix:  "UPDATE matrices SET deleted_at = NOW() WHERE uuid = $1",
		removeSubject: "UPDATE matrix_subjects SET deleted_at = NOW() WHERE matrix_id = $1 AND subject_id = $2",
	}
}
