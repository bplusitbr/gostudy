package main

import "fmt"

func main() {
	content := [...]int{3, 6, 7, 33, 88, 100, 0, 2, 1, 5, 200}
	max := -1
	for i := 0; i < len(content); i++ {
		if content[i] > max {
			max = content[i]
		}
	}

	fmt.Printf("Max value is %d\n", max)
}
