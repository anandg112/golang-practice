package main

import "fmt"

func main() {
	slice10 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range slice10 {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}
}
