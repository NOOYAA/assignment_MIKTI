package main

import (
	"fmt"
)

func main() {
	var menu_choosen int
	var keep_loop bool = true
	
	for keep_loop {
		fmt.Println("\nWelcome!")
		fmt.Println("\n1. Print Hello World\n2. Math Ops\n3. Store and Save Data User\n4. Exit.")
		fmt.Print("Type a number to select menu: ")
		fmt.Scan(&menu_choosen)
		
		if menu_choosen == 1 {
			fmt.Println("\nHello World")

		} else if menu_choosen == 2 {
			var math_menu int

			fmt.Println("\nSimple Calculator!")
			fmt.Println("\nChoose your opration!\n1. Addtion\n2. Substraction\n3. Multiplication\n4. Division\n5. Back\n6. Exit.")
			fmt.Print("Type a Number to choose Opration: ")
			fmt.Scan(&math_menu)
			
			if math_menu == 1 {
				var num1, num2 = promptShare()

				add := num1 + num2
				fmt.Println("Result: ", add)

			} else if math_menu == 2 {
				var num1, num2 = promptShare()

				substract := num1 - num2
				fmt.Println("Result: ", substract)

			} else if math_menu == 3 {
				var num1, num2 = promptShare()

				multiple := num1 * num2
				fmt.Println("Result: ", multiple)

			} else if math_menu == 4 {
				var num1, num2 = promptShare()

				divide := num1 / num2
				fmt.Println("Result: ", divide)

			} else if math_menu == 6 {
				break
			} else if math_menu == 5 {
				continue
			}
		} else if menu_choosen == 3 {

			var favfood_menu int

			fmt.Println("\nFriend's Favorite Food!")
			fmt.Println("\n1. Register friend's Name and Food\n2. Find Friend and their food\n3. Back\n4. Exit.")
			fmt.Print("Choose what you want: ")
			fmt.Scan(&favfood_menu)

			if favfood_menu == 1 {

			}
		} else if menu_choosen == 4 {
			keep_loop = false
		}
	}
}
func promptShare() (num1 int, num2 int) {

	fmt.Print("Type your first number: ")
	fmt.Scan(&num1)
	fmt.Print("Type your second number: ")
	fmt.Scan(&num2)

	return num1, num2
}

// func dataStored(name string, favorite_food string) map[string]string{} {
// 	user := map[string]string{}

// 	user[name] = favorite_food
// 	return user
// }
