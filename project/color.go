package main

func Color(input, text string) string {
	return text + input + "\033[0m"
}
