package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("question.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	maps := make(map[int][]string)
	line := strings.Split(string(data), "\n")

	//fmt.Println(strings.Join(line, "\n"))
	count := 1
	for i := 0; i < len(line)-6; i += 6 {
		maps[count] = line[i : i+6]
		count++
	}
	for i := 1; i <= 3; i++ {
		words := ""
		fmt.Println(strings.Join(maps[i], "\n"))
		fmt.Println("you answer:")
		fmt.Scanln(&words)
	}
}
