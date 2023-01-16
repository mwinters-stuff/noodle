// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"

	"github.com/mwinters-stuff/noodle/handlers"
	"github.com/mwinters-stuff/noodle/noodle"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	ldap_shim "github.com/mwinters-stuff/noodle/package-shims/ldap"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/kubernetes"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	Logger = zerologlog.Logger
)

//go:generate swagger generate server --target ../../server --name Noodle --spec ../../swagger.yaml --principal models.Principal

func configureFlags(api *operations.NoodleAPI) {
	opts := &noodle.AllNoodleOptions{}

	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "config",
		LongDescription:  "Noodle Config",
		Options:          &opts.NoodleOptions,
	})
}

func setupDatabase(config yamltypes.AppConfig) (database.Database, error) {
	db := database.NewDatabase(config)

	err := db.Connect()
	if err != nil {
		return nil, err
	}
	db.Tables().InitTables(db)

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

	return db, nil
}

func setupLDAP(config yamltypes.AppConfig) (ldap_handler.LdapHandler, error) {
	ldap := ldap_handler.NewLdapHandler(ldap_shim.NewLdapShim(), config)
	return ldap, ldap.Connect()
}

func configureAPI(api *operations.NoodleAPI) http.Handler {
	// configure the api here
	opts := &noodle.AllNoodleOptions{}
	opts.NoodleOptions = *api.CommandLineOptionsGroups[0].Options.(*noodle.NoodleOptions)

	api.ServeError = errors.ServeError

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if opts.NoodleOptions.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	yfile, err := ioutil.ReadFile(opts.NoodleOptions.Config)
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	config, err := yamltypes.UnmarshalConfig(yfile)
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	var db database.Database
	if db, err = setupDatabase(config); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	var ldap ldap_handler.LdapHandler
	if ldap, err = setupLDAP(config); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = Logger.Debug().Msgf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Remote-User" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		if token == "mathew" {
			prin := models.Principal(token)
			return &prin, nil
		}
		// api.Logger("Access attempt with username: %s", token)
		return nil, errors.New(401, "incorrect username")

	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	// USERS

	api.NoodleAPIGetNoodleUsersHandler = noodle_api.GetNoodleUsersHandlerFunc(func(params noodle_api.GetNoodleUsersParams, principal *models.Principal) middleware.Responder {
		return handlers.HandlerUsers(db, params, principal)
	})

	// GROUPS
	api.NoodleAPIGetNoodleGroupsHandler = noodle_api.GetNoodleGroupsHandlerFunc(func(params noodle_api.GetNoodleGroupsParams, principal *models.Principal) middleware.Responder {
		return handlers.HandlerGroups(db, params, principal)
	})

	// USER GROUPS
	api.NoodleAPIGetNoodleUserGroupsHandler = noodle_api.GetNoodleUserGroupsHandlerFunc(func(params noodle_api.GetNoodleUserGroupsParams, principal *models.Principal) middleware.Responder {
		return handlers.HandlerUserGroups(db, params, principal)
	})

	api.NoodleAPIGetNoodleLdapReloadHandler = noodle_api.GetNoodleLdapReloadHandlerFunc(func(params noodle_api.GetNoodleLdapReloadParams, principal *models.Principal) middleware.Responder {
		return handlers.HandleLDAPRefresh(db, ldap, params, principal)
	})

	api.KubernetesGetHealthzHandler = kubernetes.GetHealthzHandlerFunc(func(params kubernetes.GetHealthzParams) middleware.Responder {
		return kubernetes.NewGetHealthzOK().WithPayload(map[string]string{"status": "OK"})
	})

	api.KubernetesGetReadyzHandler = kubernetes.GetReadyzHandlerFunc(func(params kubernetes.GetReadyzParams) middleware.Responder {
		if db != nil {
			ok, _ := db.CheckCreated()
			if ok {
				return kubernetes.NewGetReadyzOK().WithPayload(map[string]string{"status": "OK"})
			}

		}
		return middleware.Error(http.StatusNotFound, map[string]string{"status": "UNHEALTHY"})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}