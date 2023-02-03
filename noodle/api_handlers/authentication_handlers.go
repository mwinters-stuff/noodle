package api_handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	openapi_errors "github.com/go-openapi/errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_auth"
)

var (
	RegisterAuthenticationApiHandlers = RegisterAuthenticaionApiHandlersImpl
)

func RegisterAuthenticaionApiHandlersImpl(api *operations.NoodleAPI, db database.Database, ldap ldap_handler.LdapHandler) {
	api.RemoteUserAuth = func(remoteUser string) (*models.Principal, error) {
		return RemoteUserAuthHandler(db, remoteUser)
	}

	api.TokenAuth = func(token string) (*models.Principal, error) {
		return TokenAuthHandler(db, token)
	}

	api.NoodleAuthPostAuthAuthenticateHandler = noodle_auth.PostAuthAuthenticateHandlerFunc(func(params noodle_auth.PostAuthAuthenticateParams) middleware.Responder {
		return HandlerAuthAuthenticationPost(db, ldap, params)
	})

	api.NoodleAuthGetAuthLogoutHandler = noodle_auth.GetAuthLogoutHandlerFunc(func(params noodle_auth.GetAuthLogoutParams, principal *models.Principal) middleware.Responder {
		return HandlerAuthLogoutGet(db, params, principal)
	})
}

var (
	RandToken = RandTokenImpl
)

func RandTokenImpl(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func HandlerAuthAuthenticationPost(db database.Database, ldap ldap_handler.LdapHandler, params noodle_auth.PostAuthAuthenticateParams) middleware.Responder {
	if params.Login.Username == "" || params.Login.Password == "" {
		Logger.Error().Msgf("No Username or Password")
		return noodle_auth.NewPostAuthAuthenticateConflict().WithPayload(&models.Error{Message: "Username or Password empty"})
	}

	ldapUser, err := ldap.GetUser(params.Login.Username)
	if err != nil {
		Logger.Error().Err(err)
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("LDAP Error %s", err.Error())})
	}

	if ldapUser.DN == "" {
		Logger.Error().Msgf("LDAP User %s Not Found", params.Login.Username)
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: "LDAP User not found"})
	}

	ok, err := ldap.AuthUser(ldapUser.DN, params.Login.Password.String())
	if err != nil {

		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("LDAP Error %s", err.Error())})
	}
	if !ok {
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: "Failed Login"})
	}

	dbUser, err := db.Tables().UserTable().GetDN(ldapUser.DN)
	if err != nil {
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("Database Error %s", err.Error())})
	}

	if dbUser.Username != params.Login.Username {
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: "Database Error Username mismatch"})
	}

	token := RandToken(64)

	userSession := models.UserSession{
		UserID: dbUser.ID,
		Token:  token,
	}

	err = db.Tables().UserSessionTable().Insert(&userSession)
	if err != nil {
		return noodle_auth.NewPostAuthAuthenticateUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("Database Error %s", err.Error())})
	}

	return noodle_auth.NewPostAuthAuthenticateOK().WithPayload(&userSession)
}

func HandlerAuthLogoutGet(db database.Database, params noodle_auth.GetAuthLogoutParams, principal *models.Principal) middleware.Responder {
	token := string(*principal)
	if token == "" {
		return noodle_auth.NewGetAuthLogoutUnauthorized().WithPayload(&models.Error{Message: "Invalid Parameters"})
	}
	userSession, err := db.Tables().UserSessionTable().GetToken(token)
	if err != nil {
		return noodle_auth.NewGetAuthLogoutUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("Database Error %s", err.Error())})
	}

	err = db.Tables().UserSessionTable().Delete(userSession.ID)
	if err != nil {
		return noodle_auth.NewGetAuthLogoutUnauthorized().WithPayload(&models.Error{Message: fmt.Sprintf("Database Error %s", err.Error())})
	}

	return noodle_auth.NewGetAuthLogoutOK()

}

func RemoteUserAuthHandler(db database.Database, remoteUser string) (*models.Principal, error) {
	exists, err := db.Tables().UserTable().ExistsUsername(remoteUser)

	if err == nil && exists {
		prin := models.Principal(remoteUser)
		return &prin, nil
	}
	if err != nil {
		return nil, openapi_errors.New(401, err.Error())
	}
	return nil, openapi_errors.New(401, "incorrect username")

}

func TokenAuthHandler(db database.Database, token string) (*models.Principal, error) {
	userSession, err := db.Tables().UserSessionTable().GetToken(token)
	if err != nil {
		return nil, err
	}

	if userSession.Token == "" {
		return nil, openapi_errors.New(401, "session expired")
	}

	prin := models.Principal(token)
	return &prin, nil
}
