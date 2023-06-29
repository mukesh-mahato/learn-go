package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("robots %s", string(robots))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
		}
	}()
	panic("something happened")
	fmt.Println("end")
}
