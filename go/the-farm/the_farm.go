package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	total, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	perCow := (total * factor) / float64(cows)

	return perCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	result, err := DivideFood(fodderCalculator, cows)
	if err != nil {
		return 0, err
	}
	return result, nil
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", err.cows, err.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
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
