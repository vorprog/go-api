package datastore

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
)

type DatasetItem interface {
	GetId() string
}

var database *sql.DB
var datasets = map[string]map[string]DatasetItem{}
var datasetTypes = map[string]reflect.Type{}
var datasetQueries = map[string]string{}
var dataTypes = map[string]string{
	"string":  "TEXT",
	"int":     "INTEGER",
	"int64":   "INTEGER",
	"float64": "REAL",
	"bool":    "INTEGER",
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

func createTable(datasetName string, dataType reflect.Type) error {
	var columnsSql []string
	for i := 0; i < dataType.NumField(); i++ {
		if dataType.Field(i).PkgPath != "" {
			columnsSql = append(columnsSql, dataType.Field(i).Name+" "+dataTypes[dataType.Field(i).Type.Name()])
		}
	}

	sqlStatement := "CREATE TABLE IF NOT EXISTS " + datasetName + " (id TEXT PRIMARY KEY, " + strings.Join(columnsSql, ",") + ")"
	_, err := database.Exec(sqlStatement)
	return err
}

func loadDatasets() error {
	tables, err := database.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return err
	}

	for tables.Next() {
		var tableName string
		tables.Scan(&tableName)

		// TODO: wait group and error handling

		go func() {
			contents, _ := database.Query("SELECT * FROM " + tableName)
			for contents.Next() {
				var id string
				var datasetItemValue DatasetItem
				contents.Scan(&id, &datasetItemValue)
				datasets[tableName][id] = datasetItemValue
			}
		}()
	}

	return nil
}

func Init(dataTypes map[string]reflect.Type) error {
	var err error
	database, err = sql.Open("sqlite3", util.Config.SQLiteUrl)
	if err != nil {
		return err
	}

	for datasetName, dataType := range dataTypes {
		err = createTable(datasetName, dataType)
		if err != nil {
			return err
		}

		setPersistentStoreQueryString(datasetName)
	}

	return loadDatasets()
}

func Store(datasetName string, datasetItemValue DatasetItem) (sql.Result, error) {
	values := []interface{}{}
	for i := 0; i < datasetTypes[datasetName].NumField(); i++ {
		if datasetTypes[datasetName].Field(i).PkgPath != "" {
			values = append(values, reflect.ValueOf(datasetItemValue).Field(i).Interface())
		}
	}

	return database.Exec(datasetQueries[datasetName], values...)
}

func Set(datasetName string, item DatasetItem) error {
	if _, err := Store(datasetName, item); err != nil {
		return err
	}

	datasets[datasetName][item.GetId()] = item
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

func Get(datasetName string, iteratee func(key string, value DatasetItem) (string, interface{})) map[string]interface{} {
	return lo.MapEntries(datasets[datasetName], iteratee)
}

func GetFromStore[T interface{}](datasetName string, key string) (T, error) {
	var queryResults, err = Query[T](datasetName, "SELECT * FROM "+datasetName+" WHERE id = "+key)
	if err != nil {
		return queryResults[0], err
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
