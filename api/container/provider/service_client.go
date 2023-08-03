/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: client.go
 * Desc: service client provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/api/client"
	"github.com/google/wire"
)

var ServiceClientProviderSet = wire.NewSet(
	client.NewServiceClient,
)
