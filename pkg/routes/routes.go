package routes

import (
	"github.com/anaard/simple-student-management/pkg/controllers"
	"github.com/anaard/simple-student-management/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterStudentManagementRoutes = func(router *mux.Router) {
	UserController := controllers.UserController{}

	router.HandleFunc("/register", UserController.Register).Methods("POST")
	router.HandleFunc("/login", UserController.Login).Methods("POST")

	SystemController := controllers.SystemController{}

	router.HandleFunc("/student", middleware.CheckAuth(SystemController.CreateStudent)).Methods("POST")
	router.HandleFunc("/student/all", middleware.CheckAuth(SystemController.GetStudents)).Methods("GET")

	router.HandleFunc("/student/{studentId}", middleware.CheckAuth(SystemController.GetStudentByID)).Methods("GET")
	router.HandleFunc("/student/{studentId}", middleware.CheckAuth(SystemController.UpdateStudent)).Methods("PUT")
	router.HandleFunc("/student/{studentId}", middleware.CheckAuth(SystemController.DeleteStudent)).Methods("DELETE")

	router.HandleFunc("/class", middleware.CheckAuth(SystemController.CreateClass)).Methods("POST")             // create class
	router.HandleFunc("/class/all", middleware.CheckAuth(SystemController.GetClasses)).Methods("GET")           // return brief statistics from each class (id, number of students)
	router.HandleFunc("/class/{classId}", middleware.CheckAuth(SystemController.GetClassById)).Methods("GET")   // return n° students and the students
	router.HandleFunc("/class/{classId}", middleware.CheckAuth(SystemController.DeleteClass)).Methods("DELETE") // return class and remove students from there

	router.HandleFunc("/{classId}/{studentId}", middleware.CheckAuth(SystemController.EnrollStudentInClass)).Methods("POST")     // enroll student in class
	router.HandleFunc("/{classId}/{studentId}", middleware.CheckAuth(SystemController.RemoveStudentFromClass)).Methods("DELETE") // remove student from class
}
