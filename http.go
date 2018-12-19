package wasabi

import (
	"github.com/inloco/go-wasabi/http"
)

func NewHttpClient(address, login, password string) Client {
	return http.NewHttpClient(address, login, password)
}
