// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/kubernetes"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_auth"
)

// NewNoodleAPI creates a new Noodle instance
func NewNoodleAPI(spec *loads.Document) *NoodleAPI {
	return &NoodleAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		NoodleAPIDeleteNoodleApplicationTabsHandler: noodle_api.DeleteNoodleApplicationTabsHandlerFunc(func(params noodle_api.DeleteNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.DeleteNoodleApplicationTabs has not yet been implemented")
		}),
		NoodleAPIDeleteNoodleApplicationsHandler: noodle_api.DeleteNoodleApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.DeleteNoodleApplications has not yet been implemented")
		}),
		NoodleAPIDeleteNoodleGroupApplicationsHandler: noodle_api.DeleteNoodleGroupApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.DeleteNoodleGroupApplications has not yet been implemented")
		}),
		NoodleAPIDeleteNoodleTabsHandler: noodle_api.DeleteNoodleTabsHandlerFunc(func(params noodle_api.DeleteNoodleTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.DeleteNoodleTabs has not yet been implemented")
		}),
		NoodleAPIDeleteNoodleUserApplicationsHandler: noodle_api.DeleteNoodleUserApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.DeleteNoodleUserApplications has not yet been implemented")
		}),
		NoodleAuthGetAuthLogoutHandler: noodle_auth.GetAuthLogoutHandlerFunc(func(params noodle_auth.GetAuthLogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_auth.GetAuthLogout has not yet been implemented")
		}),
		NoodleAuthGetAuthSessionHandler: noodle_auth.GetAuthSessionHandlerFunc(func(params noodle_auth.GetAuthSessionParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_auth.GetAuthSession has not yet been implemented")
		}),
		KubernetesGetHealthzHandler: kubernetes.GetHealthzHandlerFunc(func(params kubernetes.GetHealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation kubernetes.GetHealthz has not yet been implemented")
		}),
		NoodleAPIGetNoodleAppTemplatesHandler: noodle_api.GetNoodleAppTemplatesHandlerFunc(func(params noodle_api.GetNoodleAppTemplatesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleAppTemplates has not yet been implemented")
		}),
		NoodleAPIGetNoodleApplicationTabsHandler: noodle_api.GetNoodleApplicationTabsHandlerFunc(func(params noodle_api.GetNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleApplicationTabs has not yet been implemented")
		}),
		NoodleAPIGetNoodleApplicationsHandler: noodle_api.GetNoodleApplicationsHandlerFunc(func(params noodle_api.GetNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleApplications has not yet been implemented")
		}),
		NoodleAPIGetNoodleGroupApplicationsHandler: noodle_api.GetNoodleGroupApplicationsHandlerFunc(func(params noodle_api.GetNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleGroupApplications has not yet been implemented")
		}),
		NoodleAPIGetNoodleGroupsHandler: noodle_api.GetNoodleGroupsHandlerFunc(func(params noodle_api.GetNoodleGroupsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleGroups has not yet been implemented")
		}),
		NoodleAPIGetNoodleHeimdallReloadHandler: noodle_api.GetNoodleHeimdallReloadHandlerFunc(func(params noodle_api.GetNoodleHeimdallReloadParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleHeimdallReload has not yet been implemented")
		}),
		NoodleAPIGetNoodleLdapReloadHandler: noodle_api.GetNoodleLdapReloadHandlerFunc(func(params noodle_api.GetNoodleLdapReloadParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleLdapReload has not yet been implemented")
		}),
		NoodleAPIGetNoodleTabsHandler: noodle_api.GetNoodleTabsHandlerFunc(func(params noodle_api.GetNoodleTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleTabs has not yet been implemented")
		}),
		NoodleAPIGetNoodleUserAllowedApplicationsHandler: noodle_api.GetNoodleUserAllowedApplicationsHandlerFunc(func(params noodle_api.GetNoodleUserAllowedApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleUserAllowedApplications has not yet been implemented")
		}),
		NoodleAPIGetNoodleUserApplicationsHandler: noodle_api.GetNoodleUserApplicationsHandlerFunc(func(params noodle_api.GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleUserApplications has not yet been implemented")
		}),
		NoodleAPIGetNoodleUserGroupsHandler: noodle_api.GetNoodleUserGroupsHandlerFunc(func(params noodle_api.GetNoodleUserGroupsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleUserGroups has not yet been implemented")
		}),
		NoodleAPIGetNoodleUsersHandler: noodle_api.GetNoodleUsersHandlerFunc(func(params noodle_api.GetNoodleUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.GetNoodleUsers has not yet been implemented")
		}),
		KubernetesGetReadyzHandler: kubernetes.GetReadyzHandlerFunc(func(params kubernetes.GetReadyzParams) middleware.Responder {
			return middleware.NotImplemented("operation kubernetes.GetReadyz has not yet been implemented")
		}),
		NoodleAuthPostAuthAuthenticateHandler: noodle_auth.PostAuthAuthenticateHandlerFunc(func(params noodle_auth.PostAuthAuthenticateParams) middleware.Responder {
			return middleware.NotImplemented("operation noodle_auth.PostAuthAuthenticate has not yet been implemented")
		}),
		NoodleAPIPostNoodleApplicationTabsHandler: noodle_api.PostNoodleApplicationTabsHandlerFunc(func(params noodle_api.PostNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.PostNoodleApplicationTabs has not yet been implemented")
		}),
		NoodleAPIPostNoodleApplicationsHandler: noodle_api.PostNoodleApplicationsHandlerFunc(func(params noodle_api.PostNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.PostNoodleApplications has not yet been implemented")
		}),
		NoodleAPIPostNoodleGroupApplicationsHandler: noodle_api.PostNoodleGroupApplicationsHandlerFunc(func(params noodle_api.PostNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.PostNoodleGroupApplications has not yet been implemented")
		}),
		NoodleAPIPostNoodleTabsHandler: noodle_api.PostNoodleTabsHandlerFunc(func(params noodle_api.PostNoodleTabsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.PostNoodleTabs has not yet been implemented")
		}),
		NoodleAPIPostNoodleUserApplicationsHandler: noodle_api.PostNoodleUserApplicationsHandlerFunc(func(params noodle_api.PostNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation noodle_api.PostNoodleUserApplications has not yet been implemented")
		}),

		// Applies when the "Remote-User" header is set
		RemoteUserAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (remote-user) Remote-User from header param [Remote-User] has not yet been implemented")
		},
		// Applies when the "X-Token" header is set
		TokenAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (token) X-Token from header param [X-Token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*NoodleAPI Noodle */
type NoodleAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// RemoteUserAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Remote-User provided in the header
	RemoteUserAuth func(string) (*models.Principal, error)

	// TokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Token provided in the header
	TokenAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// NoodleAPIDeleteNoodleApplicationTabsHandler sets the operation handler for the delete noodle application tabs operation
	NoodleAPIDeleteNoodleApplicationTabsHandler noodle_api.DeleteNoodleApplicationTabsHandler
	// NoodleAPIDeleteNoodleApplicationsHandler sets the operation handler for the delete noodle applications operation
	NoodleAPIDeleteNoodleApplicationsHandler noodle_api.DeleteNoodleApplicationsHandler
	// NoodleAPIDeleteNoodleGroupApplicationsHandler sets the operation handler for the delete noodle group applications operation
	NoodleAPIDeleteNoodleGroupApplicationsHandler noodle_api.DeleteNoodleGroupApplicationsHandler
	// NoodleAPIDeleteNoodleTabsHandler sets the operation handler for the delete noodle tabs operation
	NoodleAPIDeleteNoodleTabsHandler noodle_api.DeleteNoodleTabsHandler
	// NoodleAPIDeleteNoodleUserApplicationsHandler sets the operation handler for the delete noodle user applications operation
	NoodleAPIDeleteNoodleUserApplicationsHandler noodle_api.DeleteNoodleUserApplicationsHandler
	// NoodleAuthGetAuthLogoutHandler sets the operation handler for the get auth logout operation
	NoodleAuthGetAuthLogoutHandler noodle_auth.GetAuthLogoutHandler
	// NoodleAuthGetAuthSessionHandler sets the operation handler for the get auth session operation
	NoodleAuthGetAuthSessionHandler noodle_auth.GetAuthSessionHandler
	// KubernetesGetHealthzHandler sets the operation handler for the get healthz operation
	KubernetesGetHealthzHandler kubernetes.GetHealthzHandler
	// NoodleAPIGetNoodleAppTemplatesHandler sets the operation handler for the get noodle app templates operation
	NoodleAPIGetNoodleAppTemplatesHandler noodle_api.GetNoodleAppTemplatesHandler
	// NoodleAPIGetNoodleApplicationTabsHandler sets the operation handler for the get noodle application tabs operation
	NoodleAPIGetNoodleApplicationTabsHandler noodle_api.GetNoodleApplicationTabsHandler
	// NoodleAPIGetNoodleApplicationsHandler sets the operation handler for the get noodle applications operation
	NoodleAPIGetNoodleApplicationsHandler noodle_api.GetNoodleApplicationsHandler
	// NoodleAPIGetNoodleGroupApplicationsHandler sets the operation handler for the get noodle group applications operation
	NoodleAPIGetNoodleGroupApplicationsHandler noodle_api.GetNoodleGroupApplicationsHandler
	// NoodleAPIGetNoodleGroupsHandler sets the operation handler for the get noodle groups operation
	NoodleAPIGetNoodleGroupsHandler noodle_api.GetNoodleGroupsHandler
	// NoodleAPIGetNoodleHeimdallReloadHandler sets the operation handler for the get noodle heimdall reload operation
	NoodleAPIGetNoodleHeimdallReloadHandler noodle_api.GetNoodleHeimdallReloadHandler
	// NoodleAPIGetNoodleLdapReloadHandler sets the operation handler for the get noodle ldap reload operation
	NoodleAPIGetNoodleLdapReloadHandler noodle_api.GetNoodleLdapReloadHandler
	// NoodleAPIGetNoodleTabsHandler sets the operation handler for the get noodle tabs operation
	NoodleAPIGetNoodleTabsHandler noodle_api.GetNoodleTabsHandler
	// NoodleAPIGetNoodleUserAllowedApplicationsHandler sets the operation handler for the get noodle user allowed applications operation
	NoodleAPIGetNoodleUserAllowedApplicationsHandler noodle_api.GetNoodleUserAllowedApplicationsHandler
	// NoodleAPIGetNoodleUserApplicationsHandler sets the operation handler for the get noodle user applications operation
	NoodleAPIGetNoodleUserApplicationsHandler noodle_api.GetNoodleUserApplicationsHandler
	// NoodleAPIGetNoodleUserGroupsHandler sets the operation handler for the get noodle user groups operation
	NoodleAPIGetNoodleUserGroupsHandler noodle_api.GetNoodleUserGroupsHandler
	// NoodleAPIGetNoodleUsersHandler sets the operation handler for the get noodle users operation
	NoodleAPIGetNoodleUsersHandler noodle_api.GetNoodleUsersHandler
	// KubernetesGetReadyzHandler sets the operation handler for the get readyz operation
	KubernetesGetReadyzHandler kubernetes.GetReadyzHandler
	// NoodleAuthPostAuthAuthenticateHandler sets the operation handler for the post auth authenticate operation
	NoodleAuthPostAuthAuthenticateHandler noodle_auth.PostAuthAuthenticateHandler
	// NoodleAPIPostNoodleApplicationTabsHandler sets the operation handler for the post noodle application tabs operation
	NoodleAPIPostNoodleApplicationTabsHandler noodle_api.PostNoodleApplicationTabsHandler
	// NoodleAPIPostNoodleApplicationsHandler sets the operation handler for the post noodle applications operation
	NoodleAPIPostNoodleApplicationsHandler noodle_api.PostNoodleApplicationsHandler
	// NoodleAPIPostNoodleGroupApplicationsHandler sets the operation handler for the post noodle group applications operation
	NoodleAPIPostNoodleGroupApplicationsHandler noodle_api.PostNoodleGroupApplicationsHandler
	// NoodleAPIPostNoodleTabsHandler sets the operation handler for the post noodle tabs operation
	NoodleAPIPostNoodleTabsHandler noodle_api.PostNoodleTabsHandler
	// NoodleAPIPostNoodleUserApplicationsHandler sets the operation handler for the post noodle user applications operation
	NoodleAPIPostNoodleUserApplicationsHandler noodle_api.PostNoodleUserApplicationsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *NoodleAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *NoodleAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *NoodleAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *NoodleAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *NoodleAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *NoodleAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *NoodleAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *NoodleAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *NoodleAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the NoodleAPI
func (o *NoodleAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RemoteUserAuth == nil {
		unregistered = append(unregistered, "RemoteUserAuth")
	}
	if o.TokenAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.NoodleAPIDeleteNoodleApplicationTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.DeleteNoodleApplicationTabsHandler")
	}
	if o.NoodleAPIDeleteNoodleApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.DeleteNoodleApplicationsHandler")
	}
	if o.NoodleAPIDeleteNoodleGroupApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.DeleteNoodleGroupApplicationsHandler")
	}
	if o.NoodleAPIDeleteNoodleTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.DeleteNoodleTabsHandler")
	}
	if o.NoodleAPIDeleteNoodleUserApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.DeleteNoodleUserApplicationsHandler")
	}
	if o.NoodleAuthGetAuthLogoutHandler == nil {
		unregistered = append(unregistered, "noodle_auth.GetAuthLogoutHandler")
	}
	if o.NoodleAuthGetAuthSessionHandler == nil {
		unregistered = append(unregistered, "noodle_auth.GetAuthSessionHandler")
	}
	if o.KubernetesGetHealthzHandler == nil {
		unregistered = append(unregistered, "kubernetes.GetHealthzHandler")
	}
	if o.NoodleAPIGetNoodleAppTemplatesHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleAppTemplatesHandler")
	}
	if o.NoodleAPIGetNoodleApplicationTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleApplicationTabsHandler")
	}
	if o.NoodleAPIGetNoodleApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleApplicationsHandler")
	}
	if o.NoodleAPIGetNoodleGroupApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleGroupApplicationsHandler")
	}
	if o.NoodleAPIGetNoodleGroupsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleGroupsHandler")
	}
	if o.NoodleAPIGetNoodleHeimdallReloadHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleHeimdallReloadHandler")
	}
	if o.NoodleAPIGetNoodleLdapReloadHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleLdapReloadHandler")
	}
	if o.NoodleAPIGetNoodleTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleTabsHandler")
	}
	if o.NoodleAPIGetNoodleUserAllowedApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleUserAllowedApplicationsHandler")
	}
	if o.NoodleAPIGetNoodleUserApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleUserApplicationsHandler")
	}
	if o.NoodleAPIGetNoodleUserGroupsHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleUserGroupsHandler")
	}
	if o.NoodleAPIGetNoodleUsersHandler == nil {
		unregistered = append(unregistered, "noodle_api.GetNoodleUsersHandler")
	}
	if o.KubernetesGetReadyzHandler == nil {
		unregistered = append(unregistered, "kubernetes.GetReadyzHandler")
	}
	if o.NoodleAuthPostAuthAuthenticateHandler == nil {
		unregistered = append(unregistered, "noodle_auth.PostAuthAuthenticateHandler")
	}
	if o.NoodleAPIPostNoodleApplicationTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.PostNoodleApplicationTabsHandler")
	}
	if o.NoodleAPIPostNoodleApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.PostNoodleApplicationsHandler")
	}
	if o.NoodleAPIPostNoodleGroupApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.PostNoodleGroupApplicationsHandler")
	}
	if o.NoodleAPIPostNoodleTabsHandler == nil {
		unregistered = append(unregistered, "noodle_api.PostNoodleTabsHandler")
	}
	if o.NoodleAPIPostNoodleUserApplicationsHandler == nil {
		unregistered = append(unregistered, "noodle_api.PostNoodleUserApplicationsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *NoodleAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *NoodleAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "remote-user":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.RemoteUserAuth(token)
			})

		case "token":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.TokenAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *NoodleAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *NoodleAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *NoodleAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *NoodleAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the noodle API
