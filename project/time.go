package main

import (
	"fmt"
	"strings"
	"time"
)

func TypeEffectFast(word []string) {
	for _, line := range word {
		for _, char := range line {
			fmt.Print(string(char))
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(900 * time.Millisecond)
	}
}

func TypeEffectSlow(word []string) {
	for _, line := range word {
		for _, char := range line {
			fmt.Print(string(char))
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(900 * time.Millisecond)
	}
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower(word[i][1:])
	}
	return  strings.Join(word, " ")
}