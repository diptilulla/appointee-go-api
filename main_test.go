package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
type Test struct {
	InsertedID string 
}

var userTestId Test
var postTestId Test

// type User struct {
// 	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
// 	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
// 	Password string             `json:"password,omitempty" bson:"password,omitempty"`
// }

func TestCreateUser(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}
	user := User{
		Name: "testname",
		Email: "tests@gmail.com",
		Password: "test",
	}
	
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	respbody, _:= ioutil.ReadAll(resp.Body)
	if 200 != resp.StatusCode {
		t.Fatal("status code not ok")
	}
	json.Unmarshal(respbody, &userTestId)
}

func TestGetUser(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}
	
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/users?id="+userTestId.InsertedID,nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	respbody, _:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(respbody))
	if 200!= resp.StatusCode {
		t.Fatal("status code not ok")
	}
}

func TestCreatePost(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}
	post := Post{
		UserId: userTestId.InsertedID,
    	Caption: "testcaption",
    	Imageurl: "testurl",
	}
	
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/posts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	respbody, _:= ioutil.ReadAll(resp.Body)
	if 200 != resp.StatusCode {
		t.Fatal("status code not ok")
	}
	json.Unmarshal(respbody, &postTestId)
}

func TestGetPostById(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}
	
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/users?id="+postTestId.InsertedID,nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	respbody, _:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(respbody))
	if 200!= resp.StatusCode {
		t.Fatal("status code not ok")
	}
}

func TestGetPostByUser(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "{ \"status\": \"good\"}")
	}
	
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/users?userid="+userTestId.InsertedID+"&limit=2&page=2",nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	respbody, _:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(respbody))
	if 200!= resp.StatusCode {
		t.Fatal("status code not ok")
	}
}