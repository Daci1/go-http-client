package internal

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"strconv"
	"strings"
)

var flags *FlagsMap = &FlagsMap{}

func ParseFlags(args []string) (*HttpReqConfig, error) {
	if len(args) < 1 {
		return nil, errors.New("Invalid usage. Run go-http-client -help")
	}

	verb := flag.String("v", "GET", "Request verb. Default is GET.")
	url := flag.String("u", "", "Request URL.")
	body := flag.String("b", "", "Request Body. Default is an empty body.")
	bodyInputFile := flag.String("i", "", "Provides a JSON file for the request body content of a request. Empty path is used by default.")
	headers := flag.String("h", "", "Request Headers.")
	headersInputFile := flag.String("hi", "", "Provides a JSON file for the request headers. Empty path is used by default.")
	pretty := flag.Bool("pretty", false, "Formats the JSON response.")

	flag.Parse()

	if err := addFlagsToMap(
		"verb", strings.ToUpper(*verb),
		"url", *url,
		"body", *body,
		"headers", *headers,
		"bodyInputFile", *bodyInputFile,
		"headersInputFile", *headersInputFile,
		"pretty", strconv.FormatBool(*pretty),
	); err != nil {
		return nil, err
	}

	parsedBody, err := getParsedBody()

	if err != nil {
		return nil, err
	}

	parsedHeaders, err := getParsedHeaders()

	if err != nil {
		return nil, err
	}

	reqConfig := &HttpReqConfig{flags.flagsMap["verb"], flags.flagsMap["url"], strings.NewReader(parsedBody), parsedHeaders}

	return reqConfig, nil
}

func addFlagsToMap(pairs ...string) error {
	if flags.flagsMap != nil {
		return errors.New("Flags map already allocated")
	}

	if len(pairs)%2 != 0 {
		return errors.New("Invalid usage of the method, the arguments array should be even")
	}

	flags.flagsMap = make(map[string]string)

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		flags.flagsMap[key] = string(value)
	}

	return nil
}

func getParsedBody() (string, error) {

	var body string

	if flags.flagsMap["bodyInputFile"] != "" && flags.flagsMap["body"] != "" {
		return "", errors.New("You should pick either input file or CLI input for body")
	}

	if flags.flagsMap["bodyInputFile"] != "" {
		data, err := os.ReadFile(flags.flagsMap["bodyInputFile"])

		if err != nil {
			return "", err
		}

		if json.Valid(data) {
			return string(data), nil
		} else {
			return "", errors.New("Body input file contains invalid JSON.")
		}
	} else {
		body = flags.flagsMap["body"]
		return body, nil
	}
}

func getParsedHeaders() (map[string]string, error) {

	if flags.flagsMap["headersInputFile"] != "" && flags.flagsMap["headers"] != "" {
		return nil, errors.New("You should pick either input file or CLI input for headers")
	}

	if flags.flagsMap["headersInputFile"] != "" {
		data, err := os.ReadFile(flags.flagsMap["headersInputFile"])

		if err != nil {
			return nil, err
		}

		if json.Valid(data) {
			return extractHeaders(string(data))
		} else {
			return nil, errors.New("Headers input file contains invalid JSON.")
		}
	} else {
		return extractHeaders(flags.flagsMap["headers"])
	}
}

func extractHeaders(str string) (map[string]string, error) {
	var headers map[string]string
	if err := json.Unmarshal([]byte(str), &headers); err != nil {
		return nil, err
	}

	return headers, nil
}
