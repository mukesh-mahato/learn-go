package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shotrs for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
