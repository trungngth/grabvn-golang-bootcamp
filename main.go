package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		element := strings.Split(input, " ")
		if len(element) != 3 {
			fmt.Printf("Missing arguments or way too many arguments! %s\n", element)
			os.Exit(1)
		}
		execute(element[0], element[1], element[2])
		fmt.Printf("\n> ")
	}
}

func execute(first, operand, second string) {
	var result float64

	firstNumber, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("First input must be an integer!")
		os.Exit(1)
	}

	secondNumber, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println("Second input must be an integer!")
		os.Exit(1)
	}

	switch operand {
	case "+":
		result = float64(firstNumber) + float64(secondNumber)
	case "-":
		result = float64(firstNumber) - float64(secondNumber)
	case "*":
		result = float64(firstNumber) * float64(secondNumber)
	case "/":
		result = float64(firstNumber) / float64(secondNumber)
	default:
		fmt.Println("Invalid operand!")
		os.Exit(1)
	}
	printResult(first, operand, second, result)
}

func printResult(first, operand, second string, result float64) {
	fmt.Printf("%v %v %v = %v", first, operand, second, result)
}
