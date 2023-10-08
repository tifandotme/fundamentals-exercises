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

package problems

import "fmt"

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

var data = []product{
	{"00001", "Coca-cola", 3000},
	{"00002", "Sprite", 2500},
	{"00003", "Fanta", 2500},
	{"00004", "Instant Noodles", 3500},
	{"00005", "Coffee", 5000},
}

func POSApp() {
	receipt := calculatePurchase([]string{"00001", "00005"})

	generateReceiptText(receipt)
}

func calculatePurchase(barcode []string) (receipt purchaseSummary) {
	receipt = purchaseSummary{
		products: make(map[string]purchaseAmount),
		total:    0,
	}

	for _, val := range barcode {
		product := searchProduct(val)

		if products, ok := receipt.products[product.name]; ok {
			fmt.Println("ok")
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
	// wait for unit test
}

func searchProduct(barcode string) (product product) {
	for _, val := range data {
		if val.barcode == barcode {
			return val
		}
	}

	return
}
