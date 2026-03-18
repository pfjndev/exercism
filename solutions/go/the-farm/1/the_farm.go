package thefarm

import (
	"fmt"
	"errors"
)

func DivideFood(fc FodderCalculator, nCows int) (float64, error) {
	totalFodderAmmount, err := fc.FodderAmount(nCows)
	if err != nil { return 0, err }
	fatteningFactor, err := fc.	FatteningFactor()
	if err != nil { return 0, err }
	
	return (totalFodderAmmount * fatteningFactor) / float64(nCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 { return 0, errors.New("invalid number of cows") }
	return DivideFood(fc, nCows)
}

type InvalidCowsError struct {
	nCows	int
	msg		string
}

func (iCE *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", iCE.nCows, iCE.msg)
}

func ValidateNumberOfCows(nCows int) error {
	switch {
		case nCows < 0:		return &InvalidCowsError {
			nCows:	nCows,
			msg:	"there are no negative cows",
		}
		case nCows == 0:	return &InvalidCowsError {
			nCows:	nCows,
			msg:	"no cows don't need food",
		}
		default:			return nil
	}
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
