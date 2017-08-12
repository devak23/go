package assorted

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"
)

func DoubleTheMoneyMain() {
	chances := 6
	bet := 0
	fmt.Println("***********************************************************")
	fmt.Printf("Roll dice to get 6 in %d chances and double your bet.\n", chances)
	fmt.Println("***********************************************************")
	fmt.Printf("Enter the betting amount: ")
	fmt.Scanln(&bet)

	if bet <= 0 {
		fmt.Println("Sorry, can't play unless there is money on the table")
		return
	}

	for {
		fmt.Printf("rolling the dice...")
		result := rollDice()
		time.Sleep(1 * time.Second)
		fmt.Printf("got: %d\n", result)
		chances--

		if chances <= 0 {
			time.Sleep(2 * time.Second)
			fmt.Println("You lose... try again later.")
			break
		}

		if (result == 6) {
			time.Sleep(3 * time.Second)
			bet := bet * 2
			fmt.Printf("\nYou win! you got double the money your bet just got doubled (%d)\n", bet)
			break
		} else {
			time.Sleep(1 * time.Second)
			fmt.Printf("You have %d more chances. Press Enter to roll the dice or \"q\" to quit: ", chances)
			reader := bufio.NewReader(os.Stdin)
			quit, _ , _ := reader.ReadLine()
			if string(quit) == "Q" || string(quit) == "q" {
				time.Sleep(1 * time.Second)
				fmt.Println("\nYou should have played the game as you wont get your money back!")
				time.Sleep(1 * time.Second)
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}

func rollDice() int {
	max, min  := 6 , 1
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}
