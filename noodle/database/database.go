package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mwinters-stuff/noodle/noodle/options"
)

// You only need **one** of these per package!
//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name Database

const DATABASE_VERSION int = 1

var (
	NewDatabase = NewDatabaseImpl
)

type Database interface {
	Connect() error
	CheckCreated() (bool, error)
	CheckUpgrade() (bool, error)
	Create() error
	Drop() error
	GetVersion() (int, error)
	Upgrade() error
	Close()

	Tables() Tables

	Pool() *pgxpool.Pool
}

type DatabaseImpl struct {
	pgConfig options.PostgresOptions
	pool     *pgxpool.Pool
	tables   Tables
}

// Tables implements Database
func (i *DatabaseImpl) Tables() Tables {
	return i.tables
}

// CheckCreated implements Database
func (i *DatabaseImpl) CheckCreated() (bool, error) {
	_, err := i.GetVersion()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Pool implements Database
func (i *DatabaseImpl) Pool() *pgxpool.Pool {
	return i.pool
}

// Close implements Database
func (i *DatabaseImpl) Close() {
	i.pool.Close()
}

// CheckUpgrade implements Database
func (i *DatabaseImpl) CheckUpgrade() (bool, error) {
	current_version, err := i.GetVersion()
	if err == nil {
		if current_version < DATABASE_VERSION {
			Logger.Info().Msgf("upgrade database required from %d to %d", current_version, DATABASE_VERSION)
			return true, nil
		} else if current_version > DATABASE_VERSION {
			Logger.Error().Msg("cannot downgrade database")
			err = fmt.Errorf("datatabase downgrade not allowed")
		} else {
			Logger.Info().Msg("no database upgrade required")
		}
	}
	return false, err
}

// GetVersion implements Database
func (i *DatabaseImpl) GetVersion() (int, error) {
	var version int

	err := i.pool.QueryRow(context.Background(), "SELECT version FROM version").Scan(&version)
	Logger.Info().Msgf("current database version %d", version)
	return version, err
}

// CheckUpgrade implements Database
func (i *DatabaseImpl) Upgrade() error {
	current_version, _ := i.GetVersion()
	Logger.Info().Msgf("upgrade database from %d to %d", current_version, DATABASE_VERSION)
	return nil
}

// Connect implements Database
func (i *DatabaseImpl) Connect() error {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		i.pgConfig.User,
		i.pgConfig.Password,
		i.pgConfig.Hostname,
		i.pgConfig.Port,
		i.pgConfig.Database)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err == nil {
		pool.Reset()
		i.pool = pool
		var num int
		err = i.pool.QueryRow(context.Background(), "SELECT 1").Scan(&num)
		if err == nil {
			Logger.Info().Msg("database connected")
		}
	}
	if err != nil {
		Logger.Error().Msgf("database connection failed %s", err)
	}
	return err
}

// Create implements Database
func (i *DatabaseImpl) Create() error {
	Logger.Info().Msg("creating database")

	_, err := i.pool.Exec(context.Background(),
		fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS version (version int);
DELETE FROM version;
INSERT INTO version (version) values (%d)
`, DATABASE_VERSION))
	if err != nil {
		return err
	}

	return i.tables.Create()
}

// Drop implements Database
func (i *DatabaseImpl) Drop() error {
	Logger.Info().Msg("dropping database")
	_, err := i.pool.Exec(context.Background(), `DROP TABLE version`)

	if err != nil {
		return err
	}

	return i.tables.Drop()

}

func NewDatabaseImpl(pgConfig options.PostgresOptions) Database {
	return &DatabaseImpl{
		pgConfig: pgConfig,
		tables:   NewTables(),
	}
}
