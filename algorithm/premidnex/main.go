package main

import "fmt"

func main() {
	content := [...]int{3, 9, 6, 7, 33, 88, 100, 0, 2, 1, 5, 200}

	middlarray := int((len(content) / 2) - 1)

	middle := content[middlarray]
	previous := content[middlarray-1]
	next := content[middlarray+1]

	fmt.Printf("Previous value is %d\n", previous)
	fmt.Printf("Middle value is %d\n", middle)
	fmt.Printf("Next value is %d\n", next)
}
