package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/58-facettes/mangopay-go-sdk/data"
	"github.com/58-facettes/mangopay-go-sdk/log"
	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/google/uuid"
)

var (
	// BaseURL is used to be able to switch between the sandbox and the production mode.
	BaseURL string
	// BasicAuth is the base64 hash used in the header.
	BasicAuth string
	// logr is the logger used in order to logs things into the app.
	logr log.Logger
	// UseBasicAuth checks if the request use oAuth or BasicAuth.
	UseBasicAuth bool
	// UseIdempotency checks if we use or not Idempotency key.
	UseIdempotency bool
	// DB is used to store the idempotency key.
	DB data.Manager
)

// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &http.Client{}

func newRequestAndExecute(method, uri string, payload interface{}) (int, []byte, error) {
	url := BaseURL + uri
	// create the request.
	var r *http.Request
	var err error
	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		r, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	} else {
		r, err = http.NewRequest(method, url, nil)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}
	// add basicAuth or oAuth token to the header.
	err = setAccessHeader(r)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	r.Header.Set("Content-Type", "application/json")

	// add idempotency key.
	if UseIdempotency {
		key := uuid.New().String()
		saveKey(key, url)
		r.Header.Set("Idempotency-Key", key)
	}

	// execute the request.
	client := DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	// read the response.
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

func queryURI(uri string, query *model.Query) string {
	if query != nil {
		return query.URI(uri)
	}
	return uri
}

// SetLogger allow to change the default logger.
func SetLogger(l log.Logger) {
	logr = l
}

var mx sync.Mutex

func setAccessHeader(r *http.Request) error {
	if UseBasicAuth {
		r.Header.Set("Authorization", BasicAuth)
		return nil
	}
	if oauthAccessToken.ExpiresAt < time.Now().Unix() {
		r.Header.Set("Authorization", "Bearer "+oauthAccessToken.AccessToken)
		return nil
	}
	tok, err := newTokenRequest()
	if err != nil {
		return err
	}
	tok.ExpiresAt = tok.ExpiresIn + time.Now().Unix()
	mx.Lock()
	oauthAccessToken = tok
	mx.Unlock()
	r.Header.Set("Authorization", "Bearer "+oauthAccessToken.AccessToken)
	return nil
}

func newTokenRequest() (*accessTokenResponse, error) {
	r, err := http.NewRequest(http.MethodPost, BaseURL+"oauth/token/", nil)
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tokResponse accessTokenResponse
	err = json.Unmarshal(data, &tokResponse)
	if err != nil {
		return nil, err
	}
	return &tokResponse, nil
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"exires_in"`
	ExpiresAt   int64  `json:"_"`
}

var oauthAccessToken *accessTokenResponse

func saveKey(key, url string) {
	go func() {
		t := time.NewTimer(5 * time.Second)
		go func() {
			for {
				select {
				case <-t.C:
					return
				}
			}
		}()
		err := DB.SaveIdempotencyKey(key, url)
		if err != nil {
			logr.Errorf("mangopay: try to save the idempotency key got error:", err)
		}
	}()
}
