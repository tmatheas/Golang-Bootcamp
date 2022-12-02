package main

import "fmt"

func uniqueNames(x []string, y []string) []string {
	var unique []string
	for _, a := range x {
		found := false
		for _, b := range y {
			if a == b {
				found = true
			}
		}
		if !found {
			unique = append(unique, a)
		}
	}

	for _, a := range y {
		found := false
		for _, b := range unique {
			if a == b {
				found = true
			}
		}
		if !found {
			unique = append(unique, a)
		}
	}
	return unique
}

func main() {
	slice1 := []string{"Ava", "Emma", "Olivia"}
	slice2 := []string{"Olivia", "Sophia", "Emma"}
	unique := uniqueNames(slice1, slice2)
	fmt.Println(unique)
}
