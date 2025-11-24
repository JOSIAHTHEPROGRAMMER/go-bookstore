package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, dst any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}
	return nil
}
