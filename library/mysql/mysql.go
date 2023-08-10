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
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/library/log/writer"
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

func InitDef() *DefMysql {
	// 获取dsn
	dsn := formatDsn()

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

	schema.RegisterSerializer("json", JSONSerializer{})

	log.Cli.Info("init mysql success.")

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
func formatDsn() string {
	return fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		config.V.GetString("mysql_def.username"),
		config.V.GetString("mysql_def.password"),
		config.V.GetString("mysql_def.protocol"),
		config.V.GetString("mysql_def.address"),
		config.V.GetInt("mysql_def.port"),
		config.V.GetString("mysql_def.dbname"),
		config.V.GetString("mysql_def.charset"),
		config.V.GetBool("mysql_def.parseTime"),
		config.V.GetString("mysql_def.loc"),
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
