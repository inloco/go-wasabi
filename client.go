package wasabi

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
)

// Client is a (mockable) interface for our Wasabi client.
// Each method is 1:1 to the routes exposed by the Wasabi application.
type Client interface {
	GenerateAssignment(ctx context.Context, experimentLabel string, userID string) (*assignments.Assignment, error)
}
