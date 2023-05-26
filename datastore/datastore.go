package datastore

import (
	"database/sql"
	"strings"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
)

var persistentDb *sql.DB
var cacheDb *sql.DB

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

func Upsert(tableName string, data map[string]string) (result sql.Result, err error) {
	colmnNames := strings.Join(lo.Keys(data), ",")
	placeholders := lo.MapToSlice(data, func(key string, value string) string { return "?" })
	sql := "INSERT INTO " + tableName + " (" + colmnNames + ") VALUES (" + strings.Join(placeholders, ",") + ") ON CONFLICT(id) DO UPDATE SET " + util.Join(columns, " = ?, ") + " = ?"
	return Store(sql, lo.Values(data)...)
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
