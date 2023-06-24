package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "We are learning golang"

	file, err := os.Create("./myFiles.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myFiles.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is\n", string(dataByte))
}
