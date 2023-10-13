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

import (
	"fmt"
	"os"
)

type price struct {
	firstHour   int
	perHour     int
	extraCharge int
}

var fees = map[string]price{
	"motorcycle": {3000, 2000, 20000},
	"car":        {7000, 5000, 50000},
}

func LotBilling() {
	fmt.Println("LOT PARKING")
	fmt.Println()
	fmt.Println("Choose vehicle:")
	fmt.Println("1. Motorcycle")
	fmt.Println("2. Car")
	fmt.Println()

	vehicle := promptVehicleType()

	duration := promptParkingDuration()

	fmt.Printf("\nParking fee: %d\n", calculateParkingFee(vehicle, duration))
}

func promptVehicleType() (vehicle string) {
	var input uint8
	fmt.Print("Input vehicle type (number):")
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	if input != 1 && input != 2 {
		fmt.Println("Choose the correct vehicle (atleast 1 or 2).")
		os.Exit(0)
	}

	if input == 1 {
		vehicle = "motorcycle"
	} else {
		vehicle = "car"
	}

	return
}

func promptParkingDuration() (duration int) {
	fmt.Print("Input parking time (hour):")
	if _, err := fmt.Scanln(&duration); err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	if duration < 0 {
		fmt.Println("Parking time can't be negative number.")
		os.Exit(0)
	}

	return
}

func calculateParkingFee(vehicle string, duration int) (result int) {
	result += fees[vehicle].firstHour
	if duration > 1 {
		result += (duration - 1) * fees[vehicle].perHour
	}
	if duration > 24 {
		result += fees[vehicle].extraCharge
	}

	return
}
