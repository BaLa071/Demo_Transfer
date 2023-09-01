package transactioninterfaces

type TransactionService interface {
	Transfer(fromid string, toid string, amt int64) (string, error)
}
