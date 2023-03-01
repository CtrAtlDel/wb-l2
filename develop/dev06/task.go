package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.String("f", "", "select fields") // парсим аргументы командной строки
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only delimited rows")
	flag.Parse()

	selectedFields := make(map[int]bool) // делаем мапу для полей
	if *fields != "" {                   // если строка не пуста
		fieldsList := strings.Split(*fields, ",") // разделяем через запятую
		for _, field := range fieldsList {
			fieldIndex := parseFieldIndex(field) // парсим поля и считаем индекс по которому расположено
			selectedFields[fieldIndex] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin) // Читаем из STDIN
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) { // если только строки с разделителм
			continue
		}
		fields := strings.Split(line, *delimiter)
		var selectedFieldsValues []string
		for i, field := range fields { //
			if len(selectedFields) == 0 || selectedFields[i+1] {
				selectedFieldsValues = append(selectedFieldsValues, field)
			}
		}
		fmt.Println(strings.Join(selectedFieldsValues, *delimiter))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading from STDIN:", err)
		os.Exit(1)
	}
}

func parseFieldIndex(field string) int { // Парсим индекс поля
	fieldIndex := 0
	if field != "" {
		_, err := fmt.Sscanf(field, "%d", &fieldIndex) //сканируем индекс
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid field index %q\n", field)
			os.Exit(1)
		}
		if fieldIndex < 1 {
			fmt.Fprintln(os.Stderr, "field index must be greater than zero")
			os.Exit(1)
		}
	}
	return fieldIndex
}
