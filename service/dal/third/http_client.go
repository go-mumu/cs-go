/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-08
 * File: http_client.go
 * Desc: http client
 */

package third

import (
	"context"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Base struct {
}

// BaseClient request client
func (Base) BaseClient(ctx context.Context) *resty.Client {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	client := resty.New()

	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	client.SetTimeout(time.Second * 3)

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		request.Time = time.Now()

		request.EnableTrace()
		request.SetContext(ctx)

		log.Log.InfoContext(ctx, "request",
			"query_params", request.QueryParam,
			"body", request.Body,
			"request_url", request.URL,
		)

		return nil
	})

	return client
}
