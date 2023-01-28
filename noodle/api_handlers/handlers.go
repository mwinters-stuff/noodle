package api_handlers

import (
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
)

func RegisterApiHandlers(api *operations.NoodleAPI, db database.Database, ldap ldap_handler.LdapHandler, heimdall heimdall.Heimdall) {

	RegisterAuthenticationApiHandlers(api, db, ldap)

	// USERS
	RegisterUserApiHandlers(api, db)

	// GROUPS
	RegisterGroupApiHandlers(api, db)

	// USER GROUPS
	RegisterUserGroupApiHandlers(api, db)

	// LDAP
	RegisterLdapApiHandlers(api, db, ldap)

	// HEIMDALL
	RegisterHeimdallApiHandlers(api, db, heimdall)

	// TABS
	RegisterTabApiHandlers(api, db)

	// APP TEMPLATES
	RegisterAppTemplatesApiHandlers(api, db)

	// APPLICATION TABS
	RegisterApplicationTabApiHandlers(api, db)

	// USER APPLICATIONS
	RegisterUserApplicationsApiHandlers(api, db)

	// GROUP APPLICATIONS
	RegisterGroupApplicationsApiHandlers(api, db)

	// APPLICATIONS
	RegisterApplicationsApiHandlers(api, db)
}
