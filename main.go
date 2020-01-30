package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	shoot()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func shoot() {

	rand.Seed(time.Now().UnixNano())

	ai := randomInt(0, 3)
	input := 0

	fmt.Println("Choose 0 for Rock, 1 for Paper, 2 for Scissors.")

	fmt.Scanln(&input)

	switch {

	case input == ai:
		fmt.Println("DRAW")

	case input == 0 && ai == 1:
		fmt.Println("You Lose. Paper beats Rock")

	case input == 0 && ai == 2:
		fmt.Println("You Win. Rock beats Scissors")

	case input == 1 && ai == 2:
		fmt.Println("You Lose. Scissors beats Paper")

	case input == 1 && ai == 0:
		fmt.Println("You Win. Paper beats Rock")

	case input == 2 && ai == 0:
		fmt.Println("You Lose. Rock beats Scissors")

	case input == 2 && ai == 1:
		fmt.Println("You Win. Scissors beats Rock")
	}

}
