package cmd

import (
	"database/sql"
	"fmt"
	"pool_backend/src/global"
	"strings"
	// "os"
	// "strings"
	// _ "github.com/lib/pq"
)

var (
	db          *sql.DB
	schemeName  string
	queryResult map[string]column
)

type column map[string]string

type dbconfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	DbName   string `json:"dbname"`
	Password string `json:"password"`
}

var template string = `
SELECT 
tb.tablename as tablename,
a.attname AS columnname,
t.typname AS type
FROM
pg_class as c,
pg_attribute as a, 
pg_type as t,
(select tablename from pg_tables where schemaname = "public" ) as tb
WHERE  a.attnum > 0 
and a.attrelid = c.oid
and a.atttypid = t.oid
and c.relname = tb.tablename 
order by tablename
`

//postgresql type -> go type
var pgMap = map[string]string{
	"int4":      "int32",
	"int8":      "int64",
	"float4":    "float32",
	"float8":    "float64",
	"double":    "float64",
	"varchar":   "string",
	"boolean":   "bool",
	"timestamp": "time.Time",
	"date":      "time.Time",
}

//QueryTables 查询数据库表
func QueryTables() {
	queryResult = make(map[string]column)
	rows, err := global.DB.GetDbR().Exec(template).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		tableName, colName, colType := "", "", ""
		rows.Scan(&tableName, &colName, &colType)
		_, ok := queryResult[tableName]
		if !ok {
			queryResult[tableName] = make(map[string]string)
		}
		queryResult[tableName][colName] = colType
	}
}

//Report 生成
func Report() {
	for tableName, ctMap := range queryResult {
		data := "\ntype %s struct {\n" + tableName
		for colName, coltype := range ctMap {
			//translate to go struct foramt
			retype := pgMap[coltype]
			if retype == "" {
				retype = coltype
			}
			jsonName := strings.ToLower(colName)
			vname := strings.ToUpper(jsonName[0:1]) + jsonName[1:]
			data = data + "\t%-10s %-10s `json:\"%s\"`\n" + vname + retype + jsonName

		}
		data = data + "}\n"
		fmt.Println("---生成struct数据----", data)
	}

}
