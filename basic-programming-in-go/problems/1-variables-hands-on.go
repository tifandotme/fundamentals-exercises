/*
Objectives
- Able to use variables
- Able to use operators
- Able to determine the appropriate data types

Description
You are tasked to calculate Shopee's annual gross profit from 10 different product.

Product			Price				Total Sold	Discount
Product A		Rp 100.000	200					0
Product B		Rp 67.000		12					20
Product C		Rp 56.000		80					0
Product D		Rp 1.000		1350				0
Product E		Rp 20.000		1						0
Product F		Rp 38.455		7						15
Product G		Rp 76.000		5644				0
Product H		Rp 530.120	30					10
Product I		Rp 143.000	54					0
Product J		Rp 16.000		109					0

The formula to calculate the gross profit is:
Gross Profit = Σ ((price x total sold) - (price x discount x total sold))

You need to return the annual gross profit value.

Questions to mentor:
The formula to calculate gross profit is the total revenue minus the cost of goods sold.

*/

package problems

import (
	"fmt"
	"strconv"
)

type name = string

type productsAnnualProfit = map[name]struct {
	price     int
	totalSold int
	discount  int
}

func VariablesHandsOn() {
	var products = productsAnnualProfit{
		"Product A": {100000, 200, 0},
		"Product B": {67000, 12, 20},
		"Product C": {56000, 80, 0},
		"Product D": {1000, 1350, 0},
		"Product E": {20000, 1, 0},
		"Product F": {38455, 7, 15},
		"Product G": {76000, 5644, 0},
		"Product H": {530120, 30, 10},
		"Product I": {143000, 54, 0},
		"Product J": {16000, 109, 0},
	}

	calculate(products)
}

// Calculate annual gross profit of all products
func calculate(products productsAnnualProfit) {
	var sum int

	for key, val := range products {
		discount := 1.0 - (float64(val.discount) / 100.0)
		discountedPrice := int(float64(val.price) * discount)

		currSum := val.price * val.totalSold
		if discount < 1.0 {
			currSum -= discountedPrice * val.totalSold
		}
		sum += currSum

		fmt.Printf("%s: Rp %s\n", key, strconv.Itoa(currSum))
	}

	fmt.Printf("Sum: Rp %s\n", strconv.Itoa(sum))
}

// Calculate annual gross profit of all products (without considering discount)
func calculateAlt(products productsAnnualProfit) {
	var sum int

	for key, val := range products {
		currSum := (val.price * val.totalSold) - (val.price * val.discount * val.totalSold)
		sum += currSum

		fmt.Printf("%s: Rp %s\n", key, strconv.Itoa(currSum))
	}

	fmt.Printf("Sum: Rp %s\n", strconv.Itoa(sum))
}
