package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func DownloadUrl(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL: %s", urlStr)
	}

	client := http.Client{} // cоздаем клиента

	req, err := http.NewRequest("GET", u.String(), nil) // GET запрос без тела
	if err != nil {
		return fmt.Errorf("cannot create request: %s", err)
	}

	resp, err := client.Do(req) // выполняем запрос
	if err != nil {
		return fmt.Errorf("cannot send request: %s", err)
	}
	defer resp.Body.Close()

	dir, err := createDir(u.Host) // создаем директорию для проекта
	if err != nil {
		return fmt.Errorf("cannot create a directory: %s", err)
	}

	filename := filepath.Join(dir, filepath.Base(u.Path)) //
	if strings.HasSuffix(filename, "/") || strings.HasSuffix(filename, "\\") {
		filename += "index.html"
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	return nil
}

func createDir(host string) (string, error) { // создаем директорию для файлов
	dir := filepath.Join(".", host)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return dir, err
}

func main() {
	// ссылка для примера
	err := DownloadUrl("https://stackoverflow.com/questions/23190311/reverse-a-map-in-value-key-format-in-golang")
	if err != nil {
		log.Fatal(err)
	}
}
