//go:build exclude

package driver

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	DriverName    = "mysql"
	DataSourceFmt = "%s:%s@tcp(%s:%s)/%s"
)
