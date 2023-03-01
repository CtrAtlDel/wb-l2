package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Process(words []string) map[string][]string {
	result := make(map[string][]string)

	for _, word := range words { // создаем множество из аннограм
		sortedWord := SortWord(word)
		result[sortedWord] = append(result[sortedWord], strings.ToLower(word)) // Добавляем по отсортированому слову ключи
	}

	for key, value := range result { // удаляем подмножества из одного элемента
		if len(value) == 1 {
			delete(result, key)
		}
	}

	for _, v := range result { // сортируем слайс строк
		sort.Strings(v)
	}

	return ChangeMapKeyToValue(result)
}

func ChangeMapKeyToValue(m map[string][]string) map[string][]string { // меняем местами ключ и значение, чтобы ключ был [0]
	res := make(map[string][]string)
	for _, v := range m {
		strings := v
		res[strings[0]] = strings
	}
	return res
}

func SortWord(word string) string { // не обязательно чтобы оно сортировало по алфавиту, достаточно чтобы единообразно это делал
	words := []rune(word)
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })
	return string(words)
}

func main() {

}
