package main

import (
	"fmt"
	"os"
	"strings"
)

func Hints(i int) {

	data, err := os.ReadFile("hint.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := strings.Split(string(data), "\n")

	fmt.Println(slice[i-1])
}

func CreatorsSuggestions(i int) {
	data, err := os.ReadFile("suggestions.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := strings.Split(string(data), "\n")

	fmt.Println(slice[i-1])
}

// func FiftyFifty(input []string) []string {
// 	data, err := os.ReadFile("fiftyfifty.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	slice := strings.Split(string(data), "\n")
// 	for i := 0; i < len(input); i++ {

// 	}

// }
