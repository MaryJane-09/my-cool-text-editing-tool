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

	if len(s) > 0 {
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

func hasLetter(s string) bool {
	for _, char := range s {
		if !(char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z') {
			return false
		}
	}
	return true
}

func hasNumber(s string) bool {
	for _, char := range s {
		if !(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}

func main() {
star:
	fmt.Println(" ")
	fmt.Println(style("INPUT : ", "\033[33m", ""))

	var word string
	fmt.Scanln(&word)

	if word == "" {
		fmt.Println(style("INPUT CANNOT BE EMPTY", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		goto star
	}

	isLetter := hasLetter(word)
	isNumber := hasNumber(word)

	if !(isLetter || isNumber) {
		fmt.Println(style("MUST EITHER BE AN ALPHABET OR AN INTEGER", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		goto star
	}

	fmt.Println(" ")
	fmt.Println(style("************* AVAILIABLE OPERATIONS ***************", "\033[1;35m", ""))
	fmt.Println(" ")
	fmt.Println(style("PLEASE SELECT AN OPERATION: ", "\033[1;33m", ""))

start:
	fmt.Println(style("1. LAST LETTER", "\033[1;34m", ""))
	fmt.Println(style("2. CAPITALIZE", "\033[1;34m", ""))
	fmt.Println(style("3. DELETE INDEX", "\033[1;34m", ""))
	fmt.Println(style("4. EXIT PROGRAM", "\033[1;34m", ""))
	fmt.Println(" ")

	fmt.Print(style("OPTION: ", "\033[1;30m", ""))
	fmt.Print(" ")
	var option string
	fmt.Scanln(&option)

	var result string
	switch option {

	case "1":
		fmt.Println(" ")
		result = lastLetter(word)

	case "2":
		if !hasLetter(word) {

			fmt.Println(style("INPUT DOES NOT MATCH OPERATION", "\033[1;31m", ""))
			fmt.Println(style("CANNOT CAPITALIZE INTEGERS", "\033[1;31m", ""))
			fmt.Println(" ")
			fmt.Println(style("WOULD YOU LIKE TO;\n 1.RESTART\n  OR\n 2.CHANGE THE OPERATION\n", "\033[1;33m", ""))
			fmt.Println(" ")
		stop:
			fmt.Print(style("OPTION: ", "\033[1;30m", ""))

			var choice string
			fmt.Scanln(&choice)

			switch choice {

			case "1":
				goto star

			case "2":
				goto start

			default:
				fmt.Println(style("PLEASE SELECT FROM THE RANGE OF OPTIONS", "\033[1;31m", ""))
				fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
				fmt.Println(" ")
				goto stop
			}

		}
		fmt.Println(" ")
		result = capitalize(word)

	case "3":
	start1:
		fmt.Println(" ")
		fmt.Println(style("INPUT THE INDEX YOU WANT TO DELETE", "\033[1;33m", ""))
		var index int
		n, err := fmt.Scanln(&index)

		if n == 0 || err != nil {
			fmt.Println(style("INDEX SHOULD BE AN INTEGER", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))

			var discard string
			fmt.Scanln(&discard)
			goto start1
		}
		if index < 0 || index >= len(word) {
			fmt.Printf(style("ERROR!! INDEX IS OUT OF RANGE IT MUST BE BETWEEN 0 AND %d\n", "\033[1;31m", ""), len(word)-1)
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			goto start1
		}

		result = deleteIndex(word, index)
	case "4":
		fmt.Println(" ")
		fmt.Println(style("EXITING..........", "\033[1m", ""))
		fmt.Println(" ")
		return
	default:
		fmt.Println(style("ERROR!! PLEASE SELECT FROM THE GIVING RANGE", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start

	}

	fmt.Println(style("RESULT:", "\033[1;32m", ""), result)
	fmt.Println(" ")
	fmt.Println(style("WOULD YOU LIKE TO PERFORM ANOTHER OPERATION??", "\033[1;35m", ""))
	fmt.Println(" ")
	fmt.Println(style("YES", "\033[1;34m", ""))
	fmt.Println(style("NO", "\033[1;34m", ""))
	fmt.Println(" ")
stap:
	fmt.Print(style("OPTION: ", "\033[1;30m", ""))
	fmt.Println(" ")
	var resuse string
	fmt.Scanln(&resuse)

	switch resuse {
	case "yes", "YES":
		fmt.Println(" ")
		goto star

	case "no", "NO":
		fmt.Println(" ")
		fmt.Println(style("IT WAS NICE WORKING WITH YOU!🎊", "\033[1;32m", ""))
		fmt.Println(style("GOODBYE", "\033[1;32m", ""))
		return

	default:
		fmt.Println(style("SELECT EITHER YES/NO", "\033[1;31m", ""))
		fmt.Println(" ")
		goto stap
	}
}
