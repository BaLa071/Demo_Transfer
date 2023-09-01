package main

import (
	"context"
	"fmt"
	"log"

	pb "transaction_grpc/transaction_protobuff"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransactionServiceClient(conn)
	response, err := client.Transfer(context.Background(), &pb.TransactionRequest{FromId: "C001", ToId: "C002", Amount: 100})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	fmt.Printf("Response: %s\n", response.Status)
}
