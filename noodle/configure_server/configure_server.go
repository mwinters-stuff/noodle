package configure_server

import (
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/noodle/options"
	ldap_shim "github.com/mwinters-stuff/noodle/package-shims/ldap"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/rs/zerolog"
)

var (
	NewConfigureServer = NewConfigureServerImpl
)

type ConfigureServer interface {
	ConfigureAPI(api *operations.NoodleAPI) (database.Database, ldap_handler.LdapHandler, heimdall.Heimdall, error)
	ConfigureFlags(api *operations.NoodleAPI)
	SetupDatabase(config options.PostgresOptions, drop bool) (database.Database, bool, error)
	SetupLDAP(config options.LDAPOptions) (ldap_handler.LdapHandler, error)
}

type ConfigureServerImpl struct {
}

// ConfgureAPI implements ConfigureServer
func (i *ConfigureServerImpl) ConfigureAPI(api *operations.NoodleAPI) (database.Database, ldap_handler.LdapHandler, heimdall.Heimdall, error) {

	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	postgresOptions := api.CommandLineOptionsGroups[1].Options.(*options.PostgresOptions)
	lDAPOptions := api.CommandLineOptionsGroups[2].Options.(*options.LDAPOptions)

	api.ServeError = errors.ServeError

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if noodleOptions.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if noodleOptions.Config != "" {
		yfile, err := os.ReadFile(noodleOptions.Config)
		if err != nil {
			Logger.Error().Msg(err.Error())
			return nil, nil, nil, err
		}

		options, err := options.UnmarshalOptions(yfile)
		if err != nil {
			Logger.Error().Msg(err.Error())
			return nil, nil, nil, err
		}
		*postgresOptions = options.PostgresOptions
		*lDAPOptions = options.LDAPOptions
	}

	db, created, err := i.SetupDatabase(*postgresOptions, noodleOptions.Drop)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return nil, nil, nil, err
	}

	var ldap ldap_handler.LdapHandler
	if ldap, err = i.SetupLDAP(*lDAPOptions); err != nil {
		Logger.Error().Msg(err.Error())
		return nil, nil, nil, err
	}

	heimdall := heimdall.NewHeimdall(db)

	if !created {
		if err := api_handlers.LDAPRefresh(db, ldap); err != nil {
			return nil, nil, nil, err
		}
		if err := heimdall.UpdateFromServer(); err != nil {
			return nil, nil, nil, err
		}
	}

	return db, ldap, heimdall, nil
}

func (i *ConfigureServerImpl) ConfigureFlags(api *operations.NoodleAPI) {
	opts := &options.AllNoodleOptions{}

	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "config",
		LongDescription:  "Noodle Config",
		Options:          &opts.NoodleOptions,
	})
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "PostgreSQL",
		LongDescription:  "PostgreSQL Options",
		Options:          &opts.PostgresOptions,
	})
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "LDAP",
		LongDescription:  "LDAP Options",
		Options:          &opts.LDAPOptions,
	})
}

func (i *ConfigureServerImpl) SetupDatabase(config options.PostgresOptions, drop bool) (database.Database, bool, error) {
	db := database.NewDatabase(config)

	if err := db.Connect(); err != nil {
		return nil, false, err
	}

	db.Tables().InitTables(db)

	if drop {
		if err := db.Drop(); err != nil {
			return nil, false, err
		}
	}

	created, _ := db.CheckCreated()
	if !created {

		if err := db.Create(); err != nil {
			return nil, false, err
		}
	} else {
		needUpgrade, err := db.CheckUpgrade()
		if err != nil {
			return nil, false, err
		}

		if needUpgrade {
			if err = db.Upgrade(); err != nil {
				return nil, false, err
			}
		}
	}

	return db, created, nil
}

func (i *ConfigureServerImpl) SetupLDAP(config options.LDAPOptions) (ldap_handler.LdapHandler, error) {
	ldap := ldap_handler.NewLdapHandler(ldap_shim.NewLdapShim(), config)
	return ldap, ldap.Connect()
}

func NewConfigureServerImpl() ConfigureServer {
	return &ConfigureServerImpl{}
}
