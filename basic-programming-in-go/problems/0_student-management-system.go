package bootcamp

import "fmt"

type student struct {
	id    int
	name  string
	age   uint8
	grade float64
}

func StudentManagementSystem() {
	data := make([]student, 0)

	var choice int8

	for {

		fmt.Println("1. Add")
		fmt.Println("2. Display")
		fmt.Println("3. Calculate")
		fmt.Println("0. Exit")
		fmt.Print("Input: ")
		fmt.Scanln(&choice)
		fmt.Printf("\n")

		switch choice {
		case 1:
			var student student

			fmt.Print("ID: ")
			fmt.Scanln(&student.id)
			fmt.Print("Name: ")
			fmt.Scanln(&student.name)
			fmt.Print("Age: ")
			fmt.Scanln(&student.age)
			fmt.Print("Grade: ")
			fmt.Scanln(&student.grade)

			data = append(data, student)

		case 2:
			for i, v := range data {
				fmt.Printf("Student %d:\n", i+1)
				fmt.Printf("ID: %d\n", v.id)
				fmt.Printf("Name: %s\n", v.name)
				fmt.Printf("Age: %d\n", v.age)
				fmt.Printf("Grade: %.2f\n\n", v.grade)
			}

		case 3:
			var sum float64
			for _, v := range data {
				sum += v.grade
			}

			average := sum / float64(len(data))

			fmt.Printf("Average: %.2f\n", average)
		case 0:
			return

		default:
			fmt.Println("Wrong choice buckaroo")
		}
	}
}
