package transactionmodels

type TransactionRequest struct {
	TransactionId string
	FromId        string
	ToId          string
	Amount        int64
}

type TransactionDbResponse struct {
	TransactionId string
	Status        string
	Message       string
}
