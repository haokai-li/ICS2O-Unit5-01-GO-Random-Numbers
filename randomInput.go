// Created by: haokai
// Created on: May 2021
// This program displays, "Random-Numbers"

package main

import (
	"fmt"

	"math/rand"

	"time"
)

func main() {

	// This function does addition
	var guess int
	var number int

	rand.Seed(time.Now().UnixNano())

	number = rand.Intn(6) + 1

	// input
	fmt.Println("This program is about Random.")
	fmt.Println()
	fmt.Print("Enter integers of number that you guess: ")
	fmt.Scanln(&guess)

	// detect
	if guess == number {
		// output of right answer
		fmt.Println("You are right!")
	} else if guess != number {
		// output of wrong answer
		fmt.Println("You are wrong!")
	}
	fmt.Println("\nDone.")
}
