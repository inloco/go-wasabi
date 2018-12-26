package http

import (
	"context"
	"testing"

	http_test "github.com/inloco/go-wasabi/http/test"

	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite

	client *HttpClient
}

func (suite *HttpTestSuite) SetupTest() {

	s := http_test.NewTestServer(suite.T())

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
			http_test.GenerateAssignmentResponseResource,
			assignment,
		)
	}

}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
