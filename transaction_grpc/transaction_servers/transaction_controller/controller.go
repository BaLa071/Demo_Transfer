package transactioncontroller

import (
	"context"
	"fmt"
	transactioninterfaces "transaction_grpc/transaction_dal/transaction_interfaces"
	pro "transaction_grpc/transaction_protobuff"
)

type RPCServer struct {
	pro.UnimplementedTransactionServiceServer
}

var (
	CustomerService transactioninterfaces.TransactionService
)

func (c *RPCServer) Transfer(ctx context.Context, req *pro.TransactionRequest) (*pro.TransactionResponse, error) {
	// dbCustomer := &transactionmodels.TransactionRequest{}
	res, err := CustomerService.Transfer(req.FromId, req.ToId, req.Amount)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}else{
		responseProfile:=&pro.TransactionResponse{

		}
		return responseProfile,nil
	}
}
