package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sockets struct {
	Guid         string
	Addr         string
	Status       string
	Wear         int32
	TakeMedicine int32
}

var watchTable, watchDb = Connection{
	Database:   IOT,
	Collection: WATCHES,
}, watchTable.Connect()

func (db *Sockets) OnOff() bool {
	var (
		filter bson.M
		doc    bson.M
	)
	if db.Status == "on" {
		filter = bson.M{"sn": db.Guid}
		doc = bson.M{"wifi": 1, "addr": db.Addr, "wifiState": db.Status}
	} else {
		filter = bson.M{"addr": db.Addr}
		doc = bson.M{"wifiState": "off", "addr": nil}
	}

	update := bson.D{{"$set", doc}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var result bson.M
	err := watchDb.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func (db *Sockets) Update() bool {
	filter := bson.M{"sn": db.Guid}
	update := bson.D{{"$set", bson.M{"sn": db.Guid, "wear": db.Wear, "takeMedicine": db.TakeMedicine}}}
	var result bson.M
	err := watchDb.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
