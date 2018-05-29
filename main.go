package main

import "fmt"

func main() {

	//different ways to declare map
	//Option 1. declaring a map where all of the keys inside the map are of type string in all the values are of type string as well
	//colors := map[string]string{
	//color and equivalent to hex code
	//	"red":   "#ff0000",
	//	"green": "#4ssrw",

	//}

	//Option 2.
	//var colors map[string]sting

	//Option 3.
	//colors := make(map[string]string)
	//colors["white"] = "#fffff"

	//Extention of Option 3
	colors := make(map[int]string)
	colors[10] = "#yyyyy"

	delete(colors, 10)

	fmt.Println(colors)
}
