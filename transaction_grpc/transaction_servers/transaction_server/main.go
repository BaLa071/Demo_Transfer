package main

import (
	"context"
	"fmt"
	"net"
	transactionconstants "transaction_grpc/transaction_config"
	transactionconfig "transaction_grpc/transaction_config/transaction_config"
	transactionservices "transaction_grpc/transaction_dal/transaction_services"
	transactioncontroller "transaction_grpc/transaction_servers/transaction_controller"

	pro "transaction_grpc/transaction_protobuff"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := transactionconfig.GetCollection(client, "Banking", "Customer")
	TransactionCollection := transactionconfig.GetCollection(client, "Banking", "Transactions")
	transactioncontroller.CustomerService = transactionservices.InitTransactionService(client, customerCollection,TransactionCollection, context.Background())
}

func main() {
	mongoclient, err := transactionconfig.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", transactionconstants.Port)
	if err != nil {
		fmt.Println("failed to listen", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterTransactionServiceServer(s, &transactioncontroller.RPCServer{})

	fmt.Println("Server listenting on", transactionconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve", err)
	}
}
