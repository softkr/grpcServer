package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	FileName string `bson:"filename"`
	Project  bool   `bson:"project"`
	Storage  bool   `bson:"storage"`
	Remove   bool   `bson:"remove"`
	Api      bool   `bson:"api"`
}

var todoTable, todoDb = Connection{
	Database:   IOT,
	Collection: CHECKLIST,
}, todoTable.Connect()

func (db *Todo) Insert() interface{} {
	defer func() {
		if err := recover();err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	res, err := todoDb.InsertOne(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	return id
}

func (db *Todo) Update() {
	result, err := todoDb.UpdateOne(
		context.TODO(),
		bson.M{"filename": db.FileName},
		bson.D{
			{"$set", db},
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func (db *Todo) Delete() {
	result, err := todoDb.DeleteOne(
		context.TODO(),
		bson.M{"filename": db.FileName},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
