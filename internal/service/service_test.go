package service_test

import (
	"testing"
	"time"
	"fmt"


	"github.com/andreipimenov/golang-training-2021/internal/service"
	"github.com/andreipimenov/golang-training-2021/internal/mock"
	"github.com/andreipimenov/golang-training-2021/internal/model"


	
	"github.com/stretchr/testify/suite"
	"github.com/rs/zerolog"
)

type serviceTestSuite struct {
	suite.Suite
	service *service.Service
	repo *mock.Repository 
	client *mock.HTTPClient

}

const (
	ticker      = "AAPL"
	date        = "2021-07-26"
	// invalidDate = "2021-07-269999999999"
	// apiFormat   = "http://127.0.0.1:8080/price/%s/%s"
)



func (suite *serviceTestSuite) SetupTest() {
	suite.repo = &mock.Repository{}
	suite.client = &mock.HTTPClient{}
	suite.service =  service.New(&zerolog.Logger{}, suite.repo, "Token doesn't exist", suite.client)
	
    
}

func TestService(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}



func (suite *serviceTestSuite) TestServiceValid() {
	price := &model.Price{
		Open:  "99.9",
		High:  "99.9",
		Low:   "99.9",
		Close: "99.9",
	}

	d, err := time.Parse("2021-07-26", date)
	key := fmt.Sprintf("%s_%s", ticker, d)
	suite.NoError(err)
	
	suite.repo.On("Load", key).Once().Return(price, true)
	res, err := suite.service.GetPrice(ticker, d)
	suite.NoError(err)
	suite.Equal(res, price)


}





// func TestSetup(t *testing.T) {
	
// }

