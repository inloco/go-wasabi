package http

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite

	client *HttpClient
}

func (suite *HttpTestSuite) SetupTest() {

	client := NewHttpClient(
		"http://wasabi-for-apps.ubee.in/api/",
		"develop",
		"admin",
		"admin",
	)

	suite.client = client

}

func (suite *HttpTestSuite) TestGenerateAssignment() {

	assignment, err := suite.client.GenerateAssignment(
		context.Background(),
		"Develop",
		"user_test_01",
	)

	spew.Dump(assignment, err)

	suite.FailNow("")
}

func (suite *HttpTestSuite) TestCreateExperiment() {

}

func (suite *HttpTestSuite) TestGetExperiments() {

}

func (suite *HttpTestSuite) TestGetExperimentByID() {

}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
