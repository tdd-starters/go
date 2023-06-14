package main

import (
	"fmt"
	"gostarter/calculator"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	result := calculator.Evaluate(input)
	fmt.Printf("%s = %d\n", input, result)
}
