package model

type WalletLoginInput struct {
	Address   string
	Challenge string
	Signature string
}

type WalletLoginOutput struct {
	Id int `json:"id"`
}

type VerifyModel struct {
	Address   string
	Signature string
	Challenge string
}
