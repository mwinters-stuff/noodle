package api_handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_auth"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AuthenticationHandlersTestSuite struct {
	suite.Suite
	mockLdap             *ldap_mocks.LdapHandler
	mockDatabase         *mocks.Database
	mockTables           *mocks.Tables
	mockUserTable        *mocks.UserTable
	mockUserSessionTable *mocks.UserSessionTable
	api                  *operations.NoodleAPI
}

func (suite *AuthenticationHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
	api_handlers.Logger = log.Output(nil)
}

func (suite *AuthenticationHandlersTestSuite) SetupTest() {
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserSessionTable = mocks.NewUserSessionTable(suite.T())
	suite.api = &operations.NoodleAPI{}

	api_handlers.RegisterAuthenticationApiHandlers(suite.api, suite.mockDatabase, suite.mockLdap)

}

func (suite *AuthenticationHandlersTestSuite) TearDownTest() {
	api_handlers.RandToken = api_handlers.RandTokenImpl
}

func (suite *AuthenticationHandlersTestSuite) TearSuite() {

}

func (suite *AuthenticationHandlersTestSuite) TestRemoteUserAuthHandler() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)
	suite.mockUserTable.EXPECT().ExistsUsername("bob").Return(true, nil)

	principal, err := suite.api.RemoteUserAuth("bob")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), principal)
	require.Equal(suite.T(), models.Principal("bob"), *principal)
}

func (suite *AuthenticationHandlersTestSuite) TestRemoteUserAuthHandlerErrorUsername() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)
	suite.mockUserTable.EXPECT().ExistsUsername("bob").Return(false, nil)

	principal, err := suite.api.RemoteUserAuth("bob")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "incorrect username", err.Error())
	require.Nil(suite.T(), principal)
}

func (suite *AuthenticationHandlersTestSuite) TestRemoteUserAuthHandlerErrorDatabase() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)
	suite.mockUserTable.EXPECT().ExistsUsername("bob").Return(false, errors.New("failed"))

	principal, err := suite.api.RemoteUserAuth("bob")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "failed", err.Error())
	require.Nil(suite.T(), principal)
}

func (suite *AuthenticationHandlersTestSuite) TestTokenAuthHandler() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{
		ID:      100,
		UserID:  2,
		Token:   "tokentoken",
		Issued:  strfmt.NewDateTime(),
		Expires: strfmt.NewDateTime(),
	}, nil)

	principal, err := suite.api.TokenAuth("tokentoken")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), principal)
	require.Equal(suite.T(), models.Principal("tokentoken"), *principal)
}

func (suite *AuthenticationHandlersTestSuite) TestTokenAuthHandlerSessionExpired() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{}, nil)

	principal, err := suite.api.TokenAuth("tokentoken")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "session expired", err.Error())
	require.Nil(suite.T(), principal)

}

