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

	// create a directory for file
	dir, err := createDir(u.Host)
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

	//if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
	//	links, err := extractLinks(resp.Body, u)
	//	if err != nil {
	//		return fmt.Errorf("error extracting links: %v", err)
	//	}
	//	for _, link := range links {
	//		if link.Host == "" {
	//			link.Host = u.Host
	//		}
	//		if link.Scheme == "" {
	//			link.Scheme = u.Scheme
	//		}
	//		err := downloadURL(link.String())
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "error downloading %s: %v\n", link, err)
	//		}
	//	}
	//}

	return nil
}

func createDir(host string) (string, error) { // Создаем директорию для файлов
	dir := filepath.Join(".", host)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return dir, err
}

func main() {
	err := DownloadUrl("https://stackoverflow.com/questions/23190311/reverse-a-map-in-value-key-format-in-golang")
	if err != nil {
		log.Fatal(err)
	}
}
