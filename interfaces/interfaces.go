package interfaces

import "go-rest/models"

type IDataRepository interface {
	InsertIntoDB(tokenInsertIntoDB models.Data) (bool, error)
}

type IConfigurationService interface {
	GetToken() string
	GetDecryptionKey() string
	GetDownstreamUrl() string
	DownstreamApiKey() string
}


type ISampleDownstreamService interface {
	FetchDownstreamData(data string) (bool, error)
}
