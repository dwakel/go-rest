package api

import (
	"fmt"
	"go-rest/interfaces"
	"go-rest/models"
	"go-rest/utils"
	"log"
	"net/http"
)

type Template struct {
	Logger              *log.Logger
	DataRepository     *interfaces.IDataRepository
	Config              *interfaces.IConfigurationService
	sampleDownstream   *interfaces.ISampleDownstreamService

}

func NewTemplate(logger *log.Logger,
	tokenRepository *interfaces.IDataRepository,
	config *interfaces.IConfigurationService,
	sampleDownstream   *interfaces.ISampleDownstreamService) *Template {
	return &Template{
		logger,
		tokenRepository,
		config,
		sampleDownstream,
	
	}
}

// swagger:route GET templateMethod templateMethod
// Submits token
// Sample request:
//
//       ///     GET /?token=123456789
//
//responses:
//	200: Success
//	500: If the request processing fails due to an exception

// Returns a redirect to TemplateMethod
func (this *Template) TemplateMethod(rw http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	this.Logger.Println(fmt.Sprintf("Token: %s", token))
	//todo: Implement logging into mongodb
	if token == "" {
		this.Logger.Println("Invalid request attempted. No request token found")
		return
	}
	this.Logger.Println(fmt.Sprintf("token i have, %s", (*this.Config).GetToken()))

	if token != ((*this.Config).GetToken()) {
		this.Logger.Println("Invalid token sent")
		return
	}

	//Decrypt 
	data, err := utils.Decryption("some string to decrypt", (*this.Config).GetDecryptionKey())
	if err != nil || data == "" {
		this.Logger.Println("Failed to Decrypt : ", err)
		return
	}

	_, err = (*this.DataRepository).InsertIntoDB(
		models.Data{
			Token:      "",
			TimeStamp:  957687868,
		},
	)
	
	this.Logger.Println("Journey Successful")
	return
}
