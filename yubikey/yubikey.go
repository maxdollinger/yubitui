package yubikey

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Yubikey struct{}

func InitYubikey() (*Yubikey, error) {
	cmd := exec.Command("which", "ykman")

	err := cmd.Run()
	if err != nil {
		return nil, errors.New("need ykman as depency")
	}

	return &Yubikey{}, nil
}

func (y *Yubikey) ListAccounts() ([]string, error) {
	cmd := exec.Command("ykman", "oath", "accounts", "list")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	accounts := make([]string, 0, 20)
	for line := range strings.SplitSeq(string(output), "\n") {
		str := strings.TrimSpace(line)
		if len(str) > 0 {
			accounts = append(accounts, str)
		}
	}

	return accounts, nil
}

func (y *Yubikey) GenerateCode(account string) (string, error) {
	cmd := exec.Command("ykman", "oath", "accounts", "code", account)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	code := strings.TrimSpace(strings.ReplaceAll(string(output), account, ""))

	return code, nil
}

func (y *Yubikey) Close() {
}

func (y *Yubikey) AddAccount(account string, secret string, digits int) error {
	cmd := exec.Command("ykman", "oath", "accounts", "add", account, secret, "-f")

	output, err := cmd.CombinedOutput()
	if err != nil {
		errStr := y.getErrorLine(string(output))
		return fmt.Errorf("error adding account:\n\n%s\n%s", err, errStr)
	}

	return nil
}

func (y *Yubikey) DeleteAccount(account string) error {
	cmd := exec.Command("ykman", "oath", "accounts", "delete", account, "-f")

	return cmd.Run()
}

func (y *Yubikey) RenameAccount(account string, name string) error {
	cmd := exec.Command("ykman", "oath", "accounts", "rename", account, name, "-f")

	output, err := cmd.CombinedOutput()
	if err != nil {
		errStr := y.getErrorLine(string(output))
		return fmt.Errorf("error renaming account:\n\n%s\n%s", err, errStr)
	}

	return nil
}

func (y *Yubikey) getErrorLine(output string) string {
	err := output
	for line := range strings.Lines(output) {
		if strings.HasPrefix(line, "Error") {
			err = strings.ReplaceAll(line, "Error:", "")
			break
		}
	}

	return strings.TrimSpace(err)
}
