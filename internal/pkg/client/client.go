package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"errors"

	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/utils"
)

type httpClient struct {
	client *http.Client
}

func NewHttpClient() HttpClient {
	return httpClient{client: &http.Client{Timeout: time.Second * 10}}
}

func (cl httpClient) ExecutePostRequest(ctx context.Context, url string, requestBody any) (string, error) {
	log := utils.GetCurrentLogger(ctx)
	log.Debugf("executing post api request, url: %s...", url)

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Errorf("failed to load request body: %v", err)
		return "", err
	}

	reqJsonBody := string(reqBody)
	log.Debugf("payload: %s", reqJsonBody)

	resp, err := cl.client.Post(url, "application/json", strings.NewReader(reqJsonBody))
	if err != nil {
		log.Errorf("failed to execute request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to load response body: %v", err)
		return "", err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		log.Debugf("successfully execute post api request, data: %s...", string(respBody))
		return string(respBody), nil
	}
	log.Errorf("the server returned, %d with response: %s", resp.StatusCode, string(respBody))
	return "", errors.New(constants.ErrorCodeExternalServerError)
}
