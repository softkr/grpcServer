package model

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	IOT       = "iot"
	MFILES    = "mfiles"
	PROJECTS  = "projects"
	USERS     = "users"
	WATCHES   = "watches"
	CHECKLIST = "checklist"
)

type Connection struct {
	Database   string
	Collection string
}

var (
	mgoCli     *mongo.Client
	client     = GetMgoCli()
	collection *mongo.Collection
)

func initEngine() {
	var err error

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func (db *Connection) Connect() *mongo.Collection {
	collection = client.Database(db.Database).Collection(db.Collection)
	if db.Database == "iot" {
		collection.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "filename", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
	}
	return collection
}
