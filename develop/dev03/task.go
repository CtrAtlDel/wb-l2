package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	row          int  // -k — указание колонки для сортировки
	number       int  // -n — сортировать по числовому значению
	reversNumber int  // -r — сортировать в обратном порядке
	isRepetive   bool // -u — не выводить повторяющиеся строки
)

func init() {
	flag.IntVar(&row, "-k", 0, "specifying the column to sort")
	flag.IntVar(&number, "-n", 0, "sort by numeric value")
	flag.IntVar(&reversNumber, "-r", 0, "sort in reverse order")
	flag.BoolVar(&isRepetive, "-u", false, "do not output duplicate lines")
}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file...")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string // collect all lines
	for scanner.Scan() { 
		lines = append(lines, scanner.Text())
	}

	


}

func main() {
	flag.Parse()
	var path = "/Users/ivankvasov/wb/l2/develop/dev03/data/file.txt"
	ReadFile(path)
}
