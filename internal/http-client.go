package internal

import (
	"io"
	"net/http"
	"strings"
)

type HttpReqConfig struct {
	verb    string
	url     string
	body    io.Reader
	headers map[string]string
}

func MakeCall(reqConfig *HttpReqConfig) (*HttpResp, error) {
	client := &http.Client{}

	req, err := http.NewRequest(reqConfig.verb, reqConfig.url, reqConfig.body)

	if err != nil {
		return nil, err
	}

	addHeadersToRequest(req, reqConfig.headers)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	httpRes := &HttpResp{res.StatusCode, ""}

	httpRes.body, err = readResponseBody(res)

	return httpRes, nil
}

func readResponseBody(res *http.Response) (string, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func addHeadersToRequest(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Add(strings.Trim(k, " "), strings.Trim(v, " "))
	}
}
