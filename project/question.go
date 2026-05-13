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
		os.Exit(1)
	}

	maps := make(map[int][]string)
	line := strings.Split(string(data), "\n")

	count := 1
	for i := 0; i < len(line)-8; i += 8 {
		maps[count] = line[i : i+8]
		count++
	}
	question := 0
	reader := bufio.NewReader(os.Stdin)
	usedLifeline := ""
	for i := 1; i <= 20; i++ {
		fmt.Printf("FOR THIS QUESTION, YOUR PRIZE MONEY IS ₦%v\n", maps[i][6])
		fmt.Println(strings.Join(maps[i][:1], "\n"))
		TypeEffectFast(maps[i][1:][:4])

	start:
		fmt.Print("ANSWER: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.ToUpper(strings.TrimSpace(answer))

		if answer == "LIFELINE" {
			fmt.Println()
			fmt.Print("LIFELINES ARE;\n 1. 50:50\n 2. HINT\n 3. SUGGESTION\n")
			fmt.Println()
			fmt.Print("OPTION: ")

			lifeline, _ := reader.ReadString('\n')
			lifeline = strings.ToLower(strings.TrimSpace(lifeline))
			if lifeline == "hint" || lifeline == "2" {
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "hint"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto start
				} else {
					Hints(i)
					usedLifeline += "hint"
					fmt.Println()
					goto start
				}
			}
			if lifeline == "suggestion" || lifeline == "3" {
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "suggestion"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto start
				} else {
					CreatorsSuggestions(i)
					usedLifeline += "suggestion"
					fmt.Println()
					goto start
				}
			}
			if lifeline == "50:50" || lifeline == "1" {
				err := CheckIfLifelineHasBeenUsed(strings.Contains(usedLifeline, "50:50"))
				if err != nil {
					fmt.Println(err)
					fmt.Println()
					goto start
				} else {
					FiftyFifty(maps[i], i)
					fmt.Println(strings.Join(maps[i][:1], "\n"))
					TypeEffectFast(maps[i][1:][:4])
					usedLifeline += "50:50"
					goto start
				}
			}

		}
		if answer == maps[i][len(maps[i])-3] {
			fmt.Println("CORRECT✅")
			fmt.Printf("YOU HAVE WON ₦%v\n", maps[i][6])
			fmt.Println()
			question++
		} else {
			fmt.Println("WRONG ❌")
			fmt.Printf("THE CORRECT ANSWER IS %v\n", maps[i][len(maps[i])-3])

			if question < 5 {
				fmt.Println("YOU HAVE NOT ANSWERED ANY GUARANTEED QUESTION THEREFORE YOU CANNOT WIN A CASH PRIZE")
				fmt.Println("THANK YOU FOR YOUR TIME. WE HOPE TO SEE YOU AGAIN")
			}
			if question < 10 {
				fmt.Println("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦20,000")
			}
			if question < 15 {
				fmt.Println("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦500,000")
			}
			if question < 20 {
				fmt.Println("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦3,000,000")
			}

			fmt.Println()
			return
		}
	}

	var words = []string {
		"CONGARTULATIONS YOU HAVE WON ₦20,000,000",
		"TO ACCESS YOUR REWARD YOU WILL HAVE TO DROP YOUR BANK DETAILS HERE",
		"FORMAT; <account number>, <bankname>",
	}
	TypeEffectFast(words)
	fmt.Println("INPUT HERE: ")

	bankDetails, _ := reader. ReadString('\n')
	bankDetails = strings.TrimSpace(strings.ToLower(bankDetails))

	fmt.Printf("OK WE'VE GOT YOUR DETAILS: %v\n", bankDetails)
	fmt.Println("FOR ANY COMPLAINT CONTACT US AT; 07036702434 or 09133440339")
	fmt.Println("THANK YOU FOR YOUR TIME")

	var word = []string {
		"SEE YOU NEXT TIME ON WHO WANTS TO BE A MILLIONAIRE",
	}
	TypeEffectFast(word)
}
