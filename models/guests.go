package models

import "gopkg.in/mgo.v2/bson"

type Guests struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	FName     string        `bson:"fname" json:"firstname"`
	LName     string        `bson:"lname" json:"lastname"`
	Attending string        `bson:"attending" json:"attending"`
	Children  int           `bson:"children" json: "children"`
	Adults    int           `bson:"adults" json:"adults"`
}
