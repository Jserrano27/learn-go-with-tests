package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("Withdraw when insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(25))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if expected != got {
		t.Errorf("Expected %s, but got %s instead", expected, got)
	}
}

func assertError(t *testing.T, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Fatal("WARNING: Expected an error but didn't get one")
	}

	if expected != err {
		t.Errorf("Expected error %q, but got %q instead", expected, err)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("WARNING: Error was not expected but it got one")
	}
}
