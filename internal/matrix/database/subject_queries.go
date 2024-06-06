package database

import "fmt"

const (
	returningSubjectColumns = `uuid, code, name, objective, credit, workload,
		created_at, updated_at, subjects.published_at`
	// CREATE.
	createSubject = "create subject"
	// READ.
	getSubject   = "get subject by uuid"
	listSubjects = "list subject"
	// UPDATE.
	updateSubject = "update subject by uuid"
	// DELETE.
	deleteSubject = "delete subject by uuid"
)

func queriesSubject() map[string]string {
	return map[string]string{
		// CREATE.
		createSubject: fmt.Sprintf(`INSERT INTO subjects (
			code, name, objective, credit, workload, published_at
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING %s`, returningSubjectColumns),
		// READ.
		getSubject: fmt.Sprintf(`SELECT
				%s
			FROM subjects
			WHERE uuid = $1
				AND subjects.deleted_at IS NULL`, returningSubjectColumns),
		listSubjects: fmt.Sprintf(`SELECT
				%s
			FROM subjects
			WHERE subjects.deleted_at IS NULL`, returningSubjectColumns),
		// UPDATE.
		updateSubject: fmt.Sprintf(`UPDATE subjects
			SET code = $2, name = $3, objective = $4, credit = $5, workload = $6, published_at = $7
			WHERE uuid = $1 AND deleted_at IS NULL
			RETURNING %s`, returningSubjectColumns),
		// DELETE.
		deleteSubject: `UPDATE subjects
			SET deleted_at = NOW()
			WHERE uuid = $1 AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
	}
}
