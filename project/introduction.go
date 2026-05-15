package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Introduction() {

	fmt.Println("PLEASE INPUT YOUR NAME TO SIGN UP")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("FULLNAME: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(strings.ToLower(name))

	fmt.Println()
	fmt.Printf("HELLO, %v\n\n", Capitalize(name))

	var words = []string{
		Color("WELCOME TO WHO WANTS TO BE A MILLIONAIRE: LEEF CENTER EDITION SERIES 1", "\033[1;36m"),
		"",
		Color("YOU HAVE A CHANCE OF WINNING ₦20,000,000 IN ONE DAY", "\033[1;32m"),
		"",
		Color("BUT YOU MUST ANSWER 20 QUESTIONS CORRECTLY TO WIN THE GRAND PRIZE", "\033[1;33m"),
		"",
		Color("GUARANTEED LEVELS ARE QUESTION 5, 10 AND 15", "\033[1;35m"),
		"",
		Color("IF YOU FAIL, YOU FALL BACK PRIZE OF THE LAST GUARANTEED LEVEL", "\033[1;31m"),
		"",
		Color("YOU HAVE 3 LIFELINES:", "\033[1;36m"),
		Color("1. 50:50 - Removes two wrong answers", "\033[1;37m"),
		Color("2. HINT - Gives a clue", "\033[1;37m"),
		Color("3. SUGGESTION - Gives possible guidance", "\033[1;37m"),
		"",
		Color("ARE YOU READY? LET'S BEGIN THE GAME...", "\033[1;33m"),
	}

	TypeEffectFast(words)

	fmt.Println()
	fmt.Print(Color("PRESS ENTER TO START...", "\033[1;35m"))
	reader.ReadString('\n')
}
