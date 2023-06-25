package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang") //Files in golang
	content := "We are learning golang"

	file, err := os.Create("./myFiles.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length) //length is: 22
	defer file.Close()
	readFile("./myFiles.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is\n", string(dataByte)) //Text data inside the file is We are learning golang
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
