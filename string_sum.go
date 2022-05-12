package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errorEmptyInput
	}

	s := regexp.MustCompile(`\d`)
	digits := s.FindAllStringIndex(input, -1)
	if len(digits) != 2 {
		return "", errorNotTwoOperands
	}

	var x, y int
	var signMinus, definedX, definedY bool

	for _, v := range input {
		val := string(v)
		if !definedX {
			cur, err := strconv.Atoi(val)
			if err != nil {
				if val == "-" {
					signMinus = true
				} else if val == "+" || val == " " {
				} else {
					return "", fmt.Errorf("incorrect input string 1")
				}
			} else {
				if signMinus {
					x = -cur
				} else {
					x = cur
				}
				signMinus = false
				definedX = true
			}
		} else if definedX && !definedY {
			cur, err := strconv.Atoi(val)
			if err != nil {
				if val == "-" {
					signMinus = true
				} else if val == "+" || val == " " {
				} else {
					return "", fmt.Errorf("incorrect input string 2")
				}
			} else {
				if signMinus {
					y = -cur
				} else {
					y = cur
				}
				signMinus = false
				definedY = true
			}
		}
	}

	return strconv.Itoa(x + y), nil
}
