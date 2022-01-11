package mocks

import (
	"errors"
	"go-rest/interfaces"
	"go-rest/models"
)

type MockDataRepository struct {
	Success bool
}

func MockNewDataRepo(success bool) interfaces.IDataRepository {
	return &MockDataRepository{Success: success}
}

func (this *MockDataRepository) InsertIntoDB(tokenInsertIntoDB models.Data) (bool, error) {
	if this.Success {
		return this.Success, nil
	}
	return this.Success, errors.New("failed to persist token to database")
}
