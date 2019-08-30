package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/JustSomeHack/git_cli/models"
)

// HTTPClient interface
type HTTPClient interface {
	Delete() ([]byte, error)
	Get() ([]byte, error)
	Patch(body io.Reader) ([]byte, error)
	Post(body io.Reader) ([]byte, error)
	Put(body io.Reader) ([]byte, error)
}

type hTTPClient struct {
	Params *models.HTTPParams
}

// NewHTTPClient returns a new http client
func NewHTTPClient(params *models.HTTPParams) HTTPClient {
	return &hTTPClient{
		Params: params,
	}
}

// Delete sends DELETE request
func (h *hTTPClient) Delete() ([]byte, error) {
	request, err := http.NewRequest("DELETE", h.Params.URL, nil)
	if err != nil {
		return nil, err
	}

	return h.sendRequest(request)
}

// Get sends GET request
func (h *hTTPClient) Get() ([]byte, error) {
	request, err := http.NewRequest("GET", h.Params.URL, nil)
	if err != nil {
		return nil, err
	}

	return h.sendRequest(request)
}

// Patch sends PATCH request
func (h *hTTPClient) Patch(body io.Reader) ([]byte, error) {
	request, err := http.NewRequest("PATCH", h.Params.URL, body)
	if err != nil {
		return nil, err
	}
	return h.sendRequest(request)
}

// Post sends POST request
func (h *hTTPClient) Post(body io.Reader) ([]byte, error) {
	request, err := http.NewRequest("POST", h.Params.URL, body)
	if err != nil {
		return nil, err
	}
	return h.sendRequest(request)
}

// Put sends PUT request
func (h *hTTPClient) Put(body io.Reader) ([]byte, error) {
	request, err := http.NewRequest("PUT", h.Params.URL, body)
	if err != nil {
		return nil, err
	}
	return h.sendRequest(request)
}

func (h *hTTPClient) setupRequest(request *http.Request) {
	if h.Params.ContentType != "" {
		request.Header.Set("Content-Type", h.Params.ContentType)
	}
	if request.Header.Get("Content-Type") == "" {
		request.Header.Set("Content-Type", "application/json")
	}
	if request.Header.Get("Accept") == "" {
		request.Header.Set("Accept", "application/json")
	}
	if h.Params.AuthorizationBearer != "" {
		request.Header.Set("Authorization", "bearer "+h.Params.AuthorizationBearer)
	}
	if h.Params.AuthorizationKey != "" {
		request.Header.Set("Authorization", "key="+h.Params.AuthorizationKey)
	}
	if h.Params.AuthorizationToken != "" {
		request.Header.Set("Authorization", "token "+h.Params.AuthorizationToken)
	}
	for key, value := range h.Params.Headers {
		request.Header.Set(key, value)
	}
	if h.Params.BasicAuthUser != "" && h.Params.BasicAuthPass != "" {
		request.SetBasicAuth(h.Params.BasicAuthUser, h.Params.BasicAuthPass)
	}

	q := request.URL.Query()
	for key, value := range h.Params.Queries {
		q.Add(key, value)
	}
	if h.Params.URLAccessToken != "" {
		q.Add("access_token", h.Params.URLAccessToken)
	}
	request.URL.RawQuery = q.Encode()
}

func (h *hTTPClient) sendRequest(request *http.Request) ([]byte, error) {
	h.setupRequest(request)

	var client *http.Client
	if h.Params.Proxy != "" {
		proxy, err := url.ParseRequestURI(h.Params.Proxy)
		if err != nil {
			return nil, fmt.Errorf("Invalid Proxy URL found '%s'", h.Params.Proxy)
		}
		t := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Transport: t,
			Timeout:   time.Duration(h.Params.Timeout) * time.Second,
		}
	} else {
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Transport: t,
			Timeout:   time.Duration(h.Params.Timeout) * time.Second,
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Received unexpected status code %d:%s", response.StatusCode, response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
