package core

import (
	"errors"
	"fmt"
	"github.com/soniah/evaler"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(expression string) (string, error) {
	if expression == "" {
		return "", nil
	}

	result, err := evaler.Eval(replaceRadix(expression))
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid expression")
	}
	return strings.TrimRight(strings.TrimRight(result.FloatString(10), "0"), "."), nil
}

func replaceRadix(expression string) string {
	hex := regexp.MustCompile(`0x[a-zA-Z0-9]*`)
	binary := regexp.MustCompile(`0b[0-1]*`)
	any := regexp.MustCompile(`([a-zA-Z0-9]*)_([0-9]*)`)

	for _, hexMatch := range hex.FindAllStringSubmatch(expression, -1) {
		expression = strings.ReplaceAll(expression, hexMatch[0], convertRadix(hexMatch[0][2:], 16))
	}
	for _, binaryMatch := range binary.FindAllStringSubmatch(expression, -1) {
		expression = strings.ReplaceAll(expression, binaryMatch[0], convertRadix(binaryMatch[0][2:], 2))
	}
	for _, anyMatch := range any.FindAllStringSubmatch(expression, -1) {
		base, err := strconv.Atoi(anyMatch[2])
		if err != nil {
			DebugError(err)
		} else {
			fmt.Println(expression)
			expression = strings.ReplaceAll(expression, anyMatch[0], convertRadix(anyMatch[1], base))
			fmt.Println(expression)
		}
	}

	return expression
}

func convertRadix(input string, base int) string {
	number, err := strconv.ParseInt(input, base, 64)
	if err != nil {
		DebugError(err)
		return input
	}
	return strconv.FormatInt(number, 10)
}
