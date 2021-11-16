package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Files struct {
	Guid         string   `json:"guid" bson:"guid"`
	FileName     string   `json:"filename" bson:"filename"`
	VideoMD5     string   `json:"videomd5" bson:"videomd5"`
	SubFile      []string `json:"subfile" bson:"subfile"`
	SubFileCount int32    `json:"subfilecount" bson:"subfilecount"`
	Value        string
}

var fileTable, fileDb = Connection{
	Database:   IOT,
	Collection: MFILES,
}, fileTable.Connect()

func (db *Files) Insert() {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"guid": db.FileName,
	}
	query := bson.M{
		"$setOnInsert": bson.M{
			"guid":     db.Guid,
			"filename": db.FileName,
			"videomd5": db.VideoMD5,
			"subfile":  db.SubFile,
		},
	}
	fileDb.UpdateOne(context.TODO(), filter, query, opts)
}

func (db *Files) Update() int32 {
	filter := bson.M{
		"subfile": db.Value,
	}
	var result Files
	fileDb.FindOneAndUpdate(context.TODO(), filter, bson.M{
		"$inc": bson.M{
			"subfilecount": 1,
		},
	}).Decode(&result)
	return result.SubFileCount
}

func (db *Files) Find() (data Files) {
	filter := bson.M{
		"subfile": db.Value,
	}
	var result Files
	fileDb.FindOne(context.TODO(), filter).Decode(&result)
	return result
}

func (db *Files) Count() int32 {
	filter := bson.M{
		"videomd5": db.Value,
	}
	var result Files
	fileDb.FindOne(context.TODO(), filter).Decode(&result)
	return result.SubFileCount - int32(len(result.SubFile))
}

func (db *Files) Remove() {
	filter := bson.M{
		"videomd5": db.Value,
	}
	_, err := fileDb.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
