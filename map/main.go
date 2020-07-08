package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4HF745",
		"white": "#FFFFFF",
	}
	// colors[1] = "#FFFFFF"
	// delete(colors, 1)
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
