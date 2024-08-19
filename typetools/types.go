package typetools

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
)

const (
	MaxUint   = uint64(^uint(0))
	MaxUint8  = uint64(255)
	MaxUint16 = uint64(65535)
	MaxUint32 = uint64(4294967295)
	MaxUint64 = uint64(18446744073709551615)

	MaxInt = int(MaxUint >> 1)
	MinInt = -MaxInt - 1

	MaxInt8 = 127
	MinInt8 = -128

	MaxInt16 = 32767
	MinInt16 = -32768

	MaxInt32 = 2147483647
	MinInt32 = -2147483648

	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808

	MaxFloat32 = float64(3.40282346638528859811704183484516925440e+38)
	MinFloat32 = float64(-3.40282346638528859811704183484516925440e+38)
)

var validNumericTypes = []string{
	"int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64", "uint8", "uint16", "uintptr",
}

var validStringlikeTypes = []string{
	"string", "[]byte", "[]rune", "byte",
}

func IsNumericType(v interface{}) bool {
	kindOf := reflect.TypeOf(v).Kind().String()

	for _, kind := range validNumericTypes {
		if kind == kindOf {
			return true
		}
	}

	return false
}

func IsStringlikeType(v interface{}) bool {
	switch v.(type) {
	case []byte, []rune, byte, rune:
		return true
	}

	kindOf := reflect.TypeOf(v).Kind().String()

	if kindOf == "slice" {
		kindOfElem := reflect.TypeOf(v).Elem().Kind()
		kindOfElemStr := kindOfElem.String()
		if kindOfElemStr == "int32" || kindOfElemStr == "uint8" {
			return true
		}
	}

	for _, kind := range validStringlikeTypes {
		if kind == kindOf {
			return true
		}
	}

	return false
}

func IsReaderType(v interface{}) bool {
	impl := reflect.TypeOf(v).Implements(reflect.TypeOf((*io.Reader)(nil)).Elem())

	return impl
}

func EnsureReader(v interface{}) io.Reader {
	if !IsReaderType(v) {
		return nil
	}

	return v.(io.Reader)
}

func EnsureInt64(v interface{}) int64 {
	if IsStringlikeType(v) {
		sureString := EnsureString(v)

		parse, err := strconv.ParseInt(sureString, 10, 64)
		if err != nil {
			return 0
		}

		return parse
	}

	if !IsNumericType(v) {
		return 0
	}

	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int()
	case uint, uintptr, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint())
	case float32, float64:
		return int64(reflect.ValueOf(v).Float())
	}

	return 0
}

func EnsureInt32(v interface{}) int32 {
	i64val := EnsureInt64(v)

	if i64val > int64(MaxInt32) || i64val < int64(MinInt32) {
		return 0
	}

	return int32(i64val)
}

func EnsureInt16(v interface{}) int16 {
	i64val := EnsureInt64(v)

	if i64val > int64(MaxInt16) || i64val < int64(MinInt16) {
		return 0
	}

	return int16(i64val)
}

func EnsureInt8(v interface{}) int8 {
	i64val := EnsureInt64(v)

	if i64val > int64(MaxInt8) || i64val < int64(MinInt8) {
		return 0
	}

	return int8(i64val)
}

func EnsureInt(v interface{}) int {
	i64val := EnsureInt64(v)

	if i64val > int64(MaxInt) || i64val < int64(MinInt) {
		return 0
	}

	return int(i64val)
}

func EnsureUint64(v interface{}) uint64 {
	if IsStringlikeType(v) {
		sureString := EnsureString(v)

		parse, err := strconv.ParseUint(sureString, 10, 64)
		if err != nil {
			return 0
		}

		return parse
	}

	if !IsNumericType(v) {
		return 0
	}

	switch v := v.(type) {
	case int, int8, int32, int64:
		if reflect.ValueOf(v).Int() < 0 {
			return 0
		}
		return uint64(reflect.ValueOf(v).Int())
	case uint, uintptr, uint8, uint16, uint32, uint64:
		return uint64(reflect.ValueOf(v).Uint())
	case float32, float64:
		if reflect.ValueOf(v).Float() < 0 {
			return 0
		}
		return uint64(reflect.ValueOf(v).Float())
	}

	return 0
}

func EnsureUint32(v interface{}) uint32 {
	ui64val := EnsureUint64(v)

	if ui64val > MaxUint32 {
		return 0
	}

	return uint32(ui64val)
}

func EnsureUint16(v interface{}) uint16 {
	ui64val := EnsureUint64(v)

	if ui64val > MaxUint16 {
		return 0
	}

	return uint16(ui64val)
}

func EnsureUint8(v interface{}) uint8 {
	ui64val := EnsureUint64(v)

	if ui64val > MaxUint8 {
		return 0
	}

	return uint8(ui64val)
}

func EnsureUint(v interface{}) uint {
	ui64val := EnsureUint64(v)

	if ui64val > MaxUint {
		return 0
	}

	return uint(ui64val)
}

func EnsureString(v interface{}) string {
	if IsNumericType(v) {
		return fmt.Sprintf("%v", v)
	}

	switch v := v.(type) {
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	}

	return ""
}

func EnsureBool(v interface{}) bool {
	if IsStringlikeType(v) {
		sureString := EnsureString(v)

		parse, err := strconv.ParseBool(sureString)
		if err != nil {
			return false
		}

		return parse
	}

	switch v := v.(type) {
	case bool:
		return v
	case byte, int, int64, uint32, uint64, float32, float64:
		return v != 0
	}
	return false
}

func EnsureFloat64(v interface{}) float64 {
	if IsStringlikeType(v) {
		sureString := EnsureString(v)

		parse, err := strconv.ParseFloat(sureString, 64)
		if err != nil {
			return 0
		}

		return parse
	}

	if !IsNumericType(v) {
		return 0
	}

	switch v := v.(type) {
	case int, int8, int32, int64:
		return float64(reflect.ValueOf(v).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint())
	case float32, float64:
		return reflect.ValueOf(v).Float()
	}

	return 0
}

func EnsureFloat32(v interface{}) float32 {
	f64val := EnsureFloat64(v)

	if f64val > MaxFloat32 || f64val < MinFloat32 {
		return 0
	}

	return float32(f64val)
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
