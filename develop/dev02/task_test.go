package main

import "testing"

func TestUnboxing(t *testing.T) {
	ex1 := "a4bc2d5e"
	res, err := UnBoxing(ex1)
	if res != "aaaabccddddde" {
		t.Errorf("Some problem with algo:  %v", err)
	}

	ex2 := "abcd"
	res, _ = UnBoxing(ex2)
	if res != "abcd" {
		t.Errorf("Some problem with algo:  %v", err)
	}

	ex3 := "45"
	_, err = UnBoxing(ex3)
	if err != nil {
		t.Errorf("Incorrect string:  %v", err)
	}
	ex4 := ""
	_, err = UnBoxing(ex4)
	if err == nil {
		t.Errorf("Incorrect string:  %v", err)
	}
}
