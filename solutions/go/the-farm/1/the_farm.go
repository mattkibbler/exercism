package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalc FodderCalculator, numCows int) (float64, error) {
	amount, err := fodderCalc.FodderAmount(numCows)
	if err != nil {
		return 0.0, err
	}
	fatFactor, err := fodderCalc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return ((amount / float64(numCows)) * fatFactor), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalc FodderCalculator, num int) (float64, error) {
	if num <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalc, num)
}

type InvalidCowsError struct {
	numCows int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numCows, err.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(num int) error {
	if num < 0 {
		return &InvalidCowsError{numCows: num, message: "there are no negative cows"}
	}
	if num == 0 {
		return &InvalidCowsError{numCows: num, message: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
