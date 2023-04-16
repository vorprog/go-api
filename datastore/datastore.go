package datastore

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/samber/lo"
)

var database *sql.DB
var datasets = map[string]map[string]interface{}{}
var datasetTypes = map[string]reflect.Type{}
var datasetQueries = map[string]string{}
var dataTypes = map[string]string{
	"string":  "TEXT",
	"int":     "INTEGER",
	"int64":   "INTEGER",
	"float64": "REAL",
	"bool":    "INTEGER",
}

type DatastoreInit struct {
	DatasetTypes map[string]reflect.Type
	SqlUrl       string
	Context      context.Context
}

func setPersistentStoreQueryString(datasetName string) {
	datasetType := datasetTypes[datasetName]

	var columns []string
	for i := 0; i < datasetType.NumField(); i++ {
		columns = append(columns, datasetType.Field(i).Name)
	}

	columnsString := strings.Join(columns, ",")
	datasetQueries[datasetName] = "INSERT INTO " + datasetName + " (id," + columnsString + ") VALUES (?,?" + strings.Repeat(",?", len(columns)-1) + ")"
}

func Init(init DatastoreInit) error {
	database, err := sql.Open("sqlite3", init.SqlUrl)
	if err != nil {
		return err
	}

	defer database.Close()

	datasetTypes = init.DatasetTypes
	for datasetName := range datasetTypes {

		dataType := datasetTypes[datasetName]
		var columns []string
		for i := 0; i < dataType.NumField(); i++ {
			if dataType.Field(i).PkgPath != "" {
				columns = append(columns, dataType.Field(i).Name+" "+dataTypes[dataType.Field(i).Type.Name()])
			}
		}

		sqlStatement := "CREATE TABLE IF NOT EXISTS " + datasetName + " (id TEXT PRIMARY KEY, " + strings.Join(columns, ",") + ")"
		_, err = database.Exec(sqlStatement)
		if err != nil {
			return err
		}

		setPersistentStoreQueryString(datasetName)
	}

	return nil
}

func Store(datasetName string, datasetItemKey string, datasetItemValue interface{}) (sql.Result, error) {
	values := []interface{}{datasetItemKey}
	for i := 0; i < datasetTypes[datasetName].NumField(); i++ {
		if datasetTypes[datasetName].Field(i).PkgPath != "" {
			values = append(values, reflect.ValueOf(datasetItemValue).Field(i).Interface())
		}
	}

	return database.Exec(datasetQueries[datasetName], values...)
}

func Set(datasetName string, datasetItemKey string, datasetItemValue interface{}) error {
	if _, err := Store(datasetName, datasetItemKey, datasetItemValue); err != nil {
		return err
	}

	datasets[datasetName][datasetItemKey] = datasetItemValue
	return nil
}

func Delete(datasetName string, datasetItemKey string) error {
	if _, err := peristentDelete(datasetName, datasetItemKey); err != nil {
		return err
	}

	delete(datasets[datasetName], datasetItemKey)
	return nil
}

func peristentDelete(datasetName string, datasetItemKey string) (sql.Result, error) {
	return database.Exec("DELETE FROM "+datasetName+" WHERE id = ?", datasetItemKey)
}

func Get(datasetName string, iteratee func(key string, value interface{}) (string, interface{})) map[string]interface{} {
	return lo.MapEntries(datasets[datasetName], iteratee)
}

func GetFromStore[T interface{}](datasetName string, key string) (T, error) {
	var queryResults, err = Query[T](datasetName, "SELECT * FROM "+datasetName+" WHERE id = "+key)
	if err != nil {
		return nil, err
	}

	return queryResults[0], nil
}

func Query[T interface{}](datasetName string, query string) ([]T, error) {
	sqlResult, err := database.Query(query)
	if err != nil {
		return nil, err
	}

	var results []T
	for sqlResult.Next() {
		var result T
		err = sqlResult.Scan(result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func GetKeys(datasetName string) []string {
	var datasetItemKeys []string
	for datasetItemKey := range datasets[datasetName] {
		datasetItemKeys = append(datasetItemKeys, datasetItemKey)
	}
	return datasetItemKeys
}
