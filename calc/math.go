package calc

import "errors"

// Returns sum of 2 integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("provide more than 2 numbers")
	} else {

		for _, num := range numbers {
			sum = sum + num
		}
		return sum, nil
	}
}
