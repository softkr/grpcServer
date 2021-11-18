package model

import (
	"context"
	"fmt"
	pb "gRPCServer/proto"

	"go.mongodb.org/mongo-driver/bson"
)

type Users struct {
	Guid    string `bson:"guid"`
	Project string `bson:"project"`
}

var userTable, userDb = Connection{
	Database:   IOT,
	Collection: USERS,
}, userTable.Connect()

func (db *Users) Insert() interface{} {
	res, err := userDb.InsertOne(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	fmt.Println(id)
	return id
}

func (db *Users) Find() pb.ProjectInfo {
	var result Users
	var project pb.ProjectInfo
	err := userDb.FindOne(context.TODO(), bson.M{"guid": db.Guid}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	projectDb.FindOne(context.TODO(), bson.M{"code": result.Project}).Decode(&project)
	return project
}
