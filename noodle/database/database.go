package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
)

// You only need **one** of these per package!
//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name Database

const DATABASE_VERSION int = 1

var (
	NewDatabase = NewDatabaseImpl
)

//counterfeiter:generate . Database
type Database interface {
	Connect() error
	CheckUpgrade() (bool, error)
	Create() error
	Drop() error
	GetVersion() (int, error)
	Upgrade(current_version int) error
	Close()

	Pool() *pgxpool.Pool
}

type DatabaseImpl struct {
	appConfig yamltypes.AppConfig
	pool      *pgxpool.Pool
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

func NewDatabaseImpl(appConfig yamltypes.AppConfig) Database {
	return &DatabaseImpl{
		appConfig: appConfig,
	}
}

// CheckUpgrade implements Database
func (i *DatabaseImpl) Upgrade(current_version int) error {
	Logger.Info().Msgf("upgrade database from %d to %d", current_version, DATABASE_VERSION)
	return nil
}

// Connect implements Database
func (i *DatabaseImpl) Connect() error {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		i.appConfig.Postgres.User,
		i.appConfig.Postgres.Password,
		i.appConfig.Postgres.Hostname,
		i.appConfig.Postgres.Port,
		i.appConfig.Postgres.Db)

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

	return err
}

// Drop implements Database
func (i *DatabaseImpl) Drop() error {
	Logger.Info().Msg("dropping database")
	_, err := i.pool.Exec(context.Background(),
		`
DROP TABLE version;
DROP TABLE application_template;
DROP TABLE users;
`)

	return err

}