func (suite *AuthenticationHandlersTestSuite) TestTokenAuthHandlerDatabaseError() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{}, errors.New("failed"))

	principal, err := suite.api.TokenAuth("tokentoken")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "failed", err.Error())
	require.Nil(suite.T(), principal)

}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthLogoutGet() {
	suite.mockDatabase.EXPECT().Tables().Twice().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Twice().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{
		ID:      100,
		UserID:  2,
		Token:   "tokentoken",
		Issued:  strfmt.NewDateTime(),
		Expires: strfmt.NewDateTime(),
	}, nil)
	suite.mockUserSessionTable.EXPECT().Delete(int64(100)).Once().Return(nil)

	principal := models.Principal("tokentoken")
	response := suite.api.NoodleAuthGetAuthLogoutHandler.Handle(noodle_auth.NewGetAuthLogoutParams(), &principal)
	require.NotNil(suite.T(), response)

	mockHeader := http.Header{}
	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().Header().Once().Return(mockHeader)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthLogoutGetNoToken() {
	principal := models.Principal("")
	response := suite.api.NoodleAuthGetAuthLogoutHandler.Handle(noodle_auth.NewGetAuthLogoutParams(), &principal)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Invalid Parameters"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthLogoutGetDBError1() {
	suite.mockDatabase.EXPECT().Tables().Once().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{}, errors.New("failed"))

	principal := models.Principal("tokentoken")
	response := suite.api.NoodleAuthGetAuthLogoutHandler.Handle(noodle_auth.NewGetAuthLogoutParams(), &principal)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Database Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthLogoutGetDBError2() {
	suite.mockDatabase.EXPECT().Tables().Twice().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserSessionTable().Twice().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().GetToken("tokentoken").Return(models.UserSession{
		ID:      100,
		UserID:  2,
		Token:   "tokentoken",
		Issued:  strfmt.NewDateTime(),
		Expires: strfmt.NewDateTime(),
	}, nil)
	suite.mockUserSessionTable.EXPECT().Delete(int64(100)).Once().Return(errors.New("failed"))

	principal := models.Principal("tokentoken")
	response := suite.api.NoodleAuthGetAuthLogoutHandler.Handle(noodle_auth.NewGetAuthLogoutParams(), &principal)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Database Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func UnRandomToken(n int) string {
	return "5754d020201b57a27106acd96bffa9f07948366002eb48399834f59341e960af266d91cc27213f45a8317d867544594a4ab3039fbf1880d0568ca7097fb1c587"
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPost() {
	suite.mockDatabase.EXPECT().Tables().Twice().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)

	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(true, nil)

	suite.mockUserTable.EXPECT().GetDN("uid=bob,ou=people,dc=example,dc=nz").Once().Return(models.User{
		ID:          100,
		Username:    "bob",
		DisplayName: "Bob Knob",
	}, nil)

	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().Insert(&models.UserSession{
		UserID: 100,
		Token:  "5754d020201b57a27106acd96bffa9f07948366002eb48399834f59341e960af266d91cc27213f45a8317d867544594a4ab3039fbf1880d0568ca7097fb1c587",
	}).Return(nil)

	api_handlers.RandToken = UnRandomToken

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(200).Once()
	mockWriter.EXPECT().Write([]byte(`{"displayName":"Bob Knob","token":"5754d020201b57a27106acd96bffa9f07948366002eb48399834f59341e960af266d91cc27213f45a8317d867544594a4ab3039fbf1880d0568ca7097fb1c587"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorNoParams() {
	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = ""
	params.Auth.Username = ""

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Username or Password empty"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorLDAPGetUserError() {
	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{}, errors.New("failed"))

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"LDAP Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorLDAPGetUserNotFound() {
	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{}, nil)

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"LDAP User not found"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorLDAPAuthUserError() {
	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(false, errors.New("failed"))

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"LDAP Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorLDAPAuthUserFailed() {
	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(false, nil)

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Failed Login"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorGetDNDBError() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)

	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(true, nil)

	suite.mockUserTable.EXPECT().GetDN("uid=bob,ou=people,dc=example,dc=nz").Once().Return(models.User{}, errors.New("failed"))

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Database Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostErrorGetDNUsernameMismatch() {
	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)

	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(true, nil)

	suite.mockUserTable.EXPECT().GetDN("uid=bob,ou=people,dc=example,dc=nz").Once().Return(models.User{}, nil)

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Database Error Username mismatch"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestHandlerAuthAuthenticationPostInsertUserSessionFailed() {
	suite.mockDatabase.EXPECT().Tables().Twice().Return(suite.mockTables)
	suite.mockTables.EXPECT().UserTable().Once().Return(suite.mockUserTable)

	suite.mockLdap.EXPECT().GetUser("bob").Once().Return(models.User{
		Username: "bob",
		DN:       "uid=bob,ou=people,dc=example,dc=nz",
	}, nil)

	suite.mockLdap.EXPECT().AuthUser("uid=bob,ou=people,dc=example,dc=nz", "letmein").Once().Return(true, nil)
	suite.mockUserTable.EXPECT().GetDN("uid=bob,ou=people,dc=example,dc=nz").Once().Return(models.User{
		ID:          100,
		Username:    "bob",
		DisplayName: "Bob Knob",
	}, nil)

	suite.mockTables.EXPECT().UserSessionTable().Once().Return(suite.mockUserSessionTable)
	suite.mockUserSessionTable.EXPECT().Insert(&models.UserSession{
		UserID: 100,
		Token:  "5754d020201b57a27106acd96bffa9f07948366002eb48399834f59341e960af266d91cc27213f45a8317d867544594a4ab3039fbf1880d0568ca7097fb1c587",
	}).Return(errors.New("failed"))

	api_handlers.RandToken = UnRandomToken

	params := noodle_auth.NewPostAuthAuthenticateParams()
	params.Auth.Password = "letmein"
	params.Auth.Username = "bob"

	response := suite.api.NoodleAuthPostAuthAuthenticateHandler.Handle(params)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(401).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"Database Error failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AuthenticationHandlersTestSuite) TestRandTokenOk() {
	random := api_handlers.RandToken(64)
	require.Len(suite.T(), random, 128)
}

func TestAuthenticationHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationHandlersTestSuite))
}
