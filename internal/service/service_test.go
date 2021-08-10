package service_test

import (
	"testing"


	"github.com/stretchr/testify/suite"
	"github.com/rs/zerolog"
)

type serviceTestSuite struct {
	suite.Suite
	service Service
}


func (suite *serviceTestSuite) SetupTest() {
    logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	repo := &mock.Repository{}
	client := &mock.HTTPClient{}
}


func TestSetup(t *testing.T) {
	
}
