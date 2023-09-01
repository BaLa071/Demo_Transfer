package transactionservices

import (
	"context"
	"fmt"
	"log"
	transactioninterfaces "transaction_grpc/transaction_dal/transaction_interfaces"
	transactionmodels "transaction_grpc/transaction_dal/transaction_models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	client                *mongo.Client
	CustomerCollection    *mongo.Collection
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

func InitTransactionService(client *mongo.Client, customer *mongo.Collection, transaction *mongo.Collection, ctx context.Context) transactioninterfaces.TransactionService {
	return &TransactionService{client, customer, transaction, ctx}
}

// func (a *TransactionService) CreateTransaction(account *transactionmodels.TransactionRequest) (*transactionmodels.TransactionDbResponse, error) {
// 	res, _ := a.TransactionCollection.InsertOne(a.ctx, &account)
// 	var newUser *transactionmodels.TransactionDbResponse
// 	query := bson.M{"_id": res.InsertedID}

// 	err := a.TransactionCollection.FindOne(a.ctx, query).Decode(&newUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return newUser, nil
// }

func (a *TransactionService) Transfer(fromid string, toid string, amt int64) (string, error) {
	session, err := a.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())

	_, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {
		_, err := a.CustomerCollection.UpdateOne(context.Background(),
			bson.M{"customer_id": fromid},
			bson.M{"$inc": bson.M{"balance": -(amt)}})
		if err != nil {
			fmt.Println("Transaction Failed1", err)
			return nil, err
		}
		_, err2 := a.CustomerCollection.UpdateOne(context.Background(), bson.M{"customer_id": toid}, bson.M{"$inc": bson.M{"balance": amt}})

		if err2 != nil {
			fmt.Println("Transaction Failed")
			return nil, err2
		}
		transactionToInsert := transactionmodels.TransactionRequest{
			TransactionId: "T0011",
			FromId:        fromid,
			ToId:          toid,
			Amount:        amt,
		}
		res, _ := a.TransactionCollection.InsertOne(context.Background(), &transactionToInsert)
		var newUser *transactionmodels.TransactionDbResponse
		query := bson.M{"_id": res.InsertedID}

		err3 := a.TransactionCollection.FindOne(a.ctx, query).Decode(&newUser)
		if err3 != nil {
			return nil, err3
		}
		return newUser, nil
	})
	if err != nil {
		return "failed", err
	}
	return "success", nil
}
