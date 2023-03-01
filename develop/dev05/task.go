package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Парсим аргументы командной строки
	after := flag.Int("A", 0, "print N lines after match")
	before := flag.Int("B", 0, "print N lines before match")
	context := flag.Int("C", 0, "print N lines before and after match")
	count := flag.Bool("c", false, "print count of matching lines")
	ignoreCase := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "invert match (exclude instead of include)")
	exactMatch := flag.Bool("F", false, "exact match with string, not pattern")
	lineNum := flag.Bool("n", false, "print line number")

	flag.Parse()

	pattern := flag.Arg(0) // получаем паттерн по которому будем искать

	scanner := bufio.NewScanner(os.Stdin) // подключаем сканер

	inMatch := false // Инициализируем переменные для сопоставления
	matchCount := 0
	linesBeforeMatch := make([]string, *before)
	linesAfterMatch := make([]string, *after)
	contextLines := make([]string, *context*2)

	for i := 1; scanner.Scan(); i++ { // Далее делаем поиск в цикле с учетом флагов
		line := scanner.Text() // получаем текст по вызову
		if *ignoreCase {       // если игнорируем регистр
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		if *exactMatch { //точное совпадение со строкой, а не с шаблоном
			if line == pattern {
				inMatch = true
			} else {
				inMatch = false
			}
		} else {
			if strings.Contains(line, pattern) { // строка содержит тот шаблон, который ищем
				inMatch = true
			} else {
				inMatch = false
			}
		}

		if inMatch && !*invert {
			matchCount++

			if *count {
				continue
			}

			if *lineNum { //
				fmt.Printf("%d:", i)
			}

			fmt.Println(line)

			if *before > 0 {
				fmt.Println(linesBeforeMatch)
			}
			if *after > 0 {
				linesAfterMatch = append(linesAfterMatch, line)
				if len(linesAfterMatch) > *after {
					linesAfterMatch = linesAfterMatch[1:]
				}
			}
			if *context > 0 {
				fmt.Println(contextLines)
			}
		} else if !inMatch && *invert {
			if *count {
				matchCount++
			} else {
				if *lineNum {
					fmt.Printf("%d:", i)
				}

				fmt.Println(line)

				if *before > 0 {
					linesBeforeMatch = append(linesBeforeMatch, line)
					if len(linesBeforeMatch) > *before {
						linesBeforeMatch = linesBeforeMatch[1:]
					}
				}
				if *after > 0 {
					fmt.Println(linesAfterMatch)
				}
				if *context > 0 {
					contextLines = append(contextLines, line)
					if len(contextLines) > *context*2 {
						contextLines = contextLines[1:]
					}
				}
			}
		}
	}

	if *count {
		fmt.Println(matchCount)
	}
}