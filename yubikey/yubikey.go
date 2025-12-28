package yubikey

import (
	"fmt"

	yk "cunicu.li/go-iso7816/devices/yubikey"
	"cunicu.li/go-iso7816/drivers/pcsc"
	"cunicu.li/go-ykoath/v2"
	"github.com/ebfe/scard"
)

type Yubikey struct {
	card *ykoath.Card
}

func InitYubikey() (*Yubikey, error) {
	ctx, err := scard.EstablishContext()
	if err != nil {
		return nil, fmt.Errorf("failed to establish context: %w", err)
	}

	sc, err := pcsc.OpenFirstCard(ctx, yk.HasOATH, true)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pcsc device: %w", err)
	}

	c, err := ykoath.NewCard(sc)
	if err != nil {
		return nil, fmt.Errorf("failed to initalize yubikey: %w", err)
	}

	if _, err = c.Select(); err != nil {
		return nil, fmt.Errorf("failed to initalize OATH session: %w", err)
	}

	return &Yubikey{
		card: c,
	}, nil
}

func (y *Yubikey) ListAccounts() ([]string, error) {
	names, err := y.card.List()
	if err != nil {
		return nil, fmt.Errorf("error invoking list cmd: %w", err)
	}

	accountNames := make([]string, 0, 10)
	for _, name := range names {
		accountNames = append(accountNames, name.Name)
	}

	return accountNames, nil
}

func (y *Yubikey) GenerateCode(account string) (string, error) {
	calc, err := y.card.Calculate(account)
	if err != nil {
		return "", fmt.Errorf("failed to calculate code for account %s: %w", account, err)
	}

	return calc, nil
}

func (y *Yubikey) Close() {
	err := y.card.Close()
	if err != nil {
		fmt.Println(err)
	}
}
