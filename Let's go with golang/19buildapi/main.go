package main

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
