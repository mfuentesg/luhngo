package luhn

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func isValidNumber(lunhNumber string) (bool, error) {
	matched, err := regexp.Match("^[0-9]{1,}$", []byte(lunhNumber))

	if !matched {
		return false, fmt.Errorf("could not match %s: %v", lunhNumber, err)
	}

	return true, nil
}

func randIntn(min, max int) int {
	return rand.Intn(max-min) + min
}

func checkSum(luhnNumber string) int {
	var reversed string
	numberLength := len(luhnNumber) - 1
	for i := 0; i <= numberLength; i++ {
		reversed += string(luhnNumber[numberLength-i])
	}

	doubled := []int{}
	for i := 0; i <= numberLength; i++ {
		n, _ := strconv.Atoi(string(reversed[i]))
		if i%2 != 0 {
			doubled = append(doubled, n*2)
		} else {
			doubled = append(doubled, n)
		}
	}

	var checkSum int
	for _, digit := range doubled {
		if digit > 9 {
			checkSum += digit%10 + 1
		} else {
			checkSum += digit
		}
	}

	return checkSum % 10
}

// Digit generate luhn digit for the provided string of digits
func Digit(luhnNumber string) (int, error) {
	valid, err := isValidNumber(luhnNumber)
	if !valid {
		return 0, err
	}

	return (10 - checkSum(luhnNumber+"0")) % 10, nil
}

// Rand generates a random valid luhn number
func Rand(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("could not create a random number, length must be greater to 1")
	}

	randStr := strconv.Itoa(randIntn(1, 9))
	for i := 0; i < length-2; i++ {
		randStr += strconv.Itoa(randIntn(0, 9))
	}

	digit, err := Digit(randStr)
	if err != nil {
		return "", err
	}

	return randStr + strconv.Itoa(digit), nil
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
	valid, err := isValidNumber(luhnNumber)

	if !valid {
		return "", err
	}

	digit, err := Digit(luhnNumber)
	if err != nil {
		return "", err
	}

	return luhnNumber + strconv.Itoa(digit), nil
}
