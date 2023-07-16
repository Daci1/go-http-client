package internal

import (
	"encoding/json"
	"errors"
	"flag"
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
	headers := flag.String("h", "", "Request Headers.")
	pretty := flag.Bool("pretty", false, "")

	flag.Parse()

	if err := addFlagsToMap(
		"verb", strings.ToUpper(*verb),
		"url", *url,
		"body", *body,
		"headers", *headers,
		"pretty", strconv.FormatBool(*pretty),
	); err != nil {
		return nil, err
	}

	parsedBody := strings.NewReader(flags.flagsMap["body"])
	var parsedHeaders map[string]string

	if len(*headers) > 0 {
		if err := json.Unmarshal([]byte(flags.flagsMap["headers"]), &parsedHeaders); err != nil {
			return nil, err
		}
	}

	reqConfig := &HttpReqConfig{flags.flagsMap["verb"], flags.flagsMap["url"], parsedBody, parsedHeaders}

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
