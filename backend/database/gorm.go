package database

import(
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func DbConnection() (*mongo.Client){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI("mongodb+srv://mongo:mongo@cluster0.uocwv.mongodb.net/testapp?retryWrites=true&w=majority"))
	if err!=nil {
		fmt.Printf("An error occured")
	}
	return client
}