/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: mysql.go
 * Desc: mysql connect
 */

package mysql

import (
	"fmt"
	"github.com/go-mumu/cs-go/config"
	"github.com/go-mumu/cs-go/log"
	"github.com/go-mumu/cs-go/log/writer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	sysLog "log"
	"time"
)

type DefMysql struct {
	*gorm.DB
}

func InitDef(config *config.Config) *DefMysql {
	// 获取dsn
	dsn := formatDsn(config.DefMysql)

	db := connection(dsn, &gorm.Config{
		Logger: newLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "mh_dsh_",
		},
	})

	// 连接池
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// connect success
	log.Cli.Info("mysql connect success.")

	return &DefMysql{db}
}

// connection
func connection(dsn string, gormConf *gorm.Config) *gorm.DB {
	// connection mysql
	db, err := gorm.Open(mysql.Open(dsn), gormConf)

	if err != nil {
		log.Log.Error("mysql connection fail", "error", err)
		return nil
	}

	return db
}

// 格式化 Dsn
func formatDsn(mysqlConf *config.MysqlConf) string {
	return fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Protocol,
		mysqlConf.Address,
		mysqlConf.Port,
		mysqlConf.Dbname,
		mysqlConf.Charset,
		mysqlConf.ParseTime,
		mysqlConf.Loc,
	)
}

func newLogger() logger.Interface {
	return logger.New(
		sysLog.New(writer.FileWriter(), "\r\n", sysLog.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Silent,
		},
	)
}
