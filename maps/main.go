package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	for c, h := range colors {
		fmt.Println("The hex code for", c, "is", h)
	}
}
