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
	fmt.Printf("HELLO, %v", Capitalize(name))

	var words = []string{
		" WELCOME TO WHO WANTS TO BE A MILLIONAIRE LEEF CENTER EDITION SERIES 1.\n",
		"HERE YOU STAND A CHANCE OF WINNING 20 MILLION NARIA\n",
		"WE HAVE 20 QUESTIONS FOR YOU TO ANSWER AND IF YOU GIVE ME CORRECT ANSWERS TO ALL OF THEM FROM QUESTION 1-20 YOU WIN 20 MILLION NARIA\n",
		"\n",
		"YOU HAVE GUARANTEED LEVELS OF QUESTION 5, 10 AND 15. IF YOU GIVE ME THE WRONG ANSWER GO ",
	}
	TypeEffect(words)
}
