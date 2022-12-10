package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)

	if want != got {
		t.Errorf("got %d want %d", got, want)
	}
}
