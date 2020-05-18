package http

import (
	"context"
	"testing"

	"github.com/inloco/go-wasabi/experiments"

	"github.com/inloco/go-wasabi/experiments/fixtures"
	"github.com/inloco/go-wasabi/http/mockserver"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite

	client *HttpClient
}

func (suite *HttpTestSuite) assertEqualExperiment(expected *experiments.Experiment, received *experiments.Experiment) {
	suite.EqualValues(expected.ApplicationName, received.ApplicationName)
	suite.EqualValues(expected.ID, received.ID)
	suite.EqualValues(expected.Label, received.Label)
}

func (suite *HttpTestSuite) SetupTest() {

	s := mockserver.NewTestServer(suite.T())

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
		"user/ID1",
	)

	if suite.NoError(err) {
		suite.Equal(
			mockserver.GenerateAssignmentResponseResource,
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

	suite.EqualValues(experiment.ApplicationName, res.ApplicationName)
	suite.EqualValues(experiment.Label, res.Label)
}

func (suite *HttpTestSuite) TestCreateBucket() {
	buckets := fixtures.Buckets()

	res, err := suite.client.CreateBucket(context.Background(), buckets[0])

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.Require().NotEmpty(res.State)

	suite.EqualValues(buckets[0].AllocationPercent, res.AllocationPercent)
	suite.EqualValues(buckets[0].IsControl, res.IsControl)
	suite.EqualValues(buckets[0].Label, res.Label)
	suite.EqualValues(buckets[0].Payload, res.Payload)
	suite.EqualValues(buckets[0].ExperimentID, res.ExperimentID)
}

func (suite *HttpTestSuite) TestUpdateExperimentState() {
	experiment := fixtures.ExperimentToUpdateState()

	res, err := suite.client.UpdateExperimentState(context.Background(), experiment.ID, experiments.ExperimentStateRunning)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.assertEqualExperiment(experiment, res)
	suite.EqualValues(res.State, experiments.ExperimentStateRunning)
}

func (suite *HttpTestSuite) TestGetExperimentByID() {
	expected := fixtures.Experiment()

	experiment, err := suite.client.GetExperimentByID(
		context.Background(),
		"experiment01",
	)

	if suite.NoError(err) {
		suite.EqualValues(expected.ApplicationName, experiment.ApplicationName)
		suite.EqualValues(expected.Label, experiment.Label)
	}
}

func (suite *HttpTestSuite) TestGetExperimentBuckets() {
	expected := fixtures.Buckets()

	buckets, err := suite.client.GetExperimentBuckets(
		context.Background(),
		"experiment01",
	)

	if suite.NoError(err) {
		suite.ElementsMatch(expected, buckets)
	}
}

func (suite *HttpTestSuite) TestUpdateExperiment() {
	experiment := fixtures.ExperimentToUpdate()

	res, err := suite.client.UpdateExperiment(context.Background(), experiment.ID, experiment)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.Require().NotEmpty(res.ID)

	suite.EqualValues(experiment, res)
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
