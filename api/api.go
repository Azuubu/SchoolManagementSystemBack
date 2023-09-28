package api

import (
	"encoding/json"
	"log"
	"net/http"
	"schoolsystem/database"
	"schoolsystem/types"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
	store database.Storage
}


var students []types.Student

func (s *ApiServer) Run() {
	r := mux.NewRouter()

	student1 := types.Student{
		User: types.User{ID: 1,FirstName: "Josh", LastName: "Kim",Email: "xp.gg@gmail.com", Password: "123"},
	} // ? it's for testing purposes

	students = append(students, student1)

	r.HandleFunc("/students", s.getStudents).Methods("GET") // this is admin only => ADMIN
	r.HandleFunc("/teachers", s.getTeachers).Methods("GET") // this is admin only => ADMIN

	//r.HandleFunc("/students", s.createStudentAcc).Methods("POST") // this is for normal user => USER

	r.HandleFunc("/student/account", s.getStudentAccount).Methods("GET") // USER
	r.HandleFunc("/student/account", s.createStudentAccount).Methods("POST") // USER
	r.HandleFunc("/student/account", s.updateStudentAccount).Methods("PUT") // USER / moze byc tez PATCH, to do ogarniecia potem

	r.HandleFunc("/teacher/account", s.getTeacherAccount).Methods("GET") // USER
	r.HandleFunc("/teacher/account", s.createTeacherAccount).Methods("POST") // USER


	//! Make and group functions by permissions, like: admin funcs, teacher funcs, student funcs

	//r.HandleFunc("/signup/teacher", AdminOnly())

	log.Println("Server running on port", s.listenAddress)
	http.ListenAndServe(s.listenAddress, r)
}

func NewApiServer(listenAddr string, store database.Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddr,
		store: store,
	}
}



func (s *ApiServer) getStudentAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *ApiServer) updateStudentAccount(w http.ResponseWriter, r *http.Request) {

}


func (s *ApiServer) createStudentAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student types.Student 
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}



func (s *ApiServer) getTeacherAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *ApiServer) createTeacherAccount(w http.ResponseWriter, r *http.Request) {

}

// either make get Users func and make it so that it works for both teachers and students,
// or split it into two funcs, 1 = getStudents, 2 = getTeachers