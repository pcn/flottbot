package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/target/flottbot/model"
	"github.com/target/flottbot/utils"
)

// HTTPResponse ...
type HTTPResponse struct {
	Status int
	Raw    string
	Data   interface{}
}

// HTTPReq handles 'http' actions for rules
func HTTPReq(args model.Action, msg *model.Message) (*HTTPResponse, error) {
	if args.Timeout == 0 {
		// Default HTTP Timeout of 10 seconds
		args.Timeout = 10
	}

	client := &http.Client{
		Timeout: time.Duration(args.Timeout) * time.Second,
	}

	// check the URL string from defined action has a variable, try to substitute it
	url, err := utils.Substitute(args.URL, msg.Vars)
	if err != nil {
		return nil, err
	}

	// TODO: refactor querydata
	// this is a temp fix for scenarios where
	// substitution above may have introduced spaces in the URL
	url = strings.Replace(url, " ", "%20", -1)

	url, payload, err := prepRequestData(url, args.Method, args.QueryData, msg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(args.Method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Close = true

	// Add custom headers to request
	for k, v := range args.CustomHeaders {
		value, err := utils.Substitute(v, msg.Vars)
		if err != nil {
			return nil, err
		}
		req.Header.Add(k, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fields, err := extractFields(bodyBytes)
	if err != nil {
		return nil, err
	}

	result := HTTPResponse{
		Status: resp.StatusCode,
		Raw:    string(bodyBytes),
		Data:   fields,
	}

	return &result, nil
}

// Depending on the type of request we want to deal with the payload accordingly
func prepRequestData(url, method string, data map[string]interface{}, msg *model.Message) (string, io.Reader, error) {
	if len(data) > 0 {
		if method == http.MethodGet {
			query, err := createGetQuery(data, msg)
			if err != nil {
				return url, nil, err
			}
			url = fmt.Sprintf("%s?%s", url, query)
			return url, nil, nil
		}

		query, err := createJSONPayload(data, msg)
		if err != nil {
			return url, nil, err
		}

		return url, strings.NewReader(query), nil
	}

	return url, nil, nil
}

// Unmarshal arbitrary JSON
func extractFields(raw []byte) (interface{}, error) {
	var resp map[string]interface{}

	err := json.Unmarshal(raw, &resp)
	if err != nil {
		return string(raw), nil
	}

	return resp, nil
}

// Create GET query string
func createGetQuery(data map[string]interface{}, msg *model.Message) (string, error) {
	u := url.Values{}
	for k, v := range data {
		subv, err := utils.Substitute(v.(string), msg.Vars)
		if err != nil {
			return "", err
		}
		u.Add(k, subv)
	}
	encoded := u.Encode()                              // uses QueryEscape
	encoded = strings.Replace(encoded, "+", "%20", -1) // replacing + with more reliable %20

	return encoded, nil
}

// Create querydata payload for non-GET requests
func createJSONPayload(data map[string]interface{}, msg *model.Message) (string, error) {
	dataNice := utils.MakeNiceJSON(data)
	str, err := json.Marshal(dataNice)
	if err != nil {
		return "", err
	}

	payload, err := utils.Substitute(string(str), msg.Vars)
	if err != nil {
		return "", err
	}

	return payload, nil
}
