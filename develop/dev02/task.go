package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

 В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	exmpl := "4d5"
	res, _ := UnBoxing(exmpl)
	fmt.Println(res == "aaaabccddddde")
}

func UnBoxing(s string) (string, error) {
	var result string

	if s == "" {
		return s, errors.New("empty string")
	}
	_, err := strconv.Atoi(s)
	if err == nil {
		return "", err
	}

	for i := 0; i < len(s)-1; i++ {
		number, _ := strconv.Atoi(string(s[i+1]))

		if unicode.IsDigit(rune(s[i])) {
			continue
		}

		digit := unicode.IsDigit(rune(s[i+1]))
		if digit {
			for j := 0; j < number; j++ {
				result += string(s[i])
			}
			continue
		}

		result += string(s[i])

	}
	result += string(s[len(s)-1])

	if s == "" {
		return s, errors.New("incorrect string")
	}

	return result, nil
}
