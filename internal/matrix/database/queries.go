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

type Query string

func queries() map[string]Query {
	return map[string]Query{
		createMatrix:  Query("INSERT INTO matrices ('name') VALUES ($1) RETURNING *"),
		deleteMatrix:  Query("UPDATE matrices SET deleted_at = NOW() WHERE uuid = $1"),
		getMatrix:     Query("SELECT * FROM matrices WHERE uuid = $1"),
		listMatrix:    Query("SELECT * FROM matrices"),
		updateMatrix:  Query("UPDATE matrices SET 'name' = $1 WHERE uuid = $2 RETURNING *"),
		addSubject:    Query("INSERT INTO matrix_subjects (matrix_id, subject_id) VALUES($1, $2)"),
		removeSubject: Query("DELETE matrix_subjects SET deleted_at = NOW() WHERE matrix_id = $1 AND subject_id = $2"),
	}
}
