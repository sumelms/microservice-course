package database

const (
	createSubject = "create subject"
	deleteSubject = "delete subject by uuid"
	getSubject    = "get subject by uuid"
	listSubject   = "list subject"
	updateSubject = "update subject by uuid"
)

func queriesSubject() map[string]string {
	return map[string]string{
		createSubject: "INSERT INTO subjects (name) VALUES ($1) RETURNING *",
		deleteSubject: "UPDATE subjects SET deleted_at = NOW() WHERE uuid = $1",
		getSubject:    "SELECT * FROM subjects WHERE uuid = $1",
		listSubject:   "SELECT * FROM subjects",
		updateSubject: "UPDATE subjects SET name = $1 WHERE uuid = $2 RETURNING *",
	}
}
