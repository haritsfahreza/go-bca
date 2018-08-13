package bca

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//Client is an interface for making call to BCA API
type Client interface {
	Call(method, path string, body io.Reader, v interface{}) error
	CallRaw(method, path string, headers http.Header, body io.Reader, v interface{}) error
}

//ClientImplementation represents config that used for HTTP client needs
type ClientImplementation struct {
	APIKey     string
	APISecret  string
	OriginHost string
	URL        string
	HTTPClient *http.Client
	LogLevel   int
	Logger     *log.Logger
}

//NewClient is used to initialize new ClientImplementation
func NewClient(cfg Config) ClientImplementation {
	return ClientImplementation{
		APIKey:     cfg.APIKey,
		APISecret:  cfg.APISecret,
		OriginHost: cfg.OriginHost,
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// 0: no logging
		// 1: errors only
		// 2: errors + informational (default)
		// 3: errors + informational + debug
		LogLevel: 2,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

//Call is the implementation for invoking BCA API with its authentication
func (c *ClientImplementation) Call(method, path, accessToken string, body io.Reader, v interface{}) error {
	var headers http.Header
	headers.Add("Authorization", "Bearer "+accessToken)
	headers.Add("Origin", c.OriginHost)
	headers.Add("X-BCA-Key", c.APIKey)

	timestamp := time.Now().Format(time.RFC3339)
	headers.Add("X-BCA-Timestamp", timestamp)

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	signature := generateSignature(c.APISecret, method, path, accessToken, buf.String(), timestamp)
	headers.Add("X-BCA-Signature", signature)

	return c.CallRaw(method, path, "application/json", headers, body, v)
}

//CallRaw is the implementation for invoking API without any wrapper
func (c *ClientImplementation) CallRaw(method, path, contentType string, headers http.Header, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, contentType, headers, body)

	if err != nil {
		return err
	}

	return c.Do(req, v)
}

//NewRequest is used to create new HTTP request of BCA API
func (c *ClientImplementation) NewRequest(method, path, contentType string, headers http.Header, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = c.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		if c.LogLevel > 0 {
			c.Logger.Printf("Cannot create Stripe request: %v\n", err)
		}
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if headers != nil {
		for k, v := range headers {
			for _, line := range v {
				req.Header.Add(k, line)
			}
		}
	}

	return req, nil
}

//Do is used by Call to execute BCA HTTP request and parse the response
func (c *ClientImplementation) Do(req *http.Request, v interface{}) error {
	logLevel := c.LogLevel
	logger := c.Logger

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("BCA response: ", resBody)
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}
	}

	return nil
}

func generateSignature(apiSecret, method, path, accessToken, requestBody, timestamp string) string {
	h := sha256.New()
	h.Write([]byte(requestBody))
	strToSign := method + ":" + path + ":" + accessToken + ":" + strings.ToLower(string(h.Sum(nil))) + ":" + timestamp

	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(strToSign))
	return string(mac.Sum(nil))
}
