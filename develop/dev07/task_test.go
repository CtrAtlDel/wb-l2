package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Parallel()

	c1 := make(chan interface{}) // Тест с двумя каналами, один из которых немедленно закрывается
	close(c1)
	c2 := make(chan interface{})
	result := or(c1, c2)
	if result != c1 {
		t.Errorf("Expected result to be c1, got %v", result)
	}

	c3 := make(chan interface{}) // Тест с двумя каналами, один закрывается с задержкой
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(c3)
	}()
	c4 := make(chan interface{})
	result = or(c3, c4)
	if result != c3 {
		t.Errorf("Expected result to be c3, got %v", result)
	}

	c5 := make(chan interface{}) // Тест с тремя каналами, все закрываются с задержкой
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(c5)
	}()
	c6 := make(chan interface{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(c6)
	}()
	c7 := make(chan interface{})
	go func() {
		time.Sleep(150 * time.Millisecond)
		close(c7)
	}()
	result = or(c5, c6, c7)
	if result != c5 {
		t.Errorf("Expected result to be c5, got %v", result)
	}
}
