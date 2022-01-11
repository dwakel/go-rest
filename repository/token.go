package repository

import (
	"context"
	"go-rest/interfaces"
	"go-rest/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)
//MONGO DB Implementation
type DataRepository struct {
	Logger *log.Logger
	Repo   *mongo.Database
}

func NewDataRepo(tr *DataRepository) interfaces.IDataRepository {
	return &DataRepository{tr.Logger, tr.Repo}
}

func (this *DataRepository) InsertIntoDB(tokenInsertIntoDB models.Data) (bool, error) {

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	_, err := this.Repo.Collection(DATA_TABLE_NAME).InsertOne(ctx, tokenInsertIntoDB)

	if err != nil {
		this.Logger.Println(err)
		return false, err
	}

	this.Logger.Println("Successfully inserted token data into database")

	return true, nil
}
