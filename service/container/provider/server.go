/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-10
 * File: server.go
 * Desc: server provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/service/server"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(
	server.NewServer,
)
