package driver

import (
	"database/sql"
	"fmt"
	"rsoi-1/config"
)

func NewSQLDatabase(cfg *config.DbConfig) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf(DataSourceFmt, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	pool, err := sql.Open(DriverName, dataSourceName)

	return pool, err
}
