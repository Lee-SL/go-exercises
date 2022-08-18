package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

var ScaleError = errors.New("negative fodder")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	fodder, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodder >= 0 {
		fodder *= 2
		return fodder / float64(cows), nil
	} else if err != nil && err != ErrScaleMalfunction {
		return 0.0, err
	} else if fodder < 0.0 {
		return 0.0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{
			message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows),
		}
	}

	return fodder / float64(cows), err
}
