package validators

import (
	"strconv"
)

func ValidatePositiveInt(val string) bool {
	valInt, err := strconv.Atoi(val)
	return err == nil && valInt > 0
}

func ValidateNonNegativeInt(val string) bool {
	valInt, err := strconv.Atoi(val)
	return err == nil && valInt >= 0
}