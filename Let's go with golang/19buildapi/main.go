package main

import (
	"encoding/json"
	"fmt"
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
