package random

import (
	"fmt"
)

func challenge1(input str[]) str[] {
	var results: str[] = [] // slice of strings of size input, with default values 'no'
	for item in input {
		results.append(false)
		for i = 1; i++; i <= len(item) {
			if item(0) != item(i) {
				return false
			}
		}
	}
}

func main() {
	challenge1()
}
