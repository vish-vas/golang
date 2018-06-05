package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	colors1 := make(map[string]string)
	// var colors2 map[string]string

	colors1["white"] = "#ffffff"
	// colors2 = map[string]string{
	// 	"white": "#ffffff",
	// }
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, " : ", hex)
	}
}
