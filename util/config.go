package util

import (
	"os"
	"strings"

	"github.com/samber/lo"
)

type AppConfig struct {
	ServerPort      string
	SQLiteUrls      []string
	CacheSqLLiteUrl string
}

var Config = AppConfig{}

func InitConfig() {
	sqliteUrlString, _ := lo.Coalesce(os.Getenv("SQLITE_URLS"), "file:///app/app.sqlite3")
	Config.SQLiteUrls = strings.Split(sqliteUrlString, ",")
	Config.CacheSqLLiteUrl, _ = lo.Coalesce(os.Getenv("CACHE_SQLITE_URL"), "file:///app/cache.sqlite3")
	Config.ServerPort, _ = lo.Coalesce(os.Getenv("SERVER_PORT"), "8080")

}
