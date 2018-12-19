package wasabi

import (
	"github.com/inloco/go-wasabi/http"
)

func NewHttpClient(address, applicationName, login, password string) Client {
	return http.NewHttpClient(address, applicationName, login, password)
}
