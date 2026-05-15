package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
	for i := 0; i+7 < len(line); i += 8 {
		maps[count] = line[i : i+8]
		count++
	}
	question := 0
	reader := bufio.NewReader(os.Stdin)
	usedLifeline := ""
	for i := 1; i <= 20; i++ {
		fmt.Printf(Color("FOR THIS QUESTION, YOUR PRIZE MONEY IS ₦%v\n", "\033[1;36m"), maps[i][6])
		fmt.Println(maps[i][:1])
		TypeEffectFast(maps[i][1:5])

	start:
		fmt.Print(Color("ANSWER: ", "\033[1;30m"))
		answer, _ := reader.ReadString('\n')
		answer = strings.ToUpper(strings.TrimSpace(answer))

		if answer == "LIFELINE" {
			fmt.Println()
			fmt.Print(Color("LIFELINES ARE;\n 1. 50:50\n 2. HINT\n 3. SUGGESTION\n", "\033[1;35m"))
			fmt.Println()
			fmt.Print(Color("OPTION: ", "\033[1;30m"))

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
					TypeEffectFast(maps[i][1:5])
					usedLifeline += "50:50"
					goto start
				}
			}

		}
		if answer == maps[i][len(maps[i])-3] {
			TypeEffectFast([]string{
				"CHECKING ANSWER...",
			})
			time.Sleep(2 * time.Second)

			TypeEffectFast([]string{
				"THE CORRECT ANSWER IS...",
			})
			time.Sleep(2 * time.Second)
			fmt.Printf(Color("%v\n", "\033[1m"), maps[i][len(maps[i])-3])

			fmt.Println(Color("CORRECT✅", "\033[1;33m"))
			fmt.Printf(Color("YOU HAVE WON ₦%v\n", "\033[1;32m"), maps[i][6])
			fmt.Println()
			question++
			ClearScreen()
		} else {
			TypeEffectFast([]string{
				"CHECKING ANSWER...",
			})
			time.Sleep(4 * time.Second)

			fmt.Println(Color("WRONG ❌", "\033[1;31m"))
			fmt.Printf(Color("THE CORRECT ANSWER IS %v\n", "\033[1m"), maps[i][len(maps[i])-3])

			if question < 5 {
				fmt.Println(Color("YOU HAVE NOT ANSWERED ANY GUARANTEED QUESTION THEREFORE YOU CANNOT WIN A CASH PRIZE", "\033[1;32m"))
				fmt.Println(Color("THANK YOU FOR YOUR TIME. WE HOPE TO SEE YOU AGAIN", "\033[1;32m"))
				ClearScreen()
			}
			if question < 10 && question >= 5 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦20,000", "\033[1;32m"))
				ClearScreen()
			}
			if question < 15 && question >= 10 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦500,000", "\033[1;32m"))
				ClearScreen()
			}
			if question < 20 && question >= 15 {
				fmt.Println(Color("THE PRIZE FOR YOUR PREVIOUS GUARANTEED LEVEL IS ₦3,000,000", "\033[1;32m"))
				ClearScreen()
			}

			fmt.Println()
			return
		}
	}

	var words = []string{
		Color("CONGRATULATIONS YOU HAVE WON ₦20,000,000", "\033[1;33m"),
		Color("TO ACCESS YOUR REWARD YOU WILL HAVE TO DROP YOUR BANK DETAILS HERE", "\033[1;33m"),
		Color("FORMAT; <account number>, <bankname>", "\033[1;33m"),
	}
	TypeEffectFast(words)
	fmt.Println(Color("INPUT HERE: ", "\033[1;30m"))

	bankDetails, _ := reader.ReadString('\n')
	bankDetails = strings.TrimSpace(strings.ToLower(bankDetails))

	fmt.Printf(Color("OK WE'VE GOT YOUR DETAILS: %v\n", "\033[1;32m"), bankDetails)
	fmt.Println(Color("FOR ANY COMPLAINT CONTACT US AT; 07036702434 or 09133440339", "\033[1;32m"))
	fmt.Println(Color("THANK YOU FOR YOUR TIME", "\033[1;32m"))

	var word = []string{
		Color("SEE YOU NEXT TIME ON WHO WANTS TO BE A MILLIONAIRE", "\033[1;42m"),
	}
	TypeEffectFast(word)
}
