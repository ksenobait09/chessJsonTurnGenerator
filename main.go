package main

import "fmt"

var englishLetters = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
}

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
}

func main() {
	var cells []string
	for _,letter := range englishLetters {
		for _, number := range numbers {
			cells = append(cells, fmt.Sprintf("%s %s", letter, number))
		}
	}
	var turnsAvailable []string
	for i,cell1 := range cells {
		for j, cell2 := range cells {
			if i != j {
				turnsAvailable = append(turnsAvailable, fmt.Sprintf("%s %s", cell1, cell2))
			}
		}
	}
	fmt.Println("[")
	for _, turn := range turnsAvailable {
		fmt.Printf("\t%q,\n", turn)
	}
	fmt.Println("]")
}
