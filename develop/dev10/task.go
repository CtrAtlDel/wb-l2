package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	host    *string
	port    *string
	timeout *time.Duration
)

func init() {
	host = flag.String("host", "", "the host to connect to")
	port = flag.String("port", "", "the port to connect to")
	timeout = flag.Duration("timeout", 10*time.Second, "the timeout for connecting to the server")
}

func main() {
	flag.Parse()

	if *host == "" || *port == "" { // Проверим указан ли хост или порт
		flag.Usage()
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", *host, *port)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error cannot close connection %s: %v\n", address, err)
		}
	}(conn)

	go func() { // Создаем поток  для чтения в подключении
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf) // Читаем буффер размером 1024
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading from %s: %v\n", address, err)
				os.Exit(1)
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	sigc := make(chan os.Signal, 1) // Проверяем нет ли вызовов
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	for { // Ждем сигнал из канала  EOF из ОС
		select {
		case <-sigc:
			err := conn.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error cannot close connection %s: %v\n", address, err)
			}
		default:
			buf := make([]byte, 1024)
			n, err := os.Stdin.Read(buf) // Считывание из STDIN
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading from %s: %v\n", address, err)
			}
			_, err = conn.Write(buf[:n]) // Запись в соединение
			if err != nil {
				fmt.Fprintf(os.Stderr, "error writing to %s: %v\n", address, err)
				os.Exit(1)
			}
		}
	}
}
