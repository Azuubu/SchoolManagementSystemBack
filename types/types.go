package types

type StudentGrade struct {
	Grade   int
	Subject string
}

type Course struct {
	Name     string
	Subjects []Subject
}

type Subject struct {
	Name           string
	SubjectTeacher Teacher
}