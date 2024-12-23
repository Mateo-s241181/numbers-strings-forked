package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {

	//Zahl n wird in ihre Ziffern zerlegt
	firstDigit := n % 10
	secondDigit := ((n - (n % 10)) % 100) / 10
	thirdDigit := (n - (n % 100)) / 100

	//Ziffern werden in die entsprechnede Funktion eingesetzt
	stringHundreds := DigitString100(thirdDigit)
	stringTens := DigitString10(secondDigit)
	stringOnes := DigitString1(firstDigit)

	return stringHundreds + stringOnes + stringTens
}
