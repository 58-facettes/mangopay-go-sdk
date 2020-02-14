package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/log"
	"github.com/58-facettes/mangopay-go-sdk/model"
)

var (
	// BaseURL is used to be able to switch between the sandbox and the production mode.
	BaseURL string
	// BasicAuth is the base64 hash used in the header.
	BasicAuth string
	// logr is the logger used in order to logs things into the app.
	logr log.Logger
)

// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &http.Client{}

func newRequestAndExecute(method, uri string, payload interface{}) (int, []byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	r, err := http.NewRequest(method, BaseURL+uri, bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	r.Header.Set("Authorization", BasicAuth)
	r.Header.Set("Content-Type", "application/json")

	client := DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, parseErrorResponse(data)
	}
	return resp.StatusCode, data, nil
}

func parseErrorResponse(data []byte) error {
	var errResp model.Error
	err := json.Unmarshal(data, &errResp)
	if err != nil {
		return err
	}
	return errResp
}

func queryURI(uri string, query ...model.Query) string {
	if query != nil && len(query) > 0 {
		return query[0].URI(uri)
	}
	return uri
}

// SetLogger allow to change the default logger.
func SetLogger(l log.Logger) {
	logr = l
}
