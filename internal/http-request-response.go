package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type HttpResp struct {
	statusCode int
	body       string
}

func (res *HttpResp) Print() error {

	prettyJson, err := strconv.ParseBool(flags.flagsMap["pretty"])

	if err != nil {
		return err
	}

	str := "Status Code: " + strconv.Itoa(res.statusCode) + "\n"
	body, err := extractResponseBody(res, prettyJson)

	if err != nil {
		return err
	}

	str += "Response Body:\n" + body
	fmt.Println(str)

	return nil
}

func prettyFormatJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func extractResponseBody(res *HttpResp, prettyJson bool) (string, error) {
	if prettyJson {
		if body, err := prettyFormatJson(res.body); err == nil {
			return body, nil
		}
	}

	return res.body, nil
}
