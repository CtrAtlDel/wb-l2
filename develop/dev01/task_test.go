package main

import (
	"testing"

	"github.com/beevik/ntp"
)

func TestMain(t *testing.T) {
	tm, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Errorf("Timeout server:  %v", err)
	}
	if tm.String() == "" {
		t.Errorf("Empty time result from server:  %v", err)
	}
}
