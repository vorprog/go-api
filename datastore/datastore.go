package datastore

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
)

var persistentDb *sql.DB
var cacheDb *sql.DB

// https://www.sqlite.org/lang_UPSERT.html
const upsertSqlTemplate = `INSERT INTO %s (%s)
VALUES (%s)
ON CONFLICT(%s) DO UPDATE SET %s`

func Init() (result sql.Result, err error) {
	persistentDb, err = sql.Open("sqlite3", util.Config.SQLiteUrl)
	if err != nil {
		return
	}

	cacheDb, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return
	}

	return cacheDb.Exec(".restore cache_db" + util.Config.SQLiteUrl)
}

func Upsert(tableName string, data map[string]interface{}) (result sql.Result, err error) {
	columnNames := lo.Keys(data)
	impliedPrimaryKey := columnNames[0]
	columnsToUpdate := lo.Slice(columnNames, 1, len(columnNames))
	insertValuePlaceholders := lo.Map(columnNames, func(columnName string, index int) string { return "?" })
	updateSql := lo.Map(columnsToUpdate, func(columnName string, index int) string { return columnName + "=excluded." + columnName })

	sql := fmt.Sprintf(upsertSqlTemplate,
		tableName,
		strings.Join(columnNames, ","),
		strings.Join(insertValuePlaceholders, ","),
		impliedPrimaryKey,
		strings.Join(updateSql, ","))

	return Cache(sql, lo.Values(data)...)
}

func Store(sql string, values ...any) (sql.Result, error) {
	return persistentDb.Exec(sql, values...)
}

func Cache(sql string, values ...any) (result sql.Result, err error) {
	result, err = Store(sql, values...)
	if err != nil {
		return
	}
	go cacheDb.Exec(sql, values...)
	return
}

func Query(sql string, values ...any) (*sql.Rows, error) {
	return persistentDb.Query(sql, values...)
}

func Get[T interface{}](sql string, values ...any) ([]T, error) {
	sqlRows, err := cacheDb.Query(sql, values...)
	if err != nil {
		return nil, err
	}

	return MapResults[T](sqlRows)
}

func MapResults[T interface{}](rows *sql.Rows) (results []T, err error) {
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err = rows.Scan(valuePtrs...)
		if err != nil {
			return
		}

		mappedFields := make(map[string]interface{})
		for i, col := range columns {
			mappedFields[col] = values[i]
		}

		var nextResult T
		err = util.FillStruct(&nextResult, mappedFields)
		if err != nil {
			return
		}

		results = append(results, nextResult)
	}

	return
}
