/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: file.go
 * Desc: 日志 writer
 */

package writer

import (
	"fmt"
	"github.com/go-mumu/cs-go/library/common/flags"
	"os"
	"path"
)

func FileWriter() *os.File {
	dir, _ := path.Split(flags.LogPath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil
		}
	}

	file, err := os.OpenFile(flags.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}
