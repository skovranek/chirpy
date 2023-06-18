package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUserParams(r *http.Request, params interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(params)
	if err != nil {
		err = fmt.Errorf("Error decoding request body: %w", err)
		return err
	}

	return nil
}
