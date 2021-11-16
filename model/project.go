package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectType struct {
	Code      string `bson:"code"`
	Name      string `bson:"name"`
	Url       string `bson:"url"`
	Account   string `bson:"AZURE_STORAGE_ACCOUNT"`
	AccessKey string `bson:"AZURE_STORAGE_ACCESS_KEY"`
}

var projectDB = ProjectCollection()

// 입력
func (data *ProjectType) Insert() interface{} {
	res, err := projectDB.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	fmt.Println(id)
	return id
}

// 리스트
func (data *ProjectType) Search(str string) ProjectType {
	filter := bson.M{
		"code": str,
	}
	var result ProjectType
	projectDB.FindOne(context.Background(), filter).Decode(&result)
	fmt.Println(result)
	return result
}

// 삭제
func (data *ProjectType) Remove(_id string) {
	//deleteResult, _ := project.DeleteOne(context.TODO(), bson.M{"_id": primitive.ObjectIDFromHex(_id)})
	//if deleteResult.DeletedCount == 0 {
	//	fmt.Println("")
	//}
}

func Projects(data *ProjectType) (*ProjectType, *error) {
	return data, nil
}
