package main

import "fmt"

func menu() { //Menu
	fmt.Println("Welcome to the shop")
	fmt.Println("We have the following items:")

	// This is the list of items in the shop.
	var shop_items = [5]string{"apple: £1.30", "banana: £1.50", "orange: £1.25", "grape: £1.50", "pineapple: £2.20"}
	fmt.Println(shop_items)

}

func order_calculater() { // This function calculates the total cost of the order.
	// The prices of the items.
	var apple float32 = 1.30
	var banana float32 = 1.50
	var orange float32 = 1.25
	var grape float32 = 1.50
	var pineapple float32 = 2.20
	// These variable are used to store the amount of each item the user wants to buy.
	var item_1 int
	var item_2 int
	var item_3 int
	var item_4 int
	var item_5 int

	fmt.Print("How many apples do you want to buy?: ")
	fmt.Scan(&item_1)
	fmt.Print("How many bananas do you want to buy?: ")
	fmt.Scan(&item_2)
	fmt.Print("How many oranges do you want to buy?: ")
	fmt.Scan(&item_3)
	fmt.Print("How many grapes do you want to buy?: ")
	fmt.Scan(&item_4)
	fmt.Print("How many pineapples do you want to buy?: ")
	fmt.Scan(&item_5)

	// This variable is used to store the total cost of the order.
	total := apple*float32(item_1) + banana*float32(item_2) + orange*float32(item_3) + grape*float32(item_4) + pineapple*float32(item_5)

	fmt.Println("Apples:", item_1, "Bananas:", item_2, "Oranges:", item_3, "Grapes:", item_4, "Pineapples:", item_5)

	// This prints the rounded total cost of the order.
	fmt.Printf("Your total is £:%.2f\n ", total)

}
func main() {

	menu()
	order_calculater()

}
