package configure_server

import (
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
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
	ConfigureAPI(api *operations.NoodleAPI) (database.Database, ldap_handler.LdapHandler, heimdall.Heimdall)
	ConfigureFlags(api *operations.NoodleAPI)
	SetupDatabase(config options.PostgresOptions, drop bool) (database.Database, error)
	SetupLDAP(config options.LDAPOptions) (ldap_handler.LdapHandler, error)
}

type ConfigureServerImpl struct {
}

// ConfgureAPI implements ConfigureServer
func (i *ConfigureServerImpl) ConfigureAPI(api *operations.NoodleAPI) (database.Database, ldap_handler.LdapHandler, heimdall.Heimdall) {
	opts := options.AllNoodleOptions{}
	opts.NoodleOptions = api.CommandLineOptionsGroups[0].Options.(options.NoodleOptions)
	opts.PostgresOptions = api.CommandLineOptionsGroups[1].Options.(options.PostgresOptions)
	opts.LDAPOptions = api.CommandLineOptionsGroups[0].Options.(options.LDAPOptions)

	api.ServeError = errors.ServeError

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if opts.NoodleOptions.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if opts.NoodleOptions.Config != "" {
		yfile, err := os.ReadFile(opts.NoodleOptions.Config)
		if err != nil {
			Logger.Fatal().Msg(err.Error())
		}

		options, err := options.UnmarshalOptions(yfile)
		if err != nil {
			Logger.Fatal().Msg(err.Error())
		}
		options.NoodleOptions = opts.NoodleOptions
		opts = options
	}

	var db database.Database
	var err error
	if db, err = i.SetupDatabase(opts.PostgresOptions, opts.NoodleOptions.Drop); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	var ldap ldap_handler.LdapHandler
	if ldap, err = i.SetupLDAP(opts.LDAPOptions); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	heimdall := heimdall.NewHeimdall(db)
	return db, ldap, heimdall
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

func (i *ConfigureServerImpl) SetupDatabase(config options.PostgresOptions, drop bool) (database.Database, error) {
	db := database.NewDatabase(config)

	err := db.Connect()
	if err != nil {
		return nil, err
	}

	db.Tables().InitTables(db)

	if drop {
		db.Drop()
	}

	created, _ := db.CheckCreated()
	if !created {
		err = db.Create()
		if err != nil {
			return nil, err
		}
	} else {
		needUpgrade, err := db.CheckUpgrade()
		if err != nil {
			return nil, err
		}

		if needUpgrade {
			err = db.Upgrade()
			if err != nil {
				return nil, err
			}
		}
	}

	db.Tables().UserTable().GetID(-1)

	return db, nil
}

func (i *ConfigureServerImpl) SetupLDAP(config options.LDAPOptions) (ldap_handler.LdapHandler, error) {
	ldap := ldap_handler.NewLdapHandler(ldap_shim.NewLdapShim(), config)
	return ldap, ldap.Connect()
}

func NewConfigureServerImpl() ConfigureServer {
	return &ConfigureServerImpl{}
}
