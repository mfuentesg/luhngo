package luhn

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func isValidNumber(lunhNumber string) (bool, error) {
	matched, err := regexp.Match("^[0-9]{1,}$", []byte(lunhNumber))

	if !matched {
		return false, fmt.Errorf("could not match %s: %v", lunhNumber, err)
	}

	return true, nil
}

func randIntn(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func checkSum(luhnNumber string) int {
	var checkSum int

	numberLength := len(luhnNumber) - 1
	for i := 0; i <= numberLength; i++ {
		n, _ := strconv.Atoi(string(luhnNumber[numberLength-i]))

		if i%2 != 0 {
			n = n * 2
		}

		if n > 9 {
			n = n%10 + 1
		}

		checkSum += n
	}
	return checkSum % 10
}

// Digit generate luhn digit for the provided string of digits
func Digit(luhnNumber string) (int, error) {
	valid, err := isValidNumber(luhnNumber)
	if !valid {
		return 0, err
	}

	return (10 - checkSum(fmt.Sprintf("%s0", luhnNumber))) % 10, nil
}

// Rand generates a random valid luhn number
func Rand(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("could not create a random number, length must be greater to 1")
	}

	randStr := strconv.Itoa(randIntn(1, 9))
	for i := 0; i < length-2; i++ {
		randStr = fmt.Sprintf("%s%d", randStr, randIntn(0, 9))
	}

	digit, _ := Digit(randStr)
	return fmt.Sprintf("%s%d", randStr, digit), nil
}

// Verify evalutes if the provided string complies with the Luhn Algorithm
func Verify(luhnNumber string) (bool, error) {
	valid, err := isValidNumber(luhnNumber)

	if !valid {
		return false, err
	}

	return checkSum(luhnNumber)%10 == 0, nil
}

// Complete appends the luhn digit to the provided string of digits
func Complete(luhnNumber string) (string, error) {
	digit, err := Digit(luhnNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%d", luhnNumber, digit), nil
}
