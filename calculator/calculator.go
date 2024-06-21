package calculator

import (
	"errors"
	"kata/converter"
	"kata/validator"
	"strconv"
	"strings"
)

var actions = []string{"+", "-", "/", "*"}

func Calculate(str string) (string, error) {
	converter := converter.NewConverter()
	err := errors.New("Panic")
	str = strings.ReplaceAll(str, " ", "")

	actionIndex := -1

	for i := 0; i < len(actions); i++ {
		if strings.Contains(str, actions[i]) {
			actionIndex = i
			break
		}
	}
	if actionIndex == -1 {
		return "", err
	}
	data := strings.Split(str, actions[actionIndex])
	if len(data) > 2 {
		return "", err
	}
	var a int
	var b int
	var result int
	var isRoman bool

	if converter.IsRoman(data[0]) && converter.IsRoman(data[1]) {

		a = converter.ToNumber(data[0])
		b = converter.ToNumber(data[1])
		err = validator.Validate(a, b)
		if err != nil {
			return "", err
		}
		isRoman = true
	} else {

		a, err = strconv.Atoi(data[0])
		if err != nil {
			return "", err
		}
		b, err = strconv.Atoi(data[1])
		if err != nil {
			return "", err
		}
		err = validator.Validate(a, b)
		if err != nil {
			return "", err
		}
	}
	switch actions[actionIndex] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	default:
		result = a / b
	}
	if isRoman {
		return converter.ToRoman(result), nil
	}
	return strconv.Itoa(result), nil
}
