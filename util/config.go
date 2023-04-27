package util

import (
	"os"

	"github.com/samber/lo"
)

type AppConfig struct {
	ServerPort string
	SQLiteUrl  string
}

var Config = AppConfig{}

func InitConfig() {
	Config.SQLiteUrl, _ = lo.Coalesce(os.Getenv("APP_SQLITE_URL"), "file:///app/app.sqlite3")
	Config.ServerPort, _ = lo.Coalesce(os.Getenv("APP_SERVER_PORT"), "8080")
}
