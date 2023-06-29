package main

import (
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseId    string  `json: "coueseid"`
	Coursename  string  `json: "coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json: "fullname"`
	Website  string `json: "website"`
}

// fake DB
var coueses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.Coursename == ""
}

func main() {

}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>We are learning to build api in golang</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coueses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	// for _, course := range Course {
	// 	if course.CourseId == params["id"] {
	// 		json.NewEncoder(w).Encode(course)
	// 		return
	// 	}
}

// 	json.NewEncoder(w).Encode("No Course found with given id")
// }
