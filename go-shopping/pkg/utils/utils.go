package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data
	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
