package main

import "fmt"

func printColor(v map[string]string) {
	for k, x := range v {
		fmt.Println(k, " : ", x)
	}
}

func main() {
	m := make(map[string]string)

	m["red"] = "#ff0000"
	m["blue"] = "#0000ff"
	m["yellow"] = "#ffff00"
	m["green"] = "#00ff00"

	printColor(m)

	fmt.Println(m["yellow"])

	delete(m, "yellow")
	fmt.Println(m["yellow"])

	printColor(m)
}
