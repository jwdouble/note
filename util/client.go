package util

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"arctron.lib/codec/json"
)

const TimeLayout = "2006-01-02 15:04:05"

var client = &http.Client{Timeout: 30 * time.Second}

func POST(url, token string, body interface{}, data interface{}) ([]byte, error) {
	reqBS := json.MustMarshal(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBS))
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/jsonx")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil // fixme
	}

	respBS, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if data != nil {
		err := json.Unmarshal(respBS, data)
		if err != nil {
			return nil, err
		}
	}

	return respBS, nil
}
