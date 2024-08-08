package tsqltools

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/scheiblingco/gofn/structtools"
)

func QueryMssql(connectionstring, query string, params map[string]interface{}) ([]map[string]interface{}, error) {
	if connectionstring == "" {
		return nil, fmt.Errorf("connectionstring is empty")
	}

	if query == "" {
		return nil, fmt.Errorf("query is empty")
	}

	conn, err := sql.Open("sqlserver", connectionstring)
	if err != nil {
		return nil, err
	}

	var resultRows []map[string]interface{}
	var qParams []any

	for k, v := range params {
		qParams = append(qParams, sql.Named(k, v))
	}

	rows, err := conn.Query(query, qParams...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		resultRows = append(resultRows, m)
	}

	return resultRows, nil
}

func appendToSlice(arrPtr interface{}, itm interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	itmv := reflect.ValueOf(itm)
	itmvind := reflect.Indirect(itmv)

	value.Set(reflect.Append(value, itmvind))
}

func QueryMssqlStruct(connectionstring, query string, params map[string]interface{}, out interface{}) error {
	if connectionstring == "" {
		return fmt.Errorf("connectionstring is empty")
	}

	if query == "" {
		return fmt.Errorf("query is empty")
	}

	conn, err := sql.Open("sqlserver", connectionstring)
	if err != nil {
		return err
	}

	var qParams []any

	for k, v := range params {
		qParams = append(qParams, sql.Named(k, v))
	}

	rows, err := conn.Query(query, qParams...)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	refType := reflect.TypeOf(out).Elem().Elem()

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		single := reflect.New(refType).Interface()

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			if err := structtools.SetField(single, colName, *val); err != nil {
				return err
			}
		}

		appendToSlice(out, single)
	}

	return nil
}
