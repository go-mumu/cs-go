/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-07
 * File: interest.go
 * Desc: 权益
 */

package interest

import (
	"context"
	"errors"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/service/dal/third"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type Interest struct {
	third.Base
}

func (i *Interest) interestList(ctx context.Context, params map[string]string) []byte {

	queryParams := make(map[string]string)

	queryParams["appId"] = "104"
	queryParams["memberId"] = params["mid"]
	queryParams["interestCode"] = params["interestCode"]

	response, err := i.client(ctx).R().
		SetQueryParams(queryParams).
		SetHeader("h-app-id", queryParams["appId"]).
		Get(config.V.GetString("domain.center") + "/api/interest/query/member/interest/list")

	if err != nil {
		log.Log.InfoContext(ctx, "err", err)
		return nil
	}

	return response.Body()
}

func (i *Interest) client(ctx context.Context) *resty.Client {
	return i.BaseClient(ctx).OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {

		code := jsoniter.Get(resp.Body(), "code").ToInt()
		success := jsoniter.Get(resp.Body(), "success").ToBool()

		log.Log.InfoContext(ctx, "response",
			"code", code,
			"success", success,
			"response", resp.Body(),
			"request_time", resp.Time(),
		)

		if code != 0 || !success {
			return errors.New("response failed")
		}

		return nil
	})
}
