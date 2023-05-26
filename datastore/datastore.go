package datastore

import (
	"database/sql"
	"reflect"
	"strings"
	"sync"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
)

var database *sql.DB
var datasets = map[string]map[[16]byte]interface{}{}
var datasetTypes = map[string]reflect.Type{}
var datasetQueries = map[string]string{}
var sqliteTypes = map[string]string{
	"string":   "TEXT",
	"int":      "INTEGER",
	"bool":     "INTEGER",
	"[16]byte": "BLOB",
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

func setTables(datasetName string, dataType reflect.Type) error {
	var columnsSql []string
	for i := 0; i < dataType.NumField(); i++ {
		if dataType.Field(i).PkgPath != "" {
			columnsSql = append(columnsSql, dataType.Field(i).Name+" "+sqliteTypes[dataType.Field(i).Type.Name()])
		}
	}

	sqlStatement := "CREATE TABLE IF NOT EXISTS " + datasetName + " (id TEXT PRIMARY KEY, " + strings.Join(columnsSql, ",") + ")"
	_, err := database.Exec(sqlStatement)

	// TODO: check for migrations and handle schema changes

	return err
}

func loadDatasets() error {
	tables, err := database.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	sqlErrors := make(chan error)

	for tables.Next() {
		var tableName string
		tables.Scan(&tableName)
		wg.Add(1)

		go func() {
			contents, err := database.Query("SELECT * FROM " + tableName)
			if err != nil {
				sqlErrors <- err
			}

			for contents.Next() {
				var id [16]byte
				var datasetItemValue interface{}
				contents.Scan(&id, &datasetItemValue)
				datasets[tableName][id] = datasetItemValue
			}
			wg.Done()
		}()
	}

	select {
	case err := <-sqlErrors:
		return err
	default:
	}

	wg.Wait()
	close(sqlErrors)
	return nil
}

func Init(initTypes map[string]reflect.Type) error {
	var err error
	database, err = sql.Open("sqlite3", util.Config.SQLiteUrl)
	if err != nil {
		return err
	}

	for datasetName, typeValue := range initTypes {
		err = setTables(datasetName, typeValue)
		if err != nil {
			return err
		}

		setPersistentStoreQueryString(datasetName)
	}

	return loadDatasets()
}

func Store(datasetName string, id [16]byte, datasetItemValue interface{}) (sql.Result, error) {
	values := []interface{}{id}
	for i := 0; i < datasetTypes[datasetName].NumField(); i++ {
		if datasetTypes[datasetName].Field(i).PkgPath != "" {
			values = append(values, reflect.ValueOf(datasetItemValue).Field(i).Interface())
		}
	}

	return database.Exec(datasetQueries[datasetName], values...)
}

func Set(datasetName string, id [16]byte, item interface{}) error {
	if _, err := Store(datasetName, id, item); err != nil {
		return err
	}

	datasets[datasetName][id] = item
	return nil
}

func Add(datasetName string, item interface{}) ([16]byte, error) {
	var id [16]byte
	randomString := lo.RandomString(16, []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'a', 'b', 'c', 'd', 'e', 'f'})

	if _, err := Store(datasetName, id, item); err != nil {
		return id, err
	}

	datasets[datasetName][id] = item
	return id, nil
}

func Delete(datasetName string, datasetItemKey [16]byte) error {
	if _, err := peristentDelete(datasetName, datasetItemKey); err != nil {
		return err
	}

	delete(datasets[datasetName], datasetItemKey)
	return nil
}

func peristentDelete(datasetName string, datasetItemKey [16]byte) (sql.Result, error) {
	return database.Exec("DELETE FROM "+datasetName+" WHERE id = ?", datasetItemKey)
}

func Get(datasetName string, iteratee func(key [16]byte, value interface{}) (string, interface{})) map[string]interface{} {
	return lo.MapEntries(datasets[datasetName], iteratee)
}

func GetFromStore[T interface{}](datasetName string, key [16]byte) (T, error) {
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

func GetKeys(datasetName string) [][16]byte {
	return lo.Keys(datasets[datasetName])
}
