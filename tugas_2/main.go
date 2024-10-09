package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var keep_loop bool = true
	order := map[string]int{}
	grandTotal := 0

	for keep_loop {
		fmt.Println("Menu:")
		menuList := menu()

		for food, price := range menuList {
			fmt.Printf("- %s: Rp%d.00\n", food, price)
		}

		fmt.Print("\nMasukan nama item (ketik \"selesai\" untuk menyelesaikan): ")
		return_input := userInput()

		formatted_input := uniformInput(return_input)

		if formatted_input == "nasigoreng" {

			fmt.Print("\nPesen Nasi Goreng berapa? ")
			jumlah := userInput()
			var intJumlah int = cstrInt(uniformInput(jumlah))
			addedItem("Nasi Goreng", intJumlah, order)

		} else if formatted_input == "miegoreng" {

			fmt.Print("\nPesen Mie Goreng berapa? ")
			jumlah := userInput()
			var intJumlah int = cstrInt(uniformInput(jumlah))
			addedItem("Mie Goreng", intJumlah, order)

		} else if formatted_input == "ayambakar" {
			fmt.Print("\nPesen Ayam Bakar berapa? ")
			jumlah := userInput()
			var intJumlah int = cstrInt(uniformInput(jumlah))
			addedItem("Ayam Bakar", intJumlah, order)
		} else if formatted_input == "selesai" {

			if len(order) != 0 {
				fmt.Println("\nPesanan kamu:")
			}

			for food, amount := range order {
				for f, p := range menuList {
					if f == food {
						totalPrice := order[food] * p
						grandTotal += totalPrice
						fmt.Printf("- %s   x%d: %d.00\n", food, amount, totalPrice)

					}
				}

			}
			fmt.Printf("\nTotal harga: Rp%d.00\n", grandTotal)
			pembayaran(grandTotal)
			keep_loop = false
		} else {
			fmt.Println("\nhmm... kayaknya ada yang salah dengan inputan mu!")
		}
	}

}

func cstrInt(numString string) int {
	intVal, _ := strconv.Atoi(numString)
	return intVal
}

func menu() map[string]int {
	var dishes = map[string]int{
		"Nasi Goreng": 25000,
		"Mie Goreng":  22000,
		"Ayam Bakar":  30000,
	}
	return dishes
}

func addedItem(food string, amount int, order map[string]int) {
	defer fmt.Printf("\n-- %s -- di tambahkan!\n\n", food)

	order[food] = amount
}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func uniformInput(sentence string) (clean_sentence string) {
	lowercase := strings.ToLower(sentence)
	clean_sentence = strings.TrimSpace(strings.ReplaceAll(lowercase, " ", ""))
	return
}

func pembayaran(money int) {
	looping := true

	for looping {
		fmt.Print("Masukan nominal yang di bayarkan: ")
		returnInput := userInput()
		formattedInput := uniformInput(returnInput)
		intInput := cstrInt(formattedInput)

		if intInput > money {
			looping = false
			change := intInput - money
			fmt.Printf("Jumlah yang di bayar valid. Kembalian: Rp%d.00\n", change)
		} else if intInput == money {
			looping = false
			fmt.Println("Jumlah yang di bayar valid. Terima Kasih!\n")
		} else {
			fmt.Printf("Nominal yang di bayar tidak sesuai.\n\nTotal harga: Rp%d.00\n\n", money)
		}
	}

}
