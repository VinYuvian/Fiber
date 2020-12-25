package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func DbConnection() (*mongo.Client){
	db_name:=os.Getenv("DB_NAME")
	password:=os.Getenv("DB_PASSWORD")
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	if cancel==nil{
		fmt.Printf("An error occured cancel")
	}
	connString:=fmt.Sprintf("mongodb+srv://mongo:%s@cluster0.uocwv.mongodb.net/%s?retryWrites=true&w=majority",password,db_name)
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(connString))
	if err!=nil {
		fmt.Printf("An error occured creating client connection")
	}
	return client
}