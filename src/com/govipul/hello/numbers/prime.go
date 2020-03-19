package numbers

import "math"

// IsPrime function to check given number is prime.
func IsPrime(number int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(number)))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}
