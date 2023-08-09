/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-09
 * File: json_serializer.go
 * Desc:
 */

package mysql

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm/schema"
	"reflect"
)

// JSONSerializer json序列化器
type JSONSerializer struct {
}

// Scan 实现 Scan 方法
func (JSONSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}

		json := jsoniter.ConfigCompatibleWithStandardLibrary

		err = json.Unmarshal(bytes, fieldValue.Interface())
	}

	fmt.Println(field, fieldValue, dbValue)

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

// Value 实现 Value 方法
func (JSONSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	result, err := json.Marshal(fieldValue)
	if string(result) == "null" {
		if field.TagSettings["NOT NULL"] != "" {
			return "", nil
		}
		return nil, err
	}
	return string(result), err
}
