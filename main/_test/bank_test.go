package test

import (
	"testing"

	"example.com/main/strc"
)

func TestBank(t *testing.T) {
	bank := strc.SaveBalance(1000_000)

	if bank != 2000_000 {
		t.Errorf("Perhitungan salah want 2000000 got %d", bank)
	}

}

func TestBankWrongType(t *testing.T) {
	bank := strc.SaveBalance(100.000)

	if bank != 1100000 {
		t.Errorf("Error %d", bank)
	}
}
