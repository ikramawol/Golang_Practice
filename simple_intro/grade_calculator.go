package main

import "fmt"

func calculateAverage(total float64, n int) float64 {
	return total / float64(n)
}

func main() {

	fmt.Println("Welcome to the Grade Calculator!")
	var studentgradelist = make(map[string]float64)
	for {
		// Declare variables
		var name string
		var Grade float64
		var n int
		var total float64
		var average float64

		fmt.Println("To exit the program, type 'exit' as your name.")
		fmt.Println("Please enter your Name: ")
		fmt.Scanln(&name)

		if name == "exit" {
			break
		}
		fmt.Println("Please enter number of subjects:")
		fmt.Scanln(&n)

		if n <= 0 {
			fmt.Println("Please enter a valid number of subjects.")
			continue
		}

		for i := range n {
			fmt.Println("Please enter your grade for subject", i+1, ":")
			fmt.Scanln(&Grade)
			if Grade < 0 || Grade > 100 {
				fmt.Println("Please enter a valid grade between 0 and 100.")
				i--
				continue
			}
			total += Grade
		}

		average = calculateAverage(total, n)
		studentgradelist[name] = average
		fmt.Println(studentgradelist)
		fmt.Printf("Your average grade is: %.2f\n", average)
	}
}
