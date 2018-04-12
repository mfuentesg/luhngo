package luhn

import (
	"strconv"
	"testing"
)

func TestInvalidLuhnNumber(t *testing.T) {
	numbers := []string{"00x", "0xff000000", "", "123a1_", "_"}

	for _, number := range numbers {
		_, err := Digit(number)
		if err == nil {
			t.Errorf("Digit(%s), expected nil", number)
		}
	}
}

func TestDigitWithValidNumbers(t *testing.T) {
	numbers := map[string]int{"7": 5, "0": 0, "383": 0, "101099877719": 5}

	for number, expected := range numbers {
		if digit, _ := Digit(number); digit != expected {
			t.Errorf("Digit(%s) must be equals to %d, got %d", number, expected, digit)
		}
	}
}

func TestVerifyWrongCases(t *testing.T) {
	numbers := []string{"73", "01", "3836", "1010998777197", "1"}
	for _, number := range numbers {
		if valid, _ := Verify(number); valid {
			t.Errorf("Verify(%s) must be invalid", number)
		}
	}
}

func TestVerifySuccessCases(t *testing.T) {
	numbers := []string{"75", "00", "3830", "1010998777195", "18"}
	for _, number := range numbers {
		if valid, _ := Verify(number); !valid {
			t.Errorf("Verify(%s) must be valid", number)
		}
	}
}

func TestRandomNumberLength(t *testing.T) {
	length := 10
	randNumber, _ := Rand(length)

	if len(randNumber) != length {
		t.Errorf("number length does not match with %d, got %d", length, len(randNumber))
	}
}

func TestDigitByRandomNumber(t *testing.T) {
	length := 5
	randNumber, _ := Rand(length)
	randDigit, _ := strconv.Atoi(string(randNumber[length-1]))
	randStr := randNumber[0 : length-1]
	digit, _ := Digit(randStr)

	if digit != randDigit {
		t.Errorf("wrong digit expected %d, got %d", randDigit, digit)
	}
}

func TestVerifyByRandomNumber(t *testing.T) {
	length := 10
	randNumber, _ := Rand(length)

	if verified, _ := Verify(randNumber); !verified {
		t.Errorf("wrong random number verification")
	}
}

func TestCompleteByRandomNumber(t *testing.T) {
	length := 10
	randNumber, _ := Rand(length)
	completedNumber, _ := Complete(randNumber[0 : length-2])

	if randNumber == completedNumber {
		t.Errorf("wrong completed number expected %s, got %s", randNumber, completedNumber)
	}
}
