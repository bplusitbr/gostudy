package main

import "fmt"

func main() {
	content := [...]int{3, 6, 7, 33, 88, 100, 0, 2, 1, 5, 200}

	middle := content[len(content)/2]

	fmt.Printf("Middle value is %d\n", middle)
}
