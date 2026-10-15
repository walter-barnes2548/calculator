package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Calculator: enter an expression such as 12.5 * 8, or q to quit")
	for {
		fmt.Print("> ")
		if !in.Scan() {
			return
		}
		line := strings.TrimSpace(in.Text())
		if strings.EqualFold(line, "q") || strings.EqualFold(line, "quit") {
			return
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			fmt.Println("use: number operator number")
			continue
		}
		a, errA := strconv.ParseFloat(fields[0], 64)
		b, errB := strconv.ParseFloat(fields[2], 64)
		if errA != nil || errB != nil {
			fmt.Println("both operands must be numbers")
			continue
		}
		var result float64
		switch fields[1] {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				fmt.Println("cannot divide by zero")
				continue
			}
			result = a / b
		default:
			fmt.Println("operator must be +, -, *, or /")
			continue
		}
		fmt.Println(result)
	}
}
