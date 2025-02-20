package main

import (
	"fmt"
	"sort"
)

type Dictionnary struct {
	words map[string]string
}

func (d *Dictionnary) Add(word, definition string) {
	if d.words == nil {
		d.words = make(map[string]string)
	}
	d.words[word] = definition
}

// ordered list by alphabetical order
func (d *Dictionnary) List() []string {
	words := make([]string, 0, len(d.words))
	for word := range d.words {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	fmt.Println("Strating the CLI...")

	dictionnary := Dictionnary{}
	dictionnary.Add("hello", "world")

	dict := dictionnary.List()
	for _, word := range dict {
		fmt.Println("mot:", word, ", definition:", dictionnary.words[word])
	}

}
