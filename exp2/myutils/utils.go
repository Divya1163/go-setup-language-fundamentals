package myutils

// Reverse reverses a string.
func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// CountVowels counts the number of vowels in a string.
func CountVowels(s string) int {
	count := 0

	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			count++
		}
	}

	return count
}

// Factorial calculates the factorial of a non-negative integer.
func Factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

// Power calculates base raised to the power of exponent.
func Power(base, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}