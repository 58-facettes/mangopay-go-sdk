package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

var (
	BaseURL   string
	BasicAuth string
)

// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &http.Client{}

func newRequestAndExecute(method, uri string, param interface{}) (int, []byte, error) {
	body, err := json.Marshal(param)
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

func addQuery(uri string, query ...model.Query) string {
	if query != nil && len(query) > 0 {
		return query[0].URI(uri)
	}
	return uri
}

type RateLimit struct {
	//
	Limit int // X-RateLimit
	// RateLimit int // X-RateLimit
	// RateLimit int // X-RateLimit
	// RateLimit int // X-RateLimit
	LimitRemaining int // X-RateLimit-Remaining
	// LimitRemaining int // X-RateLimit-Remaining
	// LimitRemaining int // X-RateLimit-Remaining
	// LimitRemaining int // X-RateLimit-Remaining
	LimitReset int // X-RateLimit-Reset
	// LimitReset int // X-RateLimit-Reset
	// LimitReset int // X-RateLimit-Reset
	// LimitReset int // X-RateLimit-Reset
}

func GetRateLimit() (*RateLimit, error) {
	r, err := http.NewRequest(http.MethodGet, BaseURL+"clients/", nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", BasicAuth)
	r.Header.Set("Content-Type", "application/json")

	client := DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	reset, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Reset"))
	return &RateLimit{
		Limit:          limit,
		LimitRemaining: remaining,
		LimitReset:     reset,
	}, nil
}
