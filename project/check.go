package main

import "fmt"


func CheckIfLifelineHasBeenUsed(input bool) error {
	if input == true {
		return fmt.Errorf("THIS LIFELINE HAS ALREADY BEEN USED")
		
	}
	return  nil
}