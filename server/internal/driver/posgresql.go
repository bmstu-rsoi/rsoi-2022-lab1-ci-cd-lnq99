package driver

import (
	_ "github.com/lib/pq"
)

const (
	DriverName    = "postgres"
	DataSourceFmt = "host=%s port=%s user=%s password=%s dbname='%s' sslmode=require"
	// TODO: add env sslmode=disable for local development
)
