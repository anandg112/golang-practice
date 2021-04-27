package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b745",
		"white": "#fffff",
	}
	for colour, hex := range colors {
		fmt.Println("Hex Code for", colour, "is", hex)
	}
}
