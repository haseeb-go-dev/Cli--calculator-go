package main

import (
	"fmt"
	"os"
)

func cal(a float64, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cant divide with 0")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("wrong operator: %s", op)
	}
}

func main() {
	var a float64
	var b float64
	var op string
	
	fmt.Println("CLI Calculator - type 'exit' to quit")
	
	for {
		fmt.Print("enter num sign num -> ")
		fmt.Scanln(&a, &op, &b)
		
		if op == "exit" {
			os.Exit(0)
	}
		
		total, err := cal(a, b, op)
		
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result = ", total)
	}
	}
