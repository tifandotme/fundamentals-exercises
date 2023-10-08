/*
Objectives
Able to use the conditional statements

Description
You are tasked to create a program to calculate parking fee.

Parking Price List

Duration	   Motorcycle	  Car
First 1 hr	 3000					7000
After 1 hr	 2000/hour		5000/hour

If vehicle parked more than 24 hours then they will be an extra charged, as below:

Vehicle Type		Extra Charge
Motorcycle			+20000
Car							+50000
*/

package problems

import "fmt"

type price struct {
	firstHour   int
	perHour     int
	extraCharge int
}

var parkingFee = map[string]price{
	"motorcycle": {
		firstHour:   3000,
		perHour:     2000,
		extraCharge: 20000,
	},
	"car": {
		firstHour:   7000,
		perHour:     5000,
		extraCharge: 50000,
	},
}

func LotBilling() {
	fmt.Printf("Lot Parking:\n")
	fmt.Println(`Choose vehicle:
1. Motorcycle
2. Car`)

	var vehicleType int8
	fmt.Print("Input vehicle type (number):")
	fmt.Scanln(&vehicleType)

	if vehicleType != 1 && vehicleType != 2 {
		fmt.Println("Invalid vehicle type.")
		return
	}

	var parkingTime int
	fmt.Print("Input parking time (hour):")
	fmt.Scanln(&parkingTime)

	if parkingTime <= 0 {
		fmt.Println("Parking time can't be 0 or negative number.")
		return
	}

	var vehicle string
	if vehicleType == 1 {
		vehicle = "motorcycle"
	} else {
		vehicle = "car"
	}

	var calculatedFee int

	calculatedFee += parkingFee[vehicle].firstHour
	if parkingTime > 1 {
		calculatedFee += (parkingTime - 1) * parkingFee[vehicle].perHour
	}
	if parkingTime > 24 {
		calculatedFee += parkingFee[vehicle].extraCharge
	}

	fmt.Printf("Parking fee: %d\n", calculatedFee)
}
