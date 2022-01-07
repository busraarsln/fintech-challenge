package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/busraarsln/fintech-challenge/config"
	"github.com/busraarsln/fintech-challenge/utils"
)

var db *sql.DB

func CreateDbConnection() (*sql.DB, error) {
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	c, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Mysql.MysqlUser,
		c.Mysql.MysqlPassword,
		c.Mysql.MysqlHost,
		c.Mysql.MysqlPort,
		c.Mysql.MysqlDbname,
	)

	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal("Unable to connect to db:", err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
