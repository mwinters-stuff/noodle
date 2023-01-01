package noodle

import (
	"fmt"
	"os"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
)

const pgport = 15432

// const pgport = 5432

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

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", pguser, pgpassword, pghostname, pgport, pgdb)

	db := database.NewDatabase(connStr)

	err := db.Connect()
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}
	db.Drop()

	// err = database.CheckUpgrade()
	// if err != nil {
	// 	var e *pgconn.PgError
	// 	if errors.As(err, &e) {
	// 		if pgerrcode.IsSyntaxErrororAccessRuleViolation(e.Code) {
	// 			err = database.Create()
	// 		}
	// 	}
	err = db.Create()
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}
	// 	Logger.Info().Msg("database created!")
	// } else {
	// 	Logger.Info().Msg("database does not need upgrading!")
	// }

	app := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads & tracking. After you set it up, it'll cover ALL your home devices, and you don't need any client-side software for that.\r\n\r\nIt operates as a DNS server that re-routes tracking domains to a \"black hole,\" thus preventing your devices from connecting to those servers. It's based on software we use for our public AdGuard DNS servers -- both share a lot of common code.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Insert(app)
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	result, err := table.Search("AdGuard")
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	print(result)
}

func NewNoodle() Noodle {
	noodle := &NoodleImpl{}
	return noodle
}
