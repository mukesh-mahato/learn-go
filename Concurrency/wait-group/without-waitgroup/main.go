package main

import (
	"fmt"
	"time"
)

func main() {

}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}
