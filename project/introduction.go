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
	fmt.Printf("WELCOME, %v", Capitalize(name))

	var words = []string{
		" TO WHO WANTS TO BE A MILLIONAIRE LEEF CENTER EDITION.\n",
	}
	TypeEffect(words)
}
