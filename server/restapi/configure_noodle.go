// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	zerologlog "github.com/rs/zerolog/log"

	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/configure_server"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/kubernetes"
)

var (
	Logger = zerologlog.Logger
)

//go:generate swagger generate server --target ../../server --name Noodle --spec ../../swagger/noodle_service.yaml --principal models.Principal

func configureFlags(api *operations.NoodleAPI) {
	serverConfig := configure_server.NewConfigureServer()
	serverConfig.ConfigureFlags(api)
}

func configureAPI(api *operations.NoodleAPI) http.Handler {
	// configure the api here
	serverConfig := configure_server.NewConfigureServer()
	db, ldap, heimdall, err := serverConfig.ConfigureAPI(api)
	if err != nil {
		Logger.Fatal().Err(err)
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

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api_handlers.RegisterApiHandlers(api, db, ldap, heimdall)
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Info().Msgf("Middlware %s", r.URL)
		handler.ServeHTTP(w, r)
	})
	// return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler

	return handleCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") || r.URL.Path == "/swagger.json" {
			handler.ServeHTTP(w, r)
		} else {
			Logger.Info().Msgf("Serve %s", r.URL)
			http.FileServer(http.Dir("/workspaces/noodle/web-client/dist/")).ServeHTTP(w, r)
		}
	}))

}
