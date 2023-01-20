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

	"github.com/mwinters-stuff/noodle/noodle"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
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
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "drop",
		LongDescription:  "Drop Database",
		Options:          &opts.NoodleOptions,
	})
}

func setupDatabase(config yamltypes.AppConfig, drop bool) (database.Database, error) {
	db := database.NewDatabase(config)

	err := db.Connect()
	if err != nil {
		return nil, err
	}

	db.Tables().InitTables(db)

	// if drop {
	// 	db.Drop()
	// }

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
	if db, err = setupDatabase(config, opts.NoodleOptions.Drop); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	var ldap ldap_handler.LdapHandler
	if ldap, err = setupLDAP(config); err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	heimdall := heimdall.NewHeimdall(db)

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
		return api_handlers.HandlerUsers(db, params, principal)
	})

	// GROUPS
	api.NoodleAPIGetNoodleGroupsHandler = noodle_api.GetNoodleGroupsHandlerFunc(func(params noodle_api.GetNoodleGroupsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerGroups(db, params, principal)
	})

	// USER GROUPS
	api.NoodleAPIGetNoodleUserGroupsHandler = noodle_api.GetNoodleUserGroupsHandlerFunc(func(params noodle_api.GetNoodleUserGroupsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerUserGroups(db, params, principal)
	})

	// LDAP
	api.NoodleAPIGetNoodleLdapReloadHandler = noodle_api.GetNoodleLdapReloadHandlerFunc(func(params noodle_api.GetNoodleLdapReloadParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandleLDAPRefresh(db, ldap, params, principal)
	})

	// HEIMDALL
	api.NoodleAPIGetNoodleHeimdallReloadHandler = noodle_api.GetNoodleHeimdallReloadHandlerFunc(func(params noodle_api.GetNoodleHeimdallReloadParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandleHeimdallRefresh(db, heimdall, params, principal)
	})

	// TABS
	api.NoodleAPIGetNoodleTabsHandler = noodle_api.GetNoodleTabsHandlerFunc(func(params noodle_api.GetNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerTabGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleTabsHandler = noodle_api.PostNoodleTabsHandlerFunc(func(params noodle_api.PostNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerTabPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleTabsHandler = noodle_api.DeleteNoodleTabsHandlerFunc(func(params noodle_api.DeleteNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerTabDelete(db, params, principal)
	})

	// APP TEMPLATES
	api.NoodleAPIGetNoodleAppTemplatesHandler = noodle_api.GetNoodleAppTemplatesHandlerFunc(func(params noodle_api.GetNoodleAppTemplatesParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerAppTemplates(db, params, principal)
	})

	// TABS
	api.NoodleAPIGetNoodleApplicationTabsHandler = noodle_api.GetNoodleApplicationTabsHandlerFunc(func(params noodle_api.GetNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerApplicationTabGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleApplicationTabsHandler = noodle_api.PostNoodleApplicationTabsHandlerFunc(func(params noodle_api.PostNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerApplicationTabPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleApplicationTabsHandler = noodle_api.DeleteNoodleApplicationTabsHandlerFunc(func(params noodle_api.DeleteNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerApplicationTabDelete(db, params, principal)
	})

	// USER APPLICATIONS
	api.NoodleAPIGetNoodleUserApplicationsHandler = noodle_api.GetNoodleUserApplicationsHandlerFunc(func(params noodle_api.GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerUserApplicationGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleUserApplicationsHandler = noodle_api.PostNoodleUserApplicationsHandlerFunc(func(params noodle_api.PostNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerUserApplicationPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleUserApplicationsHandler = noodle_api.DeleteNoodleUserApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerUserApplicationDelete(db, params, principal)
	})

	// GROUP APPLICATIONS
	api.NoodleAPIGetNoodleGroupApplicationsHandler = noodle_api.GetNoodleGroupApplicationsHandlerFunc(func(params noodle_api.GetNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerGroupApplicationGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleGroupApplicationsHandler = noodle_api.PostNoodleGroupApplicationsHandlerFunc(func(params noodle_api.PostNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerGroupApplicationPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleGroupApplicationsHandler = noodle_api.DeleteNoodleGroupApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
		return api_handlers.HandlerGroupApplicationDelete(db, params, principal)
	})

	//KUBE

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
