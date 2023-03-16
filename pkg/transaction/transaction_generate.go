package transaction

type Transaction struct {
	Dapp string `json:"dapp"`
}

func NewTransaction(dApp string) *Transaction {
	return &Transaction{
		Dapp: dApp,
	}
}
