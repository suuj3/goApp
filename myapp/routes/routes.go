package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	router.HandleFunc("/students", controller.GetAllStuds)

	//course routes
	router.HandleFunc("/course", controller.AddCourse).Methods("POST")          //to post course
	router.HandleFunc("/course/{cid}", controller.GetCour).Methods("GET")       //to get the data from url
	router.HandleFunc("/course/{cid}", controller.UpdateCour).Methods("PUT")    //to update the data
	router.HandleFunc("/course/{cid}", controller.DeleteCour).Methods("DELETE") //to delete the data

	router.HandleFunc("/courses", controller.GetAllCour).Methods("GET") //print all students

	//enroll routes
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST") //to post enroll
	router.HandleFunc("/enroll/{sid}/{cid}", controller.GetEnroll).Methods("GET")

	router.HandleFunc("/enrolls", controller.GetEnrolls).Methods("GET")

	router.HandleFunc("/enroll/{sid}/{cid}", controller.DeleteEnroll).Methods("DELETE") //delete enroll

	// admin login and signup
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/logout", controller.Logout)

	fhandler := http.FileServer(http.Dir("./index"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port 7080...")
	log.Fatal(http.ListenAndServe(":7080", router))
}
