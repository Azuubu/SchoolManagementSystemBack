package types

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

// implement the commented stuff later after other things are done

type Student struct {
	User
	CurrentSemester int
	// Subjects        []string
	// Grades          []StudentGrade
	// GPA             float64
	// DateOfBirth     string
}

type Teacher struct {
	User
	// DateOfBirth string
}

type Admin struct {
	User
}
