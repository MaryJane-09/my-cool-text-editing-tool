package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Question() {
	data, err := os.ReadFile("question.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	maps := make(map[int][]string)
	line := strings.Split(string(data), "\n")

	count := 1
	for i := 0; i < len(line)-8; i += 8 {
		maps[count] = line[i : i+8]
		count++
	}

	for i := 1; i <= 3; i++ {
		fmt.Printf("FOR THIS QUESTION, YOUR PRIZE MONEY IS ₦%v\n", maps[i][6])
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(strings.Join(maps[i][:5], "\n"))

start:
		fmt.Print("ANSWER: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))

		if answer == "lifeline" {
			fmt.Println("LIFELINES ARE;\n 50 : 50\n HINTS\n SUGGESTIONS")
			lifeline, _ := reader.ReadString('\n')
			lifeline = strings.ToLower(strings.TrimSpace(lifeline))
			if lifeline == "hint" {
				Hints()
				goto start
			}
		}

		if answer != maps[i][len(maps[i])-3] {
			fmt.Println("WRONG ❌")
			break
		} else {
			fmt.Println("CORRECT✅")
			fmt.Printf("YOU HAVE WON ₦%v\n", maps[i][6])
			fmt.Println()
		}
	}
}
