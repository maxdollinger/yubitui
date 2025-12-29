package yubikey

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

func (y *YubikeyMock) AddAccount(account string, secret string, length int) error {
	y.accounts = append(y.accounts, account)
	return nil
}

func (y *YubikeyMock) DeleteAccount(account string) error {
	remaining := make([]string, 0, len(y.accounts)-1)
	for _, val := range y.accounts {
		if val != account {
			remaining = append(remaining, val)
		}
	}
	y.accounts = remaining

	return nil
}

func (y *YubikeyMock) Close() {
}
