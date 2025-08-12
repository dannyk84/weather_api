package integrators

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func sendRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		errMsg := fmt.Sprintf("failed to send request | err=%v", err)
		return nil, errors.New(errMsg)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("failed to read response data | err=%v", err)
		return nil, errors.New(errMsg)
	}

	return respData, nil
}
