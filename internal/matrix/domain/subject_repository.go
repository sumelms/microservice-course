package domain

import (
	"github.com/google/uuid"
)

type SubjectRepository interface {
	Subject(subjectUUID uuid.UUID) (Subject, error)
	Subjects() ([]Subject, error)
	CreateSubject(subject *Subject) error
	UpdateSubject(subject *Subject) error
	DeleteSubject(subject *DeletedSubject) error
}
