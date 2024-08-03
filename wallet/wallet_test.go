package wallet

import "testing"

func TestWallet(t *testing.T) {
    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        helper(t, wallet, Bitcoin(10))
    })
    t.Run("withdraw with funds", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))
        assertNoErr(t, err)
        helper(t, wallet, Bitcoin(10))
    })
    t.Run("withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertErr(t, err, ErrInsufficientFunds)
        helper(t, wallet, Bitcoin(20))
    })
}

func helper(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %s wanted %s", got, want)
    }
}
func assertNoErr (t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error but didnt want one")
    }
}
func assertErr (t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("didnt get an error but wanted one")
    }
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
