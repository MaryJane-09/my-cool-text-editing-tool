package main

import "fmt"


func CheckIfLifelineHasBeenUsed(input bool) error {
	if input {
		return fmt.Errorf("THIS LIFELINE HAS ALREADY BEEN USED\nTRY AGAIN")
		
	}
	return  nil
}