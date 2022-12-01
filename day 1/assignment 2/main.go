package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Deck []string

func (d Deck) print() {
	for _, v := range d {
		fmt.Println(v)
	}
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}

func main() {
	d := newDeckFromFile("my_deck")
	d.print()
	d.saveToFile("my_deck")
}
