package database

const (
	createMatrix  = "create matrix"
	deleteMatrix  = "delete matrix by uuid"
	getMatrix     = "get matrix by uuid"
	listMatrix    = "list matrices"
	updateMatrix  = "update matrix by uuid"
	addSubject    = "adds subject to matrix"
	removeSubject = "remove subject from matrix"
)

func queriesMatrix() map[string]string {
	return map[string]string{
		createMatrix:  "INSERT INTO matrices (name, description) VALUES ($1, $2) RETURNING *",
		deleteMatrix:  "UPDATE matrices SET deleted_at = NOW() WHERE uuid = $1",
		getMatrix:     "SELECT * FROM matrices WHERE uuid = $1",
		listMatrix:    "SELECT * FROM matrices",
		updateMatrix:  "UPDATE matrices SET name = $1, description = $2 WHERE uuid = $3 RETURNING *",
		addSubject:    "INSERT INTO matrix_subjects (matrix_id, subject_id) VALUES($1, $2)",
		removeSubject: "UPDATE matrix_subjects SET deleted_at = NOW() WHERE matrix_id = $1 AND subject_id = $2",
	}
}
