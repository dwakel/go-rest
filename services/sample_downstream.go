package services

import (
	"encoding/json"
	"fmt"
	"go-rest/interfaces"
	"go-rest/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type SampleDownstreamService struct {
	Logger *log.Logger
	Config *interfaces.IConfigurationService
}

const VALID_DATA = "VALID_DATA"

func NewSampleDownstreamService(cms *SampleDownstreamService) interfaces.ISampleDownstreamService {
	return &SampleDownstreamService{cms.Logger, cms.Config}
}

func (this *SampleDownstreamService) FetchDownstreamData(param string) (bool, error) {
	var response models.SampleDownstreamResponse
	url := fmt.Sprintf((*this.Config).GetDownstreamUrl())

	client := http.Client{
		Timeout: 15 * time.Second,
	}
	this.Logger.Println("Calling Downstream API service")
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		this.Logger.Println(err)
		return false, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", (*this.Config).DownstreamApiKey())

	resp, err := client.Do(request)
	if err != nil {
		this.Logger.Println(err)
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		this.Logger.Println("[GET] ioutil.ReadAll Error: ", err)
		return false, err
	}
	err = json.Unmarshal(body, &response)

	if err != nil {
		this.Logger.Println("[GET] json.Unmarshal Error: %v", err)
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		this.Logger.Println(err)
		return false, err
	}

	//Check service type
	data := response.Data
	this.Logger.Println("Data: ", data)
	if data == "" {
		this.Logger.Println("Failed to fetch type from Downstream")
	}

	if strings.ToLower(data) != VALID_DATA {
		return false, nil
	}

	return true, nil
}
