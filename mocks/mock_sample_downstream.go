package mocks

import (
	"go-rest/interfaces"
)

type MockSampleDownstreamService struct {
	data string
}

func MockNewSampleDownstreamService(data string) interfaces.ISampleDownstreamService {
	return &MockSampleDownstreamService{data}
}

func (this *MockSampleDownstreamService) FetchDownstreamData(data string) (bool, error) {
	if data == this.data {
		return true, nil
	}
	return false, nil
}
