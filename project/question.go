package main

import (
	"bufio"
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
	for i := 0; i < len(line)-7; i += 7 {
		maps[count] = line[i : i+7]
		count++
	}
	fmt.Println(maps)
	for i := 1; i <= 3; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(strings.Join(maps[i][:5], "\n"))
		// fmt.Println(maps)
		fmt.Print("ANSWER: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToUpper(answer))

		if answer != maps[i][len(maps[i])-2] {
			fmt.Println("WRONG ❌")
			break
		}
		fmt.Println("CORRECT✅")
		fmt.Println()

	}
}
