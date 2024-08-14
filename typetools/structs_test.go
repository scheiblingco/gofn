package typetools_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/scheiblingco/gofn/typetools"
)

var NumberTestMap = map[string]interface{}{
	"int":           123,
	"customInt":     CustomInt,
	"int32":         int32(123),
	"customInt32":   CustomInt32,
	"int64":         int64(123),
	"customInt64":   CustomInt64,
	"uint":          uint(123),
	"customUint":    CustomUint,
	"uint32":        uint32(123),
	"customUint32":  CustomUint32,
	"uint64":        uint64(123),
	"customUint64":  CustomUint64,
	"float32":       float32(123.45),
	"customFloat32": CustomFloat32,
	"float64":       float64(123.45),
	"customFloat64": CustomFloat64,
	"uint8":         uint8(123),
	"customUint8":   CustomUint8,
	"uint16":        uint16(123),
	"customUint16":  CustomUint16,
	"uintptr":       uintptr(123),
	"customUintptr": CustomUintptr,
	"byte":          byte(123),
	"customByte":    CustomByte,
}

var StringTestMap = map[string]interface{}{
	"string":          "abc",
	"customString":    CustomString,
	"byteslice":       []byte("abc"),
	"byte":            byte('a'),
	"customByteslice": CustomByteslice,
	"rune":            'a',
	"runeSlice":       []rune("abc"),
	"customRuneSlice": CustomRuneslice,
}

func TestIsNumeric(t *testing.T) {
	for key, test := range NumberTestMap {
		if !typetools.IsNumericType(test) {
			t.Errorf("Expected %s to be a number", key)
		}
	}

	for key, test := range StringTestMap {
		if typetools.IsNumericType(test) && !strings.Contains(key, "rune") && !strings.Contains(key, "Rune") {
			t.Errorf("Expected %s to not be a number", key)
		}
	}
}

func TestIsString(t *testing.T) {
	for key, test := range StringTestMap {
		fmt.Println("Checking ", key)
		if !typetools.IsStringlikeType(test) {
			t.Errorf("Expected %s to be a string", key)
		}
	}
}

func TestIsReader(t *testing.T) {
	rd1 := strings.NewReader("abc")
	if !typetools.IsReaderType(rd1) {
		t.Errorf("Expected strings.NewReader to be a reader")
	}

	rd2 := bytes.NewBufferString("abc")
	if !typetools.IsReaderType(rd2) {
		t.Errorf("Expected bytes.NewBufferString to be a reader")
	}

	rd3 := bytes.NewBuffer([]byte("abc"))
	if !typetools.IsReaderType(rd3) {
		t.Errorf("Expected bytes.NewBuffer to be a reader")
	}

	if !typetools.IsReaderType(os.Stdin) {
		t.Errorf("Expected os.Stdin to be a reader")
	}

	if !typetools.IsReaderType(os.Stdout) {
		t.Errorf("Expected os.Stdout to be a reader")
	}

	if typetools.IsReaderType("HelloWorld") {
		t.Errorf("Didn't expect string to be a reader")
	}
}

func TestEnsureReader(t *testing.T) {
	rd := bytes.NewBuffer([]byte("abc"))
	rdr := typetools.EnsureReader(rd)

	content, err := io.ReadAll(rdr)
	if err != nil {
		t.Error(err)
	}

	if string(content) != "abc" {
		t.Errorf("Expected 'abc' but got %s", content)
	}
}

func TestEnsureInt(t *testing.T) {
	x := typetools.EnsureInt("123")
	fmt.Println(x)
}

// func IsNumber(v interface{}) bool {
// 	kindOf := reflect.TypeOf(v).Kind().String()

// 	switch kindOf {
// 	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64", "uint8", "uint16", "uintptr":
// 		return true
// 	}

// 	return false
// }

// func IsStringlike(v interface{}) bool {
// 	switch v.(type) {
// 	case string:
// 	case []byte:
// 	case []rune:
// 		return true
// 	}

