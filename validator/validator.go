package validator

import (
	"errors"
)

func Validate(a, b int) error {
	error := errors.New("panic validate")
	data := []int{a, b}
	for _, val := range data {
		if (val < 1) || (val > 10) {
			return error
		}
	}
	return nil
}
