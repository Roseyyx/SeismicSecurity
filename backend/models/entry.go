package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Website  string             `bson:"website,omitempty"`
	Notes    string             `bson:"notes,omitempty"`
}

type File struct {
	Filename string `json:"filename"`
	Key      string `json:"key"`
	Password string `json:"password"`
}
