package main

import (
	"context"
	"encoding/json"
	"strconv"

	// "io/ioutil"

	"net/http"

	// "strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Post struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId   string 			`json:"userId,omitempty" bson:"userId,omitempty"`
	Caption  string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Imageurl string             `json:"imageurl,omitempty" bson:"iurl,omitempty"`
	Stamp    string             `json:"stamp,omitempty" bson:"tstamp,omitempty"`
}


func postHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		GetPostById(response, request)
	case "POST":
		CreatePost(response, request)
	}
}

func CreatePost(response http.ResponseWriter, request *http.Request) {
	Time := time.Now().UTC()
	t := Time.String()
	response.Header().Add("content-type", "application/json")
	var post Post
	json.NewDecoder(request.Body).Decode(&post)
	post.Stamp = t
	collection := client.Database("appointee").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
	defer request.Body.Close()
}

func GetPostById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(request.URL.Query().Get("id"))
	var post Post
	collection := client.Database("appointee").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Post{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

func GetPostsByUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := request.URL.Query()
	userid := params.Get("userid")
	var posts []Post
	page, _ := strconv.Atoi(params.Get("page"))
	limit, _ := strconv.Atoi(params.Get("limit"))

	collection := client.Database("appointee").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{}
	collection.CountDocuments(ctx, filter)
	findOptions := options.Find()
	findOptions.SetSkip((int64(page) - 1) * int64(limit))
	findOptions.SetLimit(int64(limit))

	cursor, err := collection.Find(ctx, Post{UserId: userid}, findOptions)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}