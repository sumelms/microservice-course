package database

import "fmt"

const (
	returningMatrixColumns = `matrices.uuid, matrices.code, matrices.name, matrices.description,
		matrices.created_at, matrices.updated_at`
	// CREATE.
	createMatrix        = "create matrix"
	createMatrixSubject = "create matrix subject"
	// READ.
	getMatrix             = "get matrix by uuid"
	getMatrixSubject      = "get matrix subject by uuid"
	getCourseMatrixExists = "get whether there is a relationship between course and matrix"
	listMatrices          = "list matrices"
	listCourseMatrices    = "list course matrices by course_uuid"
	listMatrixSubjects    = "list matrix subjects by matrix_uuid"
	// UPDATE.
	updateMatrix        = "update matrix by uuid"
	updateMatrixSubject = "update matrix subject by matrix and subject uuid"
	// DELETE.
	deleteMatrix        = "delete matrix by uuid"
	deleteMatrixSubject = "delete matrix subject"
)

//nolint:funlen
func queriesMatrix() map[string]string {
	return map[string]string{
		// CREATE.
		createMatrix: fmt.Sprintf(`INSERT INTO matrices (
				course_id, code, name, description
			) SELECT courses.id, $2, $3, $4
				FROM courses
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
			RETURNING %s`, returningMatrixColumns),
		createMatrixSubject: `INSERT INTO matrix_subjects (
			matrix_id, subject_id, is_required, "group"
		) SELECT matrices.id, subjects.id, $3, $4
			FROM matrices, subjects
			WHERE
				matrices.deleted_at IS NULL
				AND matrices.uuid = $1
				AND subjects.deleted_at IS NULL
				AND subjects.uuid = $2
			RETURNING uuid, created_at, updated_at`,
		// READ.
		getMatrix: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				%s
			FROM matrices
			LEFT JOIN courses ON matrices.course_id = courses.id
			WHERE matrices.uuid = $1 AND matrices.deleted_at IS NULL
				AND courses.deleted_at IS NULL`, returningMatrixColumns),
		getMatrixSubject: `SELECT
				matrix_subjects.uuid,
				matrix_subjects.is_required,
				matrix_subjects.group,
				matrix_subjects.created_at,
				matrix_subjects.updated_at
			FROM matrix_subjects
			LEFT JOIN matrices ON matrix_subjects.matrix_id = matrices.id
			LEFT JOIN subjects ON matrix_subjects.subject_id = subjects.id
			WHERE matrices.uuid = $1
				AND matrices.deleted_at IS NULL
				AND subjects.uuid = $2
				AND subjects.deleted_at IS NULL
				AND matrix_subjects.deleted_at IS NULL`,
		getCourseMatrixExists: `SELECT EXISTS (
			SELECT 1
			FROM matrices
			INNER JOIN courses ON matrices.course_id = courses.id
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
				AND matrices.uuid = $2 AND matrices.deleted_at IS NULL)`,
		listMatrices: fmt.Sprintf(`SELECT
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
		listMatrixSubjects: `SELECT
				matrix_subjects.uuid,
				subjects.uuid AS subject_uuid,
				matrices.uuid AS matrix_uuid,
				matrix_subjects.is_required AS is_required,
				matrix_subjects.group AS "group",
				matrix_subjects.created_at AS created_at,
				matrix_subjects.updated_at AS updated_at
			FROM matrix_subjects
			LEFT JOIN matrices ON matrix_subjects.matrix_id = matrices.id
			LEFT JOIN subjects ON matrix_subjects.subject_id = subjects.id
			WHERE matrices.uuid = $1
				AND matrices.deleted_at IS NULL
				AND subjects.deleted_at IS NULL
				AND matrix_subjects.deleted_at IS NULL`,
		// UPDATE.
		updateMatrix: fmt.Sprintf(`UPDATE matrices
			SET code = $2, name = $3, description = $4, updated_at = NOW()
			WHERE uuid = $1
			RETURNING %s`, returningMatrixColumns),
		updateMatrixSubject: `UPDATE matrix_subjects
			SET is_required = $3, "group" = $4, updated_at = NOW()
			WHERE matrix_id in (
					SELECT id FROM matrices WHERE deleted_at IS NULL AND uuid = $1
				) AND subject_id in (
					SELECT id FROM subjects WHERE deleted_at IS NULL AND uuid = $2
				) AND deleted_at IS NULL
			RETURNING uuid, created_at, updated_at`,
		// DELETE.
		deleteMatrix: `UPDATE matrices
			SET deleted_at = NOW()
			WHERE uuid = $1
				AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
		deleteMatrixSubject: `UPDATE matrix_subjects
			SET deleted_at = NOW()
			WHERE
				matrix_id = (SELECT id FROM matrices WHERE uuid = $1 AND deleted_at IS NULL)
				AND subject_id = (SELECT id FROM subjects WHERE uuid = $2 AND deleted_at IS NULL)
				AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
	}
}
