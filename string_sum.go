package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if operandsCount(input) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	var firstOperand, secondOperand int

	operPos := strings.LastIndex(input, "+")
	if operPos == -1 {
		operPos = strings.LastIndex(input, "-")
	}

	if firstOperand, err = strconv.Atoi(strings.TrimLeft(input[0:operPos], "+")); err != nil {
		return "", fmt.Errorf("failed to convert %q: %w", input[0:operPos], err)
	}
	if secondOperand, err = strconv.Atoi(strings.TrimLeft(input[operPos:], "+")); err != nil {
		return "", fmt.Errorf("failed to convert %q: %w", input[operPos:], err)
	}

	return strconv.Itoa(firstOperand + secondOperand), nil

}

func operandsCount(input string) (count int) {
	input = strings.TrimLeft(input, "+")
	input = strings.TrimLeft(input, "-")

	for _, subStr := range strings.Split(input, "-") {
		count += len(strings.Split(subStr, "+"))
	}
	return
}
