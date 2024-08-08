package main

import (
	"fmt"

	"github.com/scheiblingco/gofn/cfgtools"
	"github.com/scheiblingco/gofn/tsqltools"
)

type Config struct {
	SqlServer string `json:"sql_server"`
}

type TestStruct struct {
	ValA string  `json:"val_a"`
	ValB int     `json:"val_b"`
	ValC bool    `json:"val_c"`
	ValD float32 `json:"val_d"`
	ValE float64 `json:"val_e"`
}

var TestMap = map[string]interface{}{
	"ValA": "test",
	"ValB": 123,
	"ValC": true,
	"ValD": 1.23,
	"ValE": 1.23,
}

type TestDbStruct struct {
	CreatedName string `json:"CreatedName"`
	CreatedDate string `json:"CreatedDate"`
}

func main() {
	cfg := Config{}
	cfgtools.LoadJsonConfig("config.json", &cfg)

	sqlq := "SELECT TOP 100 [CreatedName], [CreatedDate] FROM [EXT-HALDOR]..[TEDToAzureGuids]"

	resv := []TestDbStruct{}

	x := tsqltools.QueryMssqlStruct(cfg.SqlServer, sqlq, map[string]interface{}{}, &resv)

	fmt.Println(x)

}
