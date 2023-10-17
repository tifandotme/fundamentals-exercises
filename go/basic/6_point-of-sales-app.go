/*
Objectives
- Able to use the for statements
- Able to use the struct

Description
A cashier staff need an application to calculate the customers order by the given barcode list.

Below is list of product that listed on the store.

Barcode	Product Name		Price
00001		Coca-cola				3000
00002		Sprite					2500
00003		Fanta						2500
00004		Instant Noodles	3500
00005		Coffee					5000

There are 2 functions that you need to solve:
- CalculatePurchase: a function that receive slice of barcodes and return a PurchaseSummary. Purchase summary is a summary of the products which contain qty and sub-total, it also contain total which is sum of the all products multiply by the qty. For the detail cases, please check the unit test.
- GenerateReceiptText: a function for printing the purchase summary into the receipt. It is receive a purchase summary which is output of the func CalculatePurchase. For the detail cases, please check the unit test.

Examples
Please refer to the unit test for this exercise

Questions to mentor:
1. Guidelines for variables and types naming, sometimes they can be confusing.

*/

package exercise

import (
	"fmt"
	"os"
	"strconv"
)

type purchaseAmount struct {
	qty      int
	subtotal int
}

type purchaseSummary struct {
	products map[string]purchaseAmount
	total    int
}

type product struct {
	barcode string
	name    string
	price   int
}

var products = []product{
	{"00001", "Coca-cola", 3000},
	{"00002", "Sprite", 2500},
	{"00003", "Fanta", 2500},
	{"00004", "Instant Noodles", 3500},
	{"00005", "Coffee", 5000},
}

var scanned = []string{}

func POSApp() {
	for {
		clear()
		fmt.Println("POS App")
		fmt.Println()
		fmt.Printf("Scanned barcode: %v\n", scanned)
		fmt.Println()
		fmt.Println("1. Scan")
		fmt.Println("2. Print receipt")
		fmt.Println()
		fmt.Println("0. Exit")
		fmt.Println()

		input := prompt("Input (number): ")

		switch input {
		case 1:
			var recentlyAdded string

			for {
				clear()

				for idx, val := range products {
					fmt.Printf("%d. %s %s (Rp. %s)\n", idx+1, val.barcode, val.name, strconv.Itoa(val.price))
				}
				fmt.Println()
				fmt.Println("0. Back to menu")
				fmt.Println()
				if recentlyAdded != "" {
					fmt.Printf("Added %s!\n\n", recentlyAdded)
				}
				input := prompt("Scan product (pick the product index): ")

				if input == 0 {
					break
				}

				scanned = append(scanned, products[input-1].barcode)
				recentlyAdded = products[input-1].name
			}

		case 2:
			receipt := calculatePurchase(scanned)
			fmt.Println(receipt)

			fmt.Println()
			fmt.Println("0. Back to menu")
			fmt.Println()

			input := prompt("Input: ")

			if input == 0 {
				break
			}
			// TODO: finish this

		case 0:
			fmt.Println("Stopping program.")
			os.Exit(0)
		}
	}
}

func prompt(label string) (input int) {
	fmt.Print(label)
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	return
}

func calculatePurchase(barcode []string) (receipt purchaseSummary) {
	receipt = purchaseSummary{
		products: make(map[string]purchaseAmount),
		total:    0,
	}

	for _, val := range barcode {
		product := getProduct(val)

		if products, ok := receipt.products[product.name]; ok {
			products.qty += 1
			products.subtotal += product.price

			receipt.products[product.name] = products
		} else {
			receipt.products[product.name] = purchaseAmount{1, product.price}
		}

		receipt.total += product.price
	}

	return
}

func generateReceiptText(receipt purchaseSummary) {
	// TODO
}

func getProduct(barcode string) (product product) {
	for _, val := range products {
		if val.barcode == barcode {
			return val
		}
	}

	return
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
