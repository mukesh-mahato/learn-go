package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang") //Welcome to time study of golang

	presentTime := time.Now()
	fmt.Println(presentTime) //2023-06-25 07:02:54.357169994 +0545 +0545 m=+0.000030791

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //06-25-2023 07:02:54 Sunday

	createDate := time.Date(2020, time.May, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)                             //2020-05-12 23:23:00 +0000 UTC
	fmt.Println(createDate.Format("01-02-2006 Monday")) //05-12-2020 Tuesday
}
