package structtools

import (
	"fmt"
	"reflect"
)

func EnsureInt(v interface{}) int {
	switch v := v.(type) {
	case int:
		return v
	case int64:
	case uint32:
	case uint64:
	case float32:
	case float64:
		return int(v)
	}

	return 0
}

func EnsureUint(v interface{}) uint {
	switch v := v.(type) {
	case int:
		return uint(v)
	case int64:
		return uint(v)
	case uint32:
		return uint(v)
	case uint64:
		return uint(v)
	case float32:
		return uint(v)
	case float64:
		return uint(v)
	}

	return 0
}

func EnsureUint32(v interface{}) uint32 {
	switch v := v.(type) {
	case int:
		return uint32(v)
	case int64:
		return uint32(v)
	case uint32:
		return v
	case uint64:
		return uint32(v)
	case float32:
		return uint32(v)
	case float64:
		return uint32(v)
	}

	return 0
}

func EnsureUint64(v interface{}) uint64 {
	switch v := v.(type) {
	case int:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	}

	return 0
}

func EnsureInt64(v interface{}) int64 {
	switch v := v.(type) {
	case int:
		return int64(v)
	case int64:
		return v
	case uint32:
	case uint64:
	case float32:
	case float64:
		return int64(v)
	}

	return 0
}

func EnsureString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func EnsureBool(v interface{}) bool {
	switch v := v.(type) {
	case bool:
		return v
	case byte:
	case int:
	case int64:
	case uint32:
	case uint64:
	case float32:
	case float64:
		return v != 0
	}
	return false
}

func EnsureFloat64(v interface{}) float64 {
	switch v := v.(type) {
	case float64:
		return v
	case float32:
	case int:
	case int64:
	case uint32:
	case uint64:
		return float64(v)
	}

	return 0
}

func SetField(obj interface{}, field string, value interface{}) error {
	ref := reflect.ValueOf(obj).Elem()
	if ref.Kind() != reflect.Struct {
		return fmt.Errorf("obj is not a struct")
	}

	strField := ref.FieldByName(field)
	if !strField.IsValid() {
		return fmt.Errorf("field does not exist")
	}

	if !strField.CanSet() {
		return fmt.Errorf("field cannot be set")
	}

	switch strField.Kind() {
	case reflect.String:
		strField.SetString(EnsureString(value))
	case reflect.Bool:
		strField.SetBool(EnsureBool(value))
	case reflect.Float32:
	case reflect.Float64:
		strField.SetFloat(EnsureFloat64(value))
	case reflect.Int:
	case reflect.Int64:
		strField.SetInt(EnsureInt64(value))
	case reflect.Uint32:
	case reflect.Uint64:
		strField.SetUint(EnsureUint64(value))
	default:
		strField.Set(reflect.ValueOf(value))
	}

	return nil
}

func MapToStruct(in map[string]interface{}, out interface{}) error {
	for k, v := range in {
		if err := SetField(out, k, v); err != nil {
			return err
		}
	}

	return nil
}
