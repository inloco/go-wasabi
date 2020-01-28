package http

import (
	"io/ioutil"
	"net/http"
)

func executeRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	switch {
	case err != nil:
		return nil, err
	case res.StatusCode < int(200) || res.StatusCode >= int(300):
		return nil, handleUnexpectedResponse(res.StatusCode, body)
	}

	return body, nil
}
