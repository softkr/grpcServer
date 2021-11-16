package model

import (
	"context"
	"fmt"
	pb "gRPCServer/proto"
	"go.mongodb.org/mongo-driver/bson"
)

type UserType struct {
	Guid    string `bson:"guid"`
	Project string `bson:"project"`
}

var UserDB = UserCollection()

func (data *UserType) Insert() interface{} {
	res, err := UserDB.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	fmt.Println(id)
	return id
}

func Users(data *UserType) *UserType {
	return data
}

func FindProject(guid string) pb.ProjectInfo {
	var result UserType
	var project pb.ProjectInfo
	err := UserDB.FindOne(context.TODO(), bson.M{"guid": guid}).Decode(&result)
	// 프로젝트 워치 번호 존재 하지 않으면 개발서버로 실행
	if err != nil {
		fmt.Println(err)
		result.Project = "test"
	}
	projectDB.FindOne(context.TODO(), bson.M{"code": result.Project}).Decode(&project)
	return project
}
