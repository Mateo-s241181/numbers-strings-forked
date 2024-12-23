package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {

	finalString := ""

	firstDigit := n % 10
	secondDigit := ((n - (n % 10)) % 100) / 10
	thirdDigit := (n - (n % 100)) / 100

	if n < 10 {
		finalString = NumberUnderTen(firstDigit)
	}

	if n < 20 && n >= 10 {
		finalString = NumberBetweenTenAndTwenty(firstDigit)
	}

	if n >= 20 {

		stringHundreds := DigitString100(thirdDigit)
		stringTens := DigitString10(secondDigit)
		stringOnes := DigitString1(firstDigit)

		finalString = stringHundreds + stringOnes + stringTens
	}

	return finalString
}
