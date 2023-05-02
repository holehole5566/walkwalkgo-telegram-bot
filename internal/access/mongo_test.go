package db

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnect(t *testing.T) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		t.Fatal("MONGO_URI must be set")
	}

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			t.Fatal(err)
		}
	}()
}
