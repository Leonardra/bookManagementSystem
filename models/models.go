package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct{
	Id primitive.ObjectID `json:"id" bson:"id"`
	Author primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title string `json:"title" bson:"title"`
	Genre string `json:"genre" bson:"genre"`
	BookDescription string `json:"book_description omitempty" bson:"book_description omitempty"`
	Tags []string `json:"tags" bson:"tags"`
	Quantity int  `json:"quantity" bson:"quantity"`
	ReleaseDate string `json:"release_date" bson:"release_date"`
	DateAdded time.Time `json:"date_added" bson:"date_added"`
}

type Author struct{
	Id primitive.ObjectID  `json:"id" bson:"id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
	Books []Book
}
