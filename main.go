package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func style(text, ansi string) string {
	return ansi + text + "\033[0m"
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func lastCharacter(s string) string {
	if len(s) > 0 {
		return string(s[len(s)-1])
	}
	return s
}

func deleteIndex(s string, index int) string {
	if index < 0 || index >= len(s) {
		return s
	}
	return s[:index] + s[index+1:]
}

func reverseString(input string) string {
	words := strings.Fields(input)
	var reverse = func(s string) string {
		res := ""
		for i := len(s) - 1; i >= 0; i-- {
			res += string(s[i])
		}
		return res
	}
	for i := range words {
		words[i] = reverse(words[i])
	}
	return strings.Join(words, " ")
}

func countCharacter(s string) string {
	count := 0
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return strconv.Itoa(count)
}

func hasLetter(s string) bool {
	for _, char := range s {
		if !((char >= 'A' && char <= 'Z') ||
			(char >= 'a' && char <= 'z') ||
			char == ' ') {
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

	fmt.Println(" ")
	fmt.Println(style("*************** WELCOME TO MJ'S COOL TEXT EDITING TOOL ****************", "\033[1;33m"))
	fmt.Println(" ")

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println(" ")
		fmt.Print(style("TEXT: ", "\033[1;33m"))

		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		word = strings.Join(strings.Fields(word), " ")

		if word == "" {
			fmt.Println(style("INPUT CANNOT BE EMPTY", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			continue
		}

		if !(hasLetter(word) || hasNumber(word)) {
			fmt.Println(style("MUST BE ONLY LETTERS (SPACES ALLOWED BETWEEN THEM) OR ONLY NUMBERS (NO SPACES ALLOWED BETWEEN THEM)", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			continue
		}

		var result string

		fmt.Println(" ")
		fmt.Println(style("************* AVAILABLE OPERATIONS ***************", "\033[1;35m"))
		fmt.Println(" ")

		for {

			fmt.Println(" ")
			fmt.Println(style("PLEASE SELECT AN OPERATION: ", "\033[1;33m"))
			fmt.Println(" ")
			fmt.Println(style("1. LAST CHARACTER", "\033[1;34m"))
			fmt.Println(style("2. UPPERCASE WORDS", "\033[1;34m"))
			fmt.Println(style("3. DELETE INDEX", "\033[1;34m"))
			fmt.Println(style("4. REVERSE WORD", "\033[1;34m"))
			fmt.Println(style("5. COUNT CHARACTERS", "\033[1;34m"))
			fmt.Println(style("6. EXIT PROGRAM", "\033[1;34m"))
			fmt.Println(" ")

			fmt.Print(style("OPTION: ", "\033[1;30m"))

			option, _ := reader.ReadString('\n')
			option = strings.TrimSpace(option)

			switch option {

			case "1":
				result = lastCharacter(word)

			case "2":
				if !hasLetter(word) {
					fmt.Println(style("INPUT DOES NOT MATCH OPERATION", "\033[1;31m"))
					fmt.Println(style("CANNOT CAPITALIZE NUMBERS", "\033[1;31m"))
					fmt.Println(" ")

					fmt.Println(style("WOULD YOU LIKE TO;\n  1.CHANGE THE OPERATION   OR\n 2.EXIT PROGRAM\n", "\033[1;33m"))
					fmt.Println(" ")

					for {
						fmt.Print(style("OPTION: ", "\033[1;30m"))

						choice, _ := reader.ReadString('\n')
						choice = strings.TrimSpace(choice)

						switch choice {

						case "1":

						case "2":

							fmt.Println(" ")
							fmt.Println(style("EXITING..........", "\033[1m"))
							fmt.Println(" ")
							return

						default:
							fmt.Println(style("PLEASE SELECT FROM THE RANGE OF OPTIONS", "\033[1;31m"))
							fmt.Println(style("TRY AGAIN", "\033[1;31m"))
							fmt.Println(" ")
							continue
						}
						break
					}

					continue
				}
				result = toUpper(word)

			case "3":
				for {
					fmt.Print(style("INPUT THE INDEX YOU WANT TO DELETE: ", "\033[1;33m"))

					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(input)

					index, err := strconv.Atoi(input)

					if err != nil {
						fmt.Println(style("INDEX SHOULD BE AN INTEGER. NO SPACES ALLOWED BETWEEN THEM", "\033[1;31m"))
						fmt.Println(style("TRY AGAIN", "\033[1;31m"))
						continue
					}

					parts := strings.Fields(input)

					var indices []int

					for _, p := range parts {
						num, err := strconv.Atoi(p)
						if err != nil {
							fmt.Println(style("INDEX SHOULD BE AN INTEGER. NO SPACES ALLOWED BETWEEN THEM", "\033[1;31m"))
							fmt.Println(style("TRY AGAIN", "\033[1;31m"))
							continue

						}
					}

					if index < 0 || index >= len(word) {
						fmt.Printf(style("ERROR!! INDEX IS OUT OF RANGE IT MUST BE BETWEEN 0 AND %d\n", "\033[1;31m"), len(word)-1)
						fmt.Println(style("TRY AGAIN", "\033[1;31m"))
						fmt.Println(" ")
						continue
					}

					indices = append(indices, index)

					if len(indices) == 0 {
						fmt.Println(style("NO VALID INDICES PROVIDED", "\033[1;31m"))
						continue
					}

					removedChar := string(word[index])

					highlighted := word[:index] + "[" + style(removedChar, "\033[1;31m") + "]" + word[index+1:]

					result = fmt.Sprintf(
						"\nOriginal: %s\nRemoved: %s\nNew Result: %s",
						highlighted,
						style(removedChar, "\033[1;31m"),
						deleteIndex(word, index),
					)
					break
				}

			case "4":
				result = reverseString(word)

			case "5":
				result = countCharacter(word)

			case "6":
				fmt.Println(" ")
				fmt.Println(style("EXITING.......", "\033[1m"))
				fmt.Println(" ")
				return

			default:
				fmt.Println(style("ERROR!! PLEASE SELECT FROM THE GIVING RANGE", "\033[1;31m"))
				fmt.Println(" ")
				continue
			}

			break
		}

		fmt.Println(" ")
		fmt.Println(style("RESULT:", "\033[1;32m"), result)

		for {
			fmt.Println(" ")
			fmt.Println(style("WOULD YOU LIKE TO PERFORM ANOTHER OPERATION??", "\033[1;33m"))
			fmt.Println(" ")
			fmt.Println(style("YES", "\033[1;34m"))
			fmt.Println(style("NO", "\033[1;34m"))
			fmt.Println(" ")

			fmt.Print(style("OPTION: ", "\033[1;30m"))

			reuse, _ := reader.ReadString('\n')
			reuse = strings.TrimSpace(strings.ToLower(reuse))

			if reuse == "yes" {
				fmt.Println(" ")
				break
			} else if reuse == "no" {
				fmt.Println(" ")
				fmt.Println(style("IT WAS NICE WORKING WITH YOU!🎊", "\033[1;32m"))
				fmt.Println(style("GOODBYE 👋", "\033[1;32m"))
				return
			} else {
				fmt.Println(style("SELECT EITHER YES/NO", "\033[1;31m"))
				fmt.Println(style("TRY AGAIN", "\033[1;31m"))
				fmt.Println(" ")
			}
		}
	}
}
