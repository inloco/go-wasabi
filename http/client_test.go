package http

import (
	"context"
	"testing"

	"github.com/inloco/go-wasabi/experiments/fixtures"
	"github.com/inloco/go-wasabi/http/test"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite

	client *HttpClient
}

func (suite *HttpTestSuite) SetupTest() {

	s := test.NewTestServer(suite.T())

	client := NewHttpClient(
		s.URL,
		"test",
		"admin",
		"admin",
	)

	suite.client = client

}

func (suite *HttpTestSuite) TestGenerateAssignment() {

	assignment, err := suite.client.GenerateAssignment(
		context.Background(),
		"JustToTestGenerateAssignment",
		"userID1",
	)

	if suite.NoError(err) {
		suite.Equal(
			test.GenerateAssignmentResponseResource,
			assignment,
		)
	}

}

func (suite *HttpTestSuite) TestCreateExperiment() {
	experiment := fixtures.Experiment()

	res, err := suite.client.CreateExperiment(context.Background(), experiment)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.Require().NotEmpty(res.ID)
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
