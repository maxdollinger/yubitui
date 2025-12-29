package model

type YubiKey interface {
	ListAccountsI
	GenerateCodeI
	AddAccount(string, string) error
	Close()
}

type ListAccountsI interface {
	ListAccounts() ([]string, error)
}

type CopyToClipboardI interface {
	Copy(string) error
}

type GenerateCodeI interface {
	GenerateCode(string) (string, error)
}
