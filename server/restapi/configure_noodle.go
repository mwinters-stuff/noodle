// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/kubernetes"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle"
)

//go:generate swagger generate server --target ../../server --name Noodle --spec ../../swagger.yaml --principal models.Principal

func configureFlags(api *operations.NoodleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NoodleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

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

	if api.NoodleDeleteNoodleUsersHandler == nil {
		api.NoodleDeleteNoodleUsersHandler = noodle.DeleteNoodleUsersHandlerFunc(func(params noodle.DeleteNoodleUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle.DeleteNoodleUsers has not yet been implemented")
		})
	}
	if api.KubernetesGetHealthzHandler == nil {
		api.KubernetesGetHealthzHandler = kubernetes.GetHealthzHandlerFunc(func(params kubernetes.GetHealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation kubernetes.GetHealthz has not yet been implemented")
		})
	}

	api.NoodleGetNoodleUsersHandler = noodle.GetNoodleUsersHandlerFunc(func(params noodle.GetNoodleUsersParams, principal *models.Principal) middleware.Responder {
		return noodle.NewGetNoodleUsersOK().WithPayload([]*models.User{
			{
				DN:          "abcdef",
				DisplayName: "abcdef",
				GivenName:   "bob",
				ID:          1,
				Surname:     "joe",
				UIDNumber:   0,
				Username:    "abcdef",
			},
		})
	})

	if api.KubernetesGetReadyzHandler == nil {
		api.KubernetesGetReadyzHandler = kubernetes.GetReadyzHandlerFunc(func(params kubernetes.GetReadyzParams) middleware.Responder {
			return middleware.NotImplemented("operation kubernetes.GetReadyz has not yet been implemented")
		})
	}
	if api.NoodlePostNoodleUsersHandler == nil {
		api.NoodlePostNoodleUsersHandler = noodle.PostNoodleUsersHandlerFunc(func(params noodle.PostNoodleUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle.PostNoodleUsers has not yet been implemented")
		})
	}

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
