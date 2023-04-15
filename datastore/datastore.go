package datastore

import (
	"database/sql"
	"encoding/gob"
	"os"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
)

var datasets = map[string]map[string]interface{}{}
var datasetTypes = map[string]reflect.Type{}
var datasetQueries = map[string]string{}
var database *sql.DB

func Init(initDatasetTypes map[string]reflect.Type) error {
	datasetTypes = initDatasetTypes
	for datasetName := range datasetTypes {
		setPersistentStoreQueryString(datasetName)
	}

	for _, datasetName := range strings.Split(os.Getenv("DATASETS"), ",") {
		gobFile, err := util.GetS3File(os.Getenv("S3_BUCKET"), datasetName+".gob", os.Getenv("AWS_REGION"))
		if err != nil {
			return err
		}

		var dataset map[string]interface{}
		err = gob.NewDecoder(gobFile).Decode(&dataset)
		if err != nil {
			return err
		}

		datasets[datasetName] = dataset
	}

	database, err := sql.Open("sqlite3", "file:/app/data.db?cache=shared&mode=rwc")
	if err != nil {
		return err
	}

	defer database.Close()

	return nil
}

func Get(datasetName string, iteratee func(key string, value interface{}) (string, interface{})) map[string]interface{} {
	return lo.MapEntries(datasets[datasetName], iteratee)
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

func Set(datasetName string, datasetItemKey string, datasetItemValue interface{}) {
	persistentStore(datasetName, datasetItemKey, datasetItemValue)
	datasets[datasetName][datasetItemKey] = datasetItemValue
}

func Delete(datasetName string, datasetItemKey string) {
	peristentDelete(datasetName, datasetItemKey)
	delete(datasets[datasetName], datasetItemKey)
}

func persistentStore(datasetName string, datasetItemKey string, datasetItemValue interface{}) (sql.Result, error) {
	if datasetQueries[datasetName] == "" {
		setPersistentStoreQueryString(datasetName)
	}

	var values []interface{}
	values = append(values, datasetItemKey)
	for i := 0; i < datasetTypes[datasetName].NumField(); i++ {
		if datasetTypes[datasetName].Field(i).PkgPath != "" {
			values = append(values, reflect.ValueOf(datasetItemValue).Field(i).Interface())
		}
	}

	return database.Exec(datasetQueries[datasetName], values...)
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

func peristentDelete(datasetName string, datasetItemKey string) (sql.Result, error) {
	return database.Exec("DELETE FROM "+datasetName+" WHERE id = ?", datasetItemKey)
}
