package database

const (
	createSubject = "create subject"
	deleteSubject = "delete subject by uuid"
	getSubject    = "get subject by uuid"
	listSubject   = "list subject"
	updateSubject = "update subject by uuid"
)

type Query string

func queries() map[string]Query {
	return map[string]Query{
		createSubject: Query("INSERT INTO subjects (name) VALUES ($1) RETURNING *"),
		deleteSubject: Query("UPDATE subjects SET deleted_at = NOW() WHERE uuid = $1"),
		getSubject:    Query("SELECT * FROM subjects WHERE uuid = $1"),
		listSubject:   Query("SELECT * FROM subjects"),
		updateSubject: Query("UPDATE subjects SET name = $1 WHERE uuid = $2 RETURNING *"),
	}
}
