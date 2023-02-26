package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	ex1 := "пятак"
	ex2 := "пятка"
	ex0 := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "голубь"}
	ex0Res := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток"},
	}

	res1 := SortWord(ex1)
	res2 := SortWord(ex2)
	res3 := Process(ex0)

	if res1 != res2 {
		t.Error("Incorrect sorting ...")
	}

	if !reflect.DeepEqual(res3, ex0Res) {
		t.Error("Incorrect result ...")
	}

}
