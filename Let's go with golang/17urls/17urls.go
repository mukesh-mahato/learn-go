package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfdghd"

func main() {
	fmt.Println("Handling url in golang")
	fmt.Println(myurl) //https://lco.dev:3000/learn?courcename=reactjs&paymentid=gfdghd

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Port())   //3000
	fmt.Println(result.Path)     ///learn
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=gfdghd

	qparams := result.Query()
	fmt.Printf("The type of query parameters are %T\n", qparams) //The type of query parameters are url.Values

	fmt.Println(qparams["coursename"]) //[reactjs]
	fmt.Println(qparams["paymentid"])  //[gfdghd]

	for _, value := range qparams {
		fmt.Println("params is: ", value) //params is:  [reactjs],  params is:  [gfdghd]
	}

	pathOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=andrew",
	}

	anotherUrl := pathOfUrl.String()
	fmt.Println(anotherUrl) //https://lco.dev/learn
}
