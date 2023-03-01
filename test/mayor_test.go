package test

import (
	"testing"

	"github.com/gcleroux/IFT630-SCAM/pkg/people"
)

func TestMayor(t *testing.T) {
	got := people.MayorHello()
	expected := "Hello, Mayor!"

	if got != expected {
		t.Errorf("MayorHello() = %s; want %s", got, expected)
	}
}
