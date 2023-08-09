/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-09
 * File: funcs.go
 * Desc: do interest function
 */

package interest

import (
	"context"
	"github.com/go-mumu/cs-go/library/config"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

// UserInterest 用户权益
func (i *Interest) UserInterest(ctx context.Context, mid string) map[string]string {

	resp := i.interestList(ctx, map[string]string{
		"mid":          mid,
		"interestCode": config.C.Interest.Code,
	})

	anyJson := jsoniter.Get(resp, "data")

	interestEnd := ""
	validityDaysSurplus := 0
	for i := 0; i < anyJson.Size(); i++ {
		validityDaysSurplus += anyJson.Get('*', "validityDaysSurplus").Get(i).ToInt()

		currentInterestEnd := anyJson.Get('*', "interestEnd").Get(i).ToString()

		ct, _ := time.ParseInLocation(time.DateOnly, currentInterestEnd, time.Local)
		it, _ := time.ParseInLocation(time.DateOnly, interestEnd, time.Local)

		if ct.After(it) {
			interestEnd = currentInterestEnd
		}
	}

	return map[string]string{
		"mid":          mid,
		"interest_end": interestEnd,
		"surplus":      strconv.FormatInt(int64(validityDaysSurplus), 10),
	}
}
