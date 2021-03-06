package luhn

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestCheckSum(t *testing.T) {
	numbers := map[string]int{"70": 5, "00": 0, "3830": 0, "1010998777190": 5}
	for number, expected := range numbers {
		if sum := checkSum(number); sum != expected {
			t.Errorf("checkSum(%s) must be %d, got %d", number, expected, sum)
		}
	}
}

func TestRandInt(t *testing.T) {
	// nolint
	min, max := 1, rand.Int()
	for i := 0; i < 1000; i++ {
		if randInt := randIntn(min, max); randInt < min || randInt > max {
			t.Errorf("random int %d is out range [%d,%d]", randInt, min, max)
		}
	}
}

func TestValidNumber(t *testing.T) {
	numbers := []string{"00x", "0xff000000", "", "123a1_", "_"}

	for _, number := range numbers {
		if valid, _ := isValidNumber(number); valid {
			t.Errorf("isValidNumber(%s) must be invalid", number)
		}
	}

	numbers = []string{"006", "1", "123", "16666"}

	for _, number := range numbers {
		if valid, _ := isValidNumber(number); !valid {
			t.Errorf("isValidNumber(%s) must be valid", number)
		}
	}
}

func TestRandom(t *testing.T) {
	length := 10
	if randNumber, _ := Rand(length); len(randNumber) != length {
		t.Errorf("number length does not match with %d, got %d", length, len(randNumber))
	}

	_, err := Rand(0)

	if err == nil {
		t.Errorf("random number must have at least 1 chars")
	}
}

func TestDigit(t *testing.T) {
	numbers := []string{"00x", "0xff000000", "", "123a1_", "_"}
	for _, number := range numbers {
		if _, err := Digit(number); err == nil {
			t.Errorf("Digit(%s), expected nil", number)
		}
	}

	validNumbers := map[string]int{"7": 5, "0": 0, "383": 0, "101099877719": 5}
	for number, expected := range validNumbers {
		if digit, _ := Digit(number); digit != expected {
			t.Errorf("Digit(%s) must be equals to %d, got %d", number, expected, digit)
		}
	}

	length := 5
	randNumber, _ := Rand(length)
	randDigit, _ := strconv.Atoi(string(randNumber[length-1]))
	randStr := randNumber[0 : length-1]

	if digit, _ := Digit(randStr); digit != randDigit {
		t.Errorf("wrong digit expected %d, got %d", randDigit, digit)
	}
}

func TestVerify(t *testing.T) {
	numbers := []string{"73", "01", "3836", "1010998777197", "1", "x"}
	for _, number := range numbers {
		if valid, _ := Verify(number); valid {
			t.Errorf("Verify(%s) must be invalid", number)
		}
	}

	numbers = []string{"75", "00", "3830", "1010998777195", "18"}
	for _, number := range numbers {
		if valid, _ := Verify(number); !valid {
			t.Errorf("Verify(%s) must be valid", number)
		}
	}

	length := 10
	randNumber, _ := Rand(length)

	if verified, _ := Verify(randNumber); !verified {
		t.Errorf("wrong random number verification")
	}
}

func TestComplete(t *testing.T) {
	if _, err := Complete("x"); err == nil {
		t.Error("expected invalid error")
	}

	length := 10
	randNumber, _ := Rand(length)

	if completedNumber, _ := Complete(randNumber[0 : length-2]); randNumber == completedNumber {
		t.Errorf("wrong completed number expected %s, got %s", randNumber, completedNumber)
	}
}

func BenchmarkCheckSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkSum("10109987771900")
	}
}

func BenchmarkRandInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// nolint
		min, max := 1, rand.Int()
		randIntn(min, max)
	}
}

func BenchmarkRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Rand(50)
	}
}

func BenchmarkVerify(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r, _ := Rand(50)
		_, _ = Verify(r)
	}
}

func BenchmarkDigit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Digit("101099877719")
	}
}

func BenchmarkComplete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Complete("101099877719")
	}
}