// 	return false
// }

// func EnsureInt(v interface{}) int {
// 	switch v := v.(type) {
// 	case int:
// 		return v
// 	case int64:
// 	case uint32:
// 	case uint64:
// 	case float32:
// 	case float64:
// 		return int(v)
// 	}

// 	return 0
// }

// func EnsureUint(v interface{}) uint {
// 	switch v := v.(type) {
// 	case int:
// 		return uint(v)
// 	case int64:
// 		return uint(v)
// 	case uint32:
// 		return uint(v)
// 	case uint64:
// 		return uint(v)
// 	case float32:
// 		return uint(v)
// 	case float64:
// 		return uint(v)
// 	}

// 	return 0
// }

// func EnsureUint32(v interface{}) uint32 {
// 	switch v := v.(type) {
// 	case int:
// 		return uint32(v)
// 	case int64:
// 		return uint32(v)
// 	case uint32:
// 		return v
// 	case uint64:
// 		return uint32(v)
// 	case float32:
// 		return uint32(v)
// 	case float64:
// 		return uint32(v)
// 	}

// 	return 0
// }

// func EnsureUint64(v interface{}) uint64 {
// 	switch v := v.(type) {
// 	case int:
// 		return uint64(v)
// 	case int64:
// 		return uint64(v)
// 	case uint32:
// 		return uint64(v)
// 	case uint64:
// 		return v
// 	case float32:
// 		return uint64(v)
// 	case float64:
// 		return uint64(v)
// 	}

// 	return 0
// }

// func EnsureInt64(v interface{}) int64 {
// 	switch v := v.(type) {
// 	case int:
// 		return int64(v)
// 	case int64:
// 		return v
// 	case uint32:
// 	case uint64:
// 	case float32:
// 	case float64:
// 		return int64(v)
// 	}

// 	return 0
// }

// func EnsureString(v interface{}) string {
// 	switch v := v.(type) {
// 	case string:
// 		return v
// 	default:
// 		return ""
// 	}
// }

// func EnsureBool(v interface{}) bool {
// 	switch v := v.(type) {
// 	case bool:
// 		return v
// 	case byte:
// 	case int:
// 	case int64:
// 	case uint32:
// 	case uint64:
// 	case float32:
// 	case float64:
// 		return v != 0
// 	}
// 	return false
// }

// func EnsureFloat64(v interface{}) float64 {
// 	switch v := v.(type) {
// 	case float64:
// 		return v
// 	case float32:
// 	case int:
// 	case int64:
// 	case uint32:
// 	case uint64:
// 		return float64(v)
// 	}

// 	return 0
// }

// func SetField(obj interface{}, field string, value interface{}) error {
// 	ref := reflect.ValueOf(obj).Elem()
// 	if ref.Kind() != reflect.Struct {
// 		return fmt.Errorf("obj is not a struct")
// 	}

// 	strField := ref.FieldByName(field)
// 	if !strField.IsValid() {
// 		return fmt.Errorf("field does not exist")
// 	}

// 	if !strField.CanSet() {
// 		return fmt.Errorf("field cannot be set")
// 	}

// 	switch strField.Kind() {
// 	case reflect.String:
// 		strField.SetString(EnsureString(value))
// 	case reflect.Bool:
// 		strField.SetBool(EnsureBool(value))
// 	case reflect.Float32:
// 	case reflect.Float64:
// 		strField.SetFloat(EnsureFloat64(value))
// 	case reflect.Int:
// 	case reflect.Int64:
// 		strField.SetInt(EnsureInt64(value))
// 	case reflect.Uint32:
// 	case reflect.Uint64:
// 		strField.SetUint(EnsureUint64(value))
// 	default:
// 		strField.Set(reflect.ValueOf(value))
// 	}

// 	return nil
// }

// func MapToStruct(in map[string]interface{}, out interface{}) error {
// 	for k, v := range in {
// 		if err := SetField(out, k, v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
