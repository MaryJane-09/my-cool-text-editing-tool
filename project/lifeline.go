package main

import (
	"fmt"
	"os"
	"strings"
)

func Hints() {

	data, err := os.ReadFile("hint.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := strings.Split(string(data), "\n")

	for i := 0; i < len(slice); i++ {

		fmt.Println(slice[i-1])
	}
}
