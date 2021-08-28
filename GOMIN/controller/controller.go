package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main1/mongodb"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Book   string             `json:"book" bson:"book",omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
}

var collection = mongodb.ConnectDB()

func Greet() string {
	return "hello ! Welcome to Api"
}

//GetBooksALl
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var book Book

		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books)
}

//Get

func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book Book

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		fmt.Println("eror")
		return
	}

	json.NewEncoder(w).Encode(book)
}

//Create

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	book = Book{

		Book:   r.FormValue("book"),
		Author: r.FormValue("author"),
	}
	fmt.Println(book)
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		fmt.Println("Error")
		return
	}

	json.NewEncoder(w).Encode(result)
}

//Update
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book Book

	filter := bson.M{"_id": id}

	book = Book{

		Book:   r.FormValue("book"),
		Author: r.FormValue("author"),
	}

	update := bson.D{
		{"$set", bson.D{

			{"book", book.Book},
			{"author", book.Author},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		fmt.Println("Errroor !")
		return
	}

	book.ID = id

	json.NewEncoder(w).Encode(book)
}

//Delete
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		fmt.Println("Errroor !")
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
