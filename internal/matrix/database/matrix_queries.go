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
		createMatrix:  "INSERT INTO matrices (name, description) VALUES (:name, :description) RETURNING *",
		deleteMatrix:  "UPDATE matrices SET deleted_at = NOW() WHERE uuid = :uuid",
		getMatrix:     "SELECT * FROM matrices WHERE uuid = :uuid",
		listMatrix:    "SELECT * FROM matrices",
		updateMatrix:  "UPDATE matrices SET name = :name, description = :description WHERE uuid = :uuid RETURNING *",
		addSubject:    "INSERT INTO matrix_subjects (matrix_id, subject_id, group) VALUES(:matrix_id, :subject_id, :group)",
		removeSubject: "UPDATE matrix_subjects SET deleted_at = NOW() WHERE matrix_id = :matrix_id AND subject_id = :subject_id",
	}
}
