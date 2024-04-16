package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func iterate(input, res string) (string, error) {
	//if input is empty
	if len(input) == 0 {
		return res, nil
	}
	char := input[0]

	// check if first char is digit
	if char >= 48 && char <= 57 {
		return res, ErrInvalidString
	}

	// if next char is exist
	var nextChar uint8
	if len(input) > 1 {
		nextChar = input[1]
	} else {
		res += string(char)
		return res, nil
	}

	//counter of recurde calls
	howMuchSliced := 0
	// check if next char is digit and not nil
	//log.Println(char, nextChar, nextChar == 92, input)
	if (nextChar >= 49 && nextChar <= 57) || (nextChar == 92 && char == 92) {

		if char == 92 {
			if len(input) >= 3 && (input[2] >= 48 && input[2] <= 57) {
				for i := 0; i < int(input[2]-48); i++ {
					res += string(nextChar)
				}
				howMuchSliced = 3
			} else {
				res += string(nextChar)
				howMuchSliced = 2
			}

		} else {
			for i := 0; i < int(nextChar-48); i++ {
				res += string(char)
			}
			howMuchSliced = 2
		}
	} else if nextChar == 48 {
		res = res[len(res)-2:]
		howMuchSliced = 2
	} else {
		res += string(char)
		howMuchSliced = 1
	}
	return iterate(input[howMuchSliced:], res)
}

func Unpack(input string) (string, error) {
	// Place your code here.
	var res string
	var err error

	res, err = iterate(input, res)
	if err != nil {
		return "", err
	}
	// Place your code here.
	return res, nil
}
