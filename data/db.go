package data

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/config"
)

// DBHandle holds logger
type DBHandle struct {
	l  hclog.Logger
	DB *pg.DB
}

// NewDBHandle creates DBHandle
func NewDBHandle(l hclog.Logger) *DBHandle {
	return &DBHandle{l: l, DB: nil}
}

// ErrUnableToReachDB is custom error message when unagle to reach PG Server
var ErrUnableToReachDB = fmt.Errorf("Unable to reach PostgreSQL Server. Is it running?")

// ErrMissingDBConnection is custom error message when DB connection is nil
var ErrMissingDBConnection = fmt.Errorf("Connection to the DB doesn't exist to Initialize. Create a connection first")

// Connect returns a new db connection or error
func (dbh *DBHandle) Connect(conf *config.Config) error {
	dbh.DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.DB.Hostname, conf.DB.Port),
		User:     conf.DB.Username,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})

	ctx := context.Background()
	if err := dbh.DB.Ping(ctx); err != nil {
		return ErrUnableToReachDB
	}
	return nil
}

// Init does a DB Migration if necessary
func (dbh *DBHandle) Init() error {
	if dbh.DB == nil {
		return ErrMissingDBConnection
	}
	err := dbh.createSchema()
	if err != nil {
		return err
	}
	return nil
}

// createSchema creates database schema for User and Story models.
func (dbh *DBHandle) createSchema() error {
	models := []interface{}{
		(*Item)(nil),
	}

	for _, model := range models {

		err := dbh.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
