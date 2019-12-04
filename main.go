package main

import (
	"bufio"
	"fmt"
	"log"
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
			log.Fatal("Missing arguments or way too many arguments!", element)
		}
		execute(element[0], element[1], element[2])
		fmt.Printf("\n> ")
	}
}

func execute(first, operand, second string) {
	var result float64

	firstNumber, err := strconv.Atoi(first)
	if err != nil {
		log.Fatal("First input must be an integer!")
	}

	secondNumber, err := strconv.Atoi(second)
	if err != nil {
		log.Fatal("Second input must be an integer!")
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
		log.Fatal("Invalid operand!")
	}
	printResult(first, operand, second, result)
}

func printResult(first, operand, second string, result float64) {
	fmt.Printf("%v %v %v = %v", first, operand, second, result)
}

func test() {
	fmt.Println("Test")
}
