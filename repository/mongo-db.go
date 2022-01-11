package repository

import (
	"context"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoDBDriver struct {
	config *Config
	logger *log.Logger
}

type Config struct {
	ConnectionString string `json:"connection_string"`
	Enabled          bool   `json:"enabled"`
	Port             string `json:"port"`
	Database         string `json:"database_name"`
}

func NewMongoDB(config *Config, logger *log.Logger) *MongoDBDriver {
	return &MongoDBDriver{config, logger}
}

func (db *MongoDBDriver) ConnectMongoDB() (*mongo.Database, error) {

	//db.logger.Println(fmt.Sprintf("Connecting to DB, connection string: %v", db.config.ConnectionString))
	clientOptions := options.Client().ApplyURI(db.config.ConnectionString)
	DBClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		db.logger.Println("Mongodb Error on Creating New to mongodb client")
	}
	err = DBClient.Connect(context.Background())
	if err != nil {
		db.logger.Println("Mongodb Error on Connecting to mongodb")
	}
	// GET InsertIntoDBbase Client Here
	ping := DBClient.Ping(context.Background(), readpref.Primary())
	if ping != nil {
		db.logger.Println("Pong: Couldn't  connect to mongodb")
	} else {
		db.logger.Println("Success: Connected to Mongodb")
	}

	return DBClient.Database(db.config.Database), nil
}
