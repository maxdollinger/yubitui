package yubikey

import (
	"fmt"
)

type YubikeyMock struct {
	accounts []string
}

func InitKeyMock() (*YubikeyMock, error) {
	return &YubikeyMock{
		accounts: []string{"jira", "one"},
	}, nil
}

func (y *YubikeyMock) ListAccounts() ([]string, error) {
	return y.accounts, nil
}

func (y *YubikeyMock) GenerateCode(acc string) (string, error) {
	return "abcd1234", nil
}

func (y *YubikeyMock) AddAccount(account string, secret string) error {
	fmt.Printf("Account: %s\nSecret: %s\n", account, secret)
	return nil
}

func (y *YubikeyMock) Close() {
}
