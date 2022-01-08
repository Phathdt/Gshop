package sdk

import (
	"database/sql"
)

type ServiceConfig struct {
	*sql.DB
}
