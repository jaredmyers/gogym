package main

// given large int represented as int array of digits.
// digits are their significant digit order
// increment the large int by one and return the result

func plusOne(digits []int) []int {

	sigDigitIdx := len(digits) - 1
	return flip(digits, sigDigitIdx)
}

func flip(digits []int, sigDigitIdx int) []int {

	if sigDigitIdx < 0 {
		digits = append(digits, 0)
		digits[0] = 1
		return digits
	}

	if digits[sigDigitIdx] == 0 {

		digits[sigDigitIdx] = 0
		return flip(digits, sigDigitIdx-1)
	} else {
		digits[sigDigitIdx] += 1
	}

	return digits

}