func (o *NoodleAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *NoodleAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/noodle/application-tabs"] = noodle_api.NewDeleteNoodleApplicationTabs(o.context, o.NoodleAPIDeleteNoodleApplicationTabsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/noodle/applications"] = noodle_api.NewDeleteNoodleApplications(o.context, o.NoodleAPIDeleteNoodleApplicationsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/noodle/group-applications"] = noodle_api.NewDeleteNoodleGroupApplications(o.context, o.NoodleAPIDeleteNoodleGroupApplicationsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/noodle/tabs"] = noodle_api.NewDeleteNoodleTabs(o.context, o.NoodleAPIDeleteNoodleTabsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/noodle/user-applications"] = noodle_api.NewDeleteNoodleUserApplications(o.context, o.NoodleAPIDeleteNoodleUserApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/logout"] = noodle_auth.NewGetAuthLogout(o.context, o.NoodleAuthGetAuthLogoutHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/session"] = noodle_auth.NewGetAuthSession(o.context, o.NoodleAuthGetAuthSessionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthz"] = kubernetes.NewGetHealthz(o.context, o.KubernetesGetHealthzHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/app-templates"] = noodle_api.NewGetNoodleAppTemplates(o.context, o.NoodleAPIGetNoodleAppTemplatesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/application-tabs"] = noodle_api.NewGetNoodleApplicationTabs(o.context, o.NoodleAPIGetNoodleApplicationTabsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/applications"] = noodle_api.NewGetNoodleApplications(o.context, o.NoodleAPIGetNoodleApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/group-applications"] = noodle_api.NewGetNoodleGroupApplications(o.context, o.NoodleAPIGetNoodleGroupApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/groups"] = noodle_api.NewGetNoodleGroups(o.context, o.NoodleAPIGetNoodleGroupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/heimdall/reload"] = noodle_api.NewGetNoodleHeimdallReload(o.context, o.NoodleAPIGetNoodleHeimdallReloadHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/ldap/reload"] = noodle_api.NewGetNoodleLdapReload(o.context, o.NoodleAPIGetNoodleLdapReloadHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/tabs"] = noodle_api.NewGetNoodleTabs(o.context, o.NoodleAPIGetNoodleTabsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/user-allowed-applications"] = noodle_api.NewGetNoodleUserAllowedApplications(o.context, o.NoodleAPIGetNoodleUserAllowedApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/user-applications"] = noodle_api.NewGetNoodleUserApplications(o.context, o.NoodleAPIGetNoodleUserApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/user-groups"] = noodle_api.NewGetNoodleUserGroups(o.context, o.NoodleAPIGetNoodleUserGroupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/noodle/users"] = noodle_api.NewGetNoodleUsers(o.context, o.NoodleAPIGetNoodleUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/readyz"] = kubernetes.NewGetReadyz(o.context, o.KubernetesGetReadyzHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/authenticate"] = noodle_auth.NewPostAuthAuthenticate(o.context, o.NoodleAuthPostAuthAuthenticateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/noodle/application-tabs"] = noodle_api.NewPostNoodleApplicationTabs(o.context, o.NoodleAPIPostNoodleApplicationTabsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/noodle/applications"] = noodle_api.NewPostNoodleApplications(o.context, o.NoodleAPIPostNoodleApplicationsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/noodle/group-applications"] = noodle_api.NewPostNoodleGroupApplications(o.context, o.NoodleAPIPostNoodleGroupApplicationsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/noodle/tabs"] = noodle_api.NewPostNoodleTabs(o.context, o.NoodleAPIPostNoodleTabsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/noodle/user-applications"] = noodle_api.NewPostNoodleUserApplications(o.context, o.NoodleAPIPostNoodleUserApplicationsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *NoodleAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *NoodleAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *NoodleAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *NoodleAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *NoodleAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
