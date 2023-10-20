package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(8))
		expected := Bitcoin(12)

		assertBalance(t, wallet, expected)
		assertNoError(t, err)
	})

	t.Run("withdaw insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficentFunds)
	})
}

var assertBalance = func(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

var assertError = func(t *testing.T, err, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("expecting an error but didn't get one")
	}
	if err != expected {
		t.Errorf("expected %q but got %q", expected, err.Error())
	}
}

var assertNoError = func(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("wasn't expecting an error but got one")
	}
}
