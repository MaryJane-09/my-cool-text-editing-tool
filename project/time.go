package main

import (
	"fmt"
	"strings"
	"time"
)

func TypeEffect(word []string) {
	for _, line := range word {
		for _, char := range line {
			fmt.Print(string(char))
			time.Sleep(60 * time.Millisecond)
		}
		time.Sleep(800 * time.Millisecond)
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