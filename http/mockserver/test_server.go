package mockserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(&handler{t})
}

type handler struct {
	t *testing.T
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodGet && req.URL.EscapedPath() == generateAssignmentPath:
		handleGenerateAssignemnt(w, req)
	case req.Method == http.MethodPost && req.URL.EscapedPath() == createExperimentPath:
		handleCreateExperiment(w, req)
	default:
		failureMessage := fmt.Sprintf(
			"test server does not recognize the request %s %s",
			req.Method,
			req.URL.EscapedPath(),
		)

		w.WriteHeader(http.StatusInternalServerError)
		assert.FailNow(h.t, failureMessage)
	}
}
