package types

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string // gonna be hashed in the db
}

// implement the commented stuff later after other things are done

type Student struct {
	User
	StudentCourse Course
	// Grades          []StudentGrade
	// GPA             float64
	// DateOfBirth     string
}

type Teacher struct {
	User
}

type Admin struct {
	User
}
