package main

import (
	"fmt"
	"strings"
)

func style(text, style, color string) string {
	return style + color + text + "\033[0m"
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s)
}

func lastLetter(s string) string {

	if len(s) > 1 {
		return string(s[len(s)-1])
	}
	return s
}

func deleteIndex(s string, index int) string {

	if index < 0 || index >= len(s) {
		return "ERROR!! INDEX IS OUT OF RANGE IT MUST BE BETWEEN 0 AND %d"
	}
	return s[:index] + s[index+1:]
}

func main() {
star:
	fmt.Println(" ")
	fmt.Println(style("INPUT WORD: ", "\033[33m", ""))

	var word string
	fmt.Scanln(&word)

	fmt.Println(" ")
	fmt.Println(style("************* AVAILIABLE OPERATIONS ***************", "\033[35m", ""))
	fmt.Println(" ")
	fmt.Println(style("PLEASE SELECT AN OPERATION: ", "\033[33m", ""))

start:
	fmt.Println(style("1. LAST LETTER", "\033[34m", ""))
	fmt.Println(style("2. CAPITALIZE", "\033[34m", ""))
	fmt.Println(style("3. DELETE INDEX", "\033[34m", ""))
	fmt.Println(" ")

	fmt.Print(style("OPTION: ", "\033[30m", ""))
	var option string
	fmt.Scanln(&option)

	var result string
	switch option {

	case "1":
		fmt.Println(" ")
		result = lastLetter(word)

	case "2":
		fmt.Println(" ")
		result = capitalize(word)

	case "3":
	start1:
		fmt.Println(" ")
		fmt.Println(style("INPUT THE INDEX YOU WANT TO DELETE", "\033[33m", ""))
		var index int
		fmt.Scanln(&index)
		if index < 0 || index >= len(word) {
			fmt.Printf(style("ERROR!! INDEX IS OUT OF RANGE IT MUST BE BETWEEN 0 AND %d\n", "\033[31m", ""), len(word)-1)
			fmt.Println(style("TRY AGAIN", "\033[31m", ""))
			goto start1
		}
		result = deleteIndex(word, index)
	default:
		fmt.Println(style("ERROR!! PLEASE SELECT FROM THE GIVING RANGE", "\033[31m", ""))
		fmt.Println(" ")
		goto start

	}

	fmt.Println(style("RESULT:", "\033[32m", ""), result)
	fmt.Println(" ")

	fmt.Println(style("WOULD YOU LIKE TO PERFORM ANOTHER OPERATION??", "\033[35m", ""))
	fmt.Println(" ")
	fmt.Println(style("YES", "\033[34m", ""))
	fmt.Println(style("NO", "\033[34m", ""))
	fmt.Println(" ")
stap:
	fmt.Print(style("OPTION: ", "\033[30m", ""))
	var resuse string
	fmt.Scanln(&resuse)

	switch resuse {
	case "yes":
		fmt.Println(" ")
		goto star

	case "no":
		fmt.Println(" ")
		fmt.Println(style("IT WAS NICE WORKING WITH YOU!", "\033[32m", ""))
		fmt.Println(style("GOODBYE", "\033[32m", ""))
		return

	default:
		fmt.Println(style("SELECT EITHER YES/NO", "\033[31m", ""))
		fmt.Println(" ")
		goto stap
	}
}
