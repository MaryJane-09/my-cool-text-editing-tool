package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AccessingPrize() {

	reader := bufio.NewReader(os.Stdin)
	var words = []string{
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
	fmt.Println()

	var word = []string{
		"\n",
		Color("SEE YOU NEXT TIME ON WHO WANTS TO BE A MILLIONAIRE", "\033[1;96m"),
		"\n",
	}
	TypeEffectFast(word)
}
