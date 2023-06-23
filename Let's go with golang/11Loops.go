package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	num := 1
	for num < 10 {

		if num == 3 {
			goto lgo
		}

		if num == 5 {
			num++
			continue
		}
		if num == 8 {
			break
		}
		fmt.Println("Value is", num)
		num++
	}

lgo:
	fmt.Println("We are learning Golang")

}
