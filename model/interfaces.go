package model

type YubiKey interface {
	ListAccountsI
	GenerateCodeI
	AddAccountI
	DeleteAccountI
	RenameAccountI
	Close()
}

type RenameAccountI interface {
	RenameAccount(string, string) error
}

type DeleteAccountI interface {
	DeleteAccount(string) error
}

type AddAccountI interface {
	AddAccount(string, string, int) error
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

type PasteI interface {
	Paste() string
}
