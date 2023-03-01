package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainn(m *testing.T) {
	input := "Go is an Open source programming language\n"
	input += "Developed at Google in 2007\n"
	input += "Go is statically typed, Compiled and easy to use.\n"
	input += "go is now being used by many companies around the world\n"
	input += "Go is awesome\n"
	input += "GO is awesome too\n"

	args := []string{"-c", "-i", "-F", "Go"} // Тест с точным совпадением, флагом подсчета и флагом игнорирования регистра
	out := "4\n"
	testRun(m, args, input, out)

}

func testRun(t *testing.T, args []string, input string, expected string) {

	oldStdout := os.Stdout // перенаправляем вывод
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", "-test.run=TestHelperProcess"} // здесь задаем аргументы для зхапуска программы
	os.Args = append(os.Args, args...)
	main()

	w.Close() // Закрываем запись и читаем то, что в output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	if buf.String() != expected { // Проверяем соответствует ли вывод ожидаемому значению
		t.Errorf("Unexpected output: expected %q, got %q", expected, buf.String())
	}
}
