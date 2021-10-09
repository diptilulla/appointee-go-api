package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func welcomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("welcome"))
}

func main() {
	clientOptions := options.Client().
    ApplyURI("mongodb+srv://abcd:abcd@cluster0.u7fpe.mongodb.net/appointee?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, clientOptions)
	
	http.Handle("/users", http.HandlerFunc(userHandler))
	http.Handle("/posts", http.HandlerFunc(postHandler))
	http.Handle("/posts/user", http.HandlerFunc(GetPostsByUser))
	http.Handle("/", http.HandlerFunc(welcomeHandler))
	fmt.Println("Listening")
	http.ListenAndServe("localhost:3000", nil)
}