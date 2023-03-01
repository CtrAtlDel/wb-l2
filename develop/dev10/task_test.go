package main

import (
	"bytes"
	"net"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	// Поднимаем сервер для теста
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	defer ln.Close()
	// Запускаем клиент в качестве подпроцессора
	cmd := exec.Command("go", "run", "telnet.go", "--timeout=3s", "--host="+ln.Addr().(*net.TCPAddr).IP.String(), "--port="+string(rune(ln.Addr().(*net.TCPAddr).Port)))
	cmdIn, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("failed to get stdin pipe: %v", err)
	}
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start telnet client: %v", err)
	}
	defer cmd.Process.Kill()

	conn, err := ln.Accept() // дождитесь подключения
	if err != nil {
		t.Fatalf("failed to accept connection: %v", err)
	}
	defer conn.Close()

	data := []byte("hello\n") // отправляем дату клиенту
	_, err = cmdIn.Write(data)
	if err != nil {
		t.Fatalf("failed to write to stdin: %v", err)
	}

	time.Sleep(100 * time.Millisecond) // ждем пока клиент отправит данные
	if cmdOut.String() != "hello\n" {
		t.Errorf("expected %q, got %q", "hello\n", cmdOut.String())
	}

	cmd.Process.Signal(syscall.SIGINT) // отправляем SIGINT сигнал завершения клиенту

	if err := cmd.Wait(); err != nil { // ждем завершения клиента
		t.Fatalf("telnet client exited with error: %v", err)
	}
}
