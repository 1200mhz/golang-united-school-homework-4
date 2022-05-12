package string_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFailWhenEmptyInputStringSum(t *testing.T) {
	s := "   "
	expected := ""
	actual, err := StringSum(s)

	if !reflect.DeepEqual(actual, expected) || err == nil || err.Error() != "input is empty" {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf("error, man")
	}
}

func TestFailWhenIncorrectNumberOfOperands(t *testing.T) {
	s := "1 - 2 + 88"
	expected := ""
	actual, err := StringSum(s)

	errorMessage := "expecting two operands, but received more or less"
	if !reflect.DeepEqual(actual, expected) || err == nil || err.Error() != errorMessage {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf("error, man")
	}
}

func TestSumOfPositiveOperands(t *testing.T) {
	s := "1 + 2"

	expected := "3"
	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}

func TestSumOfDifferentOperands(t *testing.T) {
	s := "-5 + 1"

	expected := "-4"
	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}

func TestSumOfNegativeOperands(t *testing.T) {
	s := "- 4 + -2"

	expected := "-6"
	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}
