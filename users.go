package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func userHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		GetUser(response, request)
	case "POST":
		CreateUser(response, request)
	}
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user User
	databytes, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(databytes, &user)
	collection := client.Database("appointee").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
	defer request.Body.Close()
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(request.URL.Query().Get("id"))
	var user User
	collection := client.Database("appointee").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(user)
}