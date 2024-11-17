package client

import (
	"context"
)

type HttpClient interface {
	ExecutePostRequest(ctx context.Context, url string, payload any) (string, error)
}
