package intermediate

import (
	"fmt"
	"math/rand/v2"
)

func random_numbers() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// // fmt.Println(rand.Intn(100)+1)

	// fmt.Println(val.Intn(101))

	fmt.Println("Game: Roll the dice")
	for {
		fmt.Println("Menu")
		fmt.Println("1. Roll Dice")
		fmt.Println("2. Exit")

		var choice string
		_, err := fmt.Scan(&choice)

		if err != nil || (choice != "1" && choice != "2") {
			fmt.Println("Invalid choice. Please choose 1 or 2")
			continue
		}

		if choice == "2" {
			fmt.Println("Thanks for playing")
			break
		}

		die1 := rand.IntN(6) + 1
		die2 := rand.IntN(6) + 1

		fmt.Println("Die 1:", die1)
		fmt.Println("Die 2:", die2)


		fmt.Println("Do you wanna play again ?")

		var restart string
		_, err1 := fmt.Scan(&restart)
		if err1 != nil || (restart != "Y" && restart != "N") {
			fmt.Println("invalid choice assuming No, Goodbye")
			break
		}

		if restart == "N" {
			fmt.Println("Thanks for playing")
			break
		}

	}
}
