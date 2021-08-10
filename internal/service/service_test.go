package service_test

import (
	// "testing"


	"github.com/andreipimenov/golang-training-2021/internal/service"
	"github.com/andreipimenov/golang-training-2021/internal/mock"


	
	"github.com/stretchr/testify/suite"
	"github.com/rs/zerolog"
)

type serviceTestSuite struct {
	suite.Suite
	service *service.Service
	repo *mock.Repository 
	client *mock.HTTPClient

}



func (suite *serviceTestSuite) SetupTest() {
	suite.repo = &mock.Repository{}
	suite.client = &mock.HTTPClient{}
	suite.service =  service.New(&zerolog.Logger{}, suite.repo, "Token doesn't exist", suite.client)
	
    
}





// func TestSetup(t *testing.T) {
	
// }

