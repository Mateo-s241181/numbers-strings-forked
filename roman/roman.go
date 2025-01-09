package roman

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {
	// TODO
	return ""
}

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	// TODO
	return ""
}

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {

	//Seperate Digits into ones, tens and hundreds

	Digit_ones := n % 10
	Digit_tens := ((n - (n % 10)) % 100) / 10
	Digit_hundreds := (n - (n % 100)) / 100

	//Let RomanBelow10 compute ones
	OnesString := RomanBelow10(Digit_ones)

	//Let RomanBelow10 compute tens in ones-code and translate it into ten-code
	ones_Code := []rune{'I', 'V', 'X'}
	tens_Code := []rune{'X', 'L', 'C'}

	//Falsestring is the tens string, but still in ones-code
	FalseString := RomanBelow10(Digit_tens)
	TensString := ""

	//if hundreds != 0 => return "C"
	if Digit_hundreds != 0 {
		return string(tens_Code[2])
	}

	//Code in ones_code is converted to tens_code
	for _, el := range FalseString {

		switch el {
		case ones_Code[0]:
			TensString += string(tens_Code[0])

		case ones_Code[1]:
			TensString += string(tens_Code[1])

		case ones_Code[2]:
			TensString += string(tens_Code[2])
		}

	}

	// Result is assembled by merging ones and tens
	return TensString + OnesString
}

func RomanBelow10(n int) string {

	if n < 0 {
		return "FEHLER"
	}

	romanString := ""

	//die Zahl 4 hat in römischen Zahlen auch ein V
	if n >= 4 && n != 10 {
		romanString = "V"
	}

	//Die Zahl 9 hat in römisch auch ein X
	if n >= 9 {
		romanString = "X"
	}

	//Restliche Stellen werden mir Modulorechnung berechnet
	switch n % 5 {

	case 0:
		romanString += ""
	case 1:
		romanString += "I"
	case 2:
		romanString += "II"
	case 3:
		romanString += "III"
	case 4:
		romanString = "I" + romanString
	}

	return romanString
}
