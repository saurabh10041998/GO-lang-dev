package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#7b45fe",
		"white": "#ffffff",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex value for color", color, "is", hex)
	}
}
