package api_handlers_test

import (
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	heimdall_mocks "github.com/mwinters-stuff/noodle/noodle/heimdall/mocks"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HandlersTestSuite struct {
	suite.Suite
	mockDatabase *mocks.Database
	mockLdap     *ldap_mocks.LdapHandler
	mockHiemdall *heimdall_mocks.Heimdall

	api *operations.NoodleAPI
}

func (suite *HandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())
	suite.mockHiemdall = heimdall_mocks.NewHeimdall(suite.T())

	suite.api = &operations.NoodleAPI{}

}

func (suite *HandlersTestSuite) TestRegisterAPIHandlers() {

	api_handlers.RegisterApiHandlers(suite.api, suite.mockDatabase, suite.mockLdap, suite.mockHiemdall)

	require.NotNil(suite.T(), suite.api.NoodleAuthGetAuthLogoutHandler)
	require.NotNil(suite.T(), suite.api.NoodleAuthPostAuthAuthenticateHandler)
	require.NotNil(suite.T(), suite.api.RemoteUserAuth)
	require.NotNil(suite.T(), suite.api.TokenAuth)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleAppTemplatesHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleApplicationsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleApplicationTabsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleApplicationTabsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleApplicationTabsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleGroupApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleGroupApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleGroupApplicationsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleGroupsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleHeimdallReloadHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleLdapReloadHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleTabsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleTabsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleTabsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUserApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleUserApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleUserApplicationsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUserGroupsHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUsersHandler)

	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleUploadIconHandler)
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
