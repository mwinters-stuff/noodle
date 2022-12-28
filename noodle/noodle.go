package noodle

import (
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mwinters-stuff/noodle/noodle/database"
)

type Noodle interface {
	Run()
}

type NoodleImpl struct {
}

// Run implements Noodle
func (*NoodleImpl) Run() {
	Logger.Info().Msg("Noodle Start")

	pguser := os.Getenv("POSTGRES_USER")
	pgpassword := os.Getenv("POSTGRES_PASSWORD")
	pgdb := os.Getenv("POSTGRES_DB")
	pghostname := os.Getenv("POSTGRES_HOSTNAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pguser, pgpassword, pghostname, pgdb)

	database := database.NewDatabase(connStr)

	err := database.Connect()
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}
	// database.Drop()

	err = database.CheckUpgrade()
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) {
			if pgerrcode.IsSyntaxErrororAccessRuleViolation(e.Code) {
				err = database.Create()
			}
		}
		if err != nil {
			Logger.Fatal().Msg(err.Error())
		}
		Logger.Info().Msg("database created!")
	} else {
		Logger.Info().Msg("database does not need upgrading!")
	}

}

func NewNoodle() Noodle {
	noodle := &NoodleImpl{}
	return noodle
}
