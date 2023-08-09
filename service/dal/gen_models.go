/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: gen_models.go
 * Desc: generate model
 */

package dal

import (
	"gorm.io/gorm"
)
import "gorm.io/gen"

// GenDefModels gen default models
func GenDefModels(db *gorm.DB) {
	cfg := gen.Config{
		OutPath:      "./service/dal/query",
		OutFile:      "defGen.go",
		ModelPkgPath: "./service/dal/model",

		FieldNullable:    true,
		FieldCoverable:   true,
		FieldSignable:    true,
		FieldWithTypeTag: true,
	}

	g := gen.NewGenerator(cfg)

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// bool mapping
		"datetime": func(columnType gorm.ColumnType) (dataType string) {
			return "datetime.Datetime"
		},
	}

	g.WithDataTypeMap(dataMap)

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("wxuser"))

	g.Execute()
}
