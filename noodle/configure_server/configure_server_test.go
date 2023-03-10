package configure_server_test

import (
	"errors"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/configure_server"
	"github.com/mwinters-stuff/noodle/noodle/database"
	database_mocks "github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	heimdall_mocks "github.com/mwinters-stuff/noodle/noodle/heimdall/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/noodle/options"
	ldap_shim "github.com/mwinters-stuff/noodle/package-shims/ldap"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type configureServerLogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *configureServerLogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type ConfigureServerTestSuite struct {
	suite.Suite
	loghook configureServerLogHook

	mockHeimdall *heimdall_mocks.Heimdall
	mockDatabase *database_mocks.Database
	mockLdap     *ldap_mocks.LdapHandler
	mockTables   *database_mocks.Tables

	tempFile *os.File
}

func (suite *ConfigureServerTestSuite) SetupSuite() {
	suite.loghook = configureServerLogHook{}
	configure_server.Logger = log.Hook(&suite.loghook).Output(nil)

	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())
	suite.mockHeimdall = heimdall_mocks.NewHeimdall(suite.T())
	suite.mockDatabase = database_mocks.NewDatabase(suite.T())
	suite.mockTables = database_mocks.NewTables(suite.T())
}

func (suite *ConfigureServerTestSuite) SetupTest() {

	database.NewDatabase = func(pgConfig options.PostgresOptions) database.Database { return suite.mockDatabase }
	ldap_handler.NewLdapHandler = func(ldapShim ldap_shim.LdapShim, ldapConfig options.LDAPOptions) ldap_handler.LdapHandler {
		return suite.mockLdap
	}
	heimdall.NewHeimdall = func(database database.Database, options options.NoodleOptions) heimdall.Heimdall {
		return suite.mockHeimdall
	}

	suite.tempFile = nil
}

func (suite *ConfigureServerTestSuite) TearDownTest() {
	database.NewDatabase = database.NewDatabaseImpl
	ldap_handler.NewLdapHandler = ldap_handler.NewLdapHandlerImpl
	heimdall.NewHeimdall = heimdall.NewHeimdallImpl

	if suite.tempFile != nil {
		os.Remove(suite.tempFile.Name())
	}
}

func (suite *ConfigureServerTestSuite) TearDownSuite() {

}

func (suite *ConfigureServerTestSuite) TestConfigureAPI() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().CheckUpgrade().Once().Return(false, nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(nil)

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), db)
	require.NotNil(suite.T(), ldap)
	require.NotNil(suite.T(), heimdall)

	require.Equal(suite.T(), suite.mockDatabase, db)
	require.Equal(suite.T(), suite.mockLdap, ldap)
	require.Equal(suite.T(), suite.mockHeimdall, heimdall)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIDropError() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.ErrorContains(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

}

type _Server struct {
}

func (suite *ConfigureServerTestSuite) TestConfigureAPIReadConfigENVOk() {
	os.Setenv("NOODLE_DEBUG", "true")

	os.Setenv("NOODLE_ICON_SAVE_PATH", "/disk/web-client/icons")
	os.Setenv("NOODLE_WEB_CLIENT_PATH", "/disk/web-client")
	os.Setenv("NOODLE_HEIMDALL_LIST_JSON_URL", "http://site/list.json")
	os.Setenv("NOODLE_HEIMDALL_ICON_BASE_URL", "http://site/icons")

	os.Setenv("NOODLE_POSTGRES_USER", "postgresuser")
	os.Setenv("NOODLE_POSTGRES_PASSWORD", "postgrespass")
	os.Setenv("NOODLE_POSTGRES_DB", "postgres")
	os.Setenv("NOODLE_POSTGRES_PORT", "5432")
	os.Setenv("NOODLE_POSTGRES_HOSTNAME", "localhost")

	os.Setenv("NOODLE_LDAP_URL", "ldap://example.com")
	os.Setenv("NOODLE_LDAP_BASE_DN", "dc=example,dc=com")
	os.Setenv("NOODLE_LDAP_USER", "CN=readonly,DC=example,DC=com")
	os.Setenv("NOODLE_LDAP_PASSWORD", "readonly")

	os.Setenv("NOODLE_LDAP_USER_FILTER", "(&(objectClass=organizationalPerson)(uid=%s))")
	os.Setenv("NOODLE_LDAP_ALL_USERS_FILTER", "(objectclass=organizationalPerson)")
	os.Setenv("NOODLE_LDAP_ALL_GROUPS_FILTER", "(objectclass=groupOfUniqueNames)")
	os.Setenv("NOODLE_LDAP_USER_GROUPS_FILTER", "(&(uniquemember={dn})(objectclass=groupOfUniqueNames))")
	os.Setenv("NOODLE_LDAP_GROUP_USERS_FILTER", "(&(objectClass=groupOfUniqueNames)(cn=%s))")
	os.Setenv("NOODLE_LDAP_USERNAME_ATTRIBUTE", "uid")
	os.Setenv("NOODLE_LDAP_GROUP_NAME_ATTRIBUTE", "cn")
	os.Setenv("NOODLE_LDAP_USER_DISPLAY_NAME_ATTRIBUTE", "displayName")
	os.Setenv("NOODLE_LDAP_GROUP_MEMBER_ATTRIBUTE", "uniqueMember")

	api := &operations.NoodleAPI{}
	serv := _Server{}

	parser := flags.NewParser(&serv, flags.IgnoreUnknown)
	parser.ShortDescription = "Noodle"
	parser.LongDescription = "Noodle"

	configure_server.NewConfigureServer().ConfigureFlags(api)

	for _, optsGroup := range api.CommandLineOptionsGroups {
		parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
	}

	_, err := parser.Parse()

	require.NoError(suite.T(), err)

	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().CheckUpgrade().Once().Return(false, nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(nil)

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), db)
	require.NotNil(suite.T(), ldap)
	require.NotNil(suite.T(), heimdall)

	require.Equal(suite.T(), suite.mockDatabase, db)
	require.Equal(suite.T(), suite.mockLdap, ldap)
	require.Equal(suite.T(), suite.mockHeimdall, heimdall)

	require.True(suite.T(), noodleOptions.Debug)
	require.True(suite.T(), noodleOptions.Drop)

	require.Equal(suite.T(), "http://site/list.json", noodleOptions.HeimdallListJsonURL)
	require.Equal(suite.T(), "http://site/icons", noodleOptions.HeimdallIconBaseURL)
	require.Equal(suite.T(), "/disk/web-client/icons", noodleOptions.IconSavePath)
	require.Equal(suite.T(), "/disk/web-client", noodleOptions.WebClientPath)

	postgresOptions := api.CommandLineOptionsGroups[1].Options.(*options.PostgresOptions)
	require.Equal(suite.T(), "postgresuser", postgresOptions.User)
	require.Equal(suite.T(), "postgrespass", postgresOptions.Password)
	require.Equal(suite.T(), "postgres", postgresOptions.Database)
	require.Equal(suite.T(), "localhost", postgresOptions.Hostname)
	require.Equal(suite.T(), 5432, postgresOptions.Port)

	ldapOptions := api.CommandLineOptionsGroups[2].Options.(*options.LDAPOptions)
	require.Equal(suite.T(), "ldap://example.com", ldapOptions.URL)
	require.Equal(suite.T(), "dc=example,dc=com", ldapOptions.BaseDN)
	require.Equal(suite.T(), "uid", ldapOptions.UserNameAttribute)
	require.Equal(suite.T(), "(&(objectClass=organizationalPerson)(uid=%s))", ldapOptions.UserFilter)
	require.Equal(suite.T(), "(objectclass=organizationalPerson)", ldapOptions.AllUsersFilter)
	require.Equal(suite.T(), "(objectclass=groupOfUniqueNames)", ldapOptions.AllGroupsFilter)
	require.Equal(suite.T(), "(&(uniquemember={dn})(objectclass=groupOfUniqueNames))", ldapOptions.UserGroupsFilter)
	require.Equal(suite.T(), "(&(objectClass=groupOfUniqueNames)(cn=%s))", ldapOptions.GroupUsersFilter)
	require.Equal(suite.T(), "cn", ldapOptions.GroupNameAttribute)
	require.Equal(suite.T(), "displayName", ldapOptions.UserDisplayNameAttribute)
	require.Equal(suite.T(), "CN=readonly,DC=example,DC=com", ldapOptions.User)
	require.Equal(suite.T(), "readonly", ldapOptions.Password)
	require.Equal(suite.T(), "uniqueMember", ldapOptions.GroupMemberAttribute)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIErrorDBConnectFailed() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Connect().Once().Return(errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.NotNil(suite.T(), err)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

	require.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "failed"
	}, time.Second, time.Millisecond*500)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPILDAPError() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().CheckUpgrade().Once().Return(false, nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.NotNil(suite.T(), err)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

	require.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "failed"
	}, time.Second, time.Millisecond*500)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIErrorDBNotCreatedFailedCreate() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(false, nil)
	suite.mockDatabase.EXPECT().Create().Once().Return(errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.NotNil(suite.T(), err)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

	require.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "failed"
	}, time.Second, time.Millisecond*500)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIErrorDBCheckUpgradeFailed() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().CheckUpgrade().Once().Return(false, errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.NotNil(suite.T(), err)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

	require.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "failed"
	}, time.Second, time.Millisecond*500)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIErrorDBUpgradeFailed() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().CheckUpgrade().Once().Return(true, nil)
	suite.mockDatabase.EXPECT().Upgrade().Once().Return(errors.New("failed"))

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.NotNil(suite.T(), err)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)

	require.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "failed"
	}, time.Second, time.Millisecond*500)

}

func (suite *ConfigureServerTestSuite) TestConfigureAPIPreloadDBOK() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(false, nil)
	suite.mockDatabase.EXPECT().Create().Once().Return(nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(nil)
	suite.mockHeimdall.EXPECT().UpdateFromServer().Once().Return(nil)

	api_handlers.LDAPRefresh = func(db database.Database, ldap ldap_handler.LdapHandler) error { return nil }

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.Nil(suite.T(), err)

	require.NotNil(suite.T(), db)
	require.NotNil(suite.T(), ldap)
	require.NotNil(suite.T(), heimdall)
}

func (suite *ConfigureServerTestSuite) TestConfigureAPIPreloadDBLdapFailed() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(false, nil)
	suite.mockDatabase.EXPECT().Create().Once().Return(nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(nil)
	// suite.mockHeimdall.EXPECT().UpdateFromServer().Once().Return(nil)

	api_handlers.LDAPRefresh = func(db database.Database, ldap ldap_handler.LdapHandler) error { return errors.New("failed") }

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.ErrorContains(suite.T(), err, "failed")

	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)
}

func (suite *ConfigureServerTestSuite) TestConfigureAPIPreloadDBHeimdallFailed() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)
	noodleOptions := api.CommandLineOptionsGroups[0].Options.(*options.NoodleOptions)
	noodleOptions.Debug = true
	noodleOptions.Drop = true

	suite.mockDatabase.EXPECT().Tables().Once().Return(suite.mockTables)
	suite.mockDatabase.EXPECT().Connect().Once().Return(nil)
	suite.mockTables.EXPECT().InitTables(suite.mockDatabase).Once()
	suite.mockDatabase.EXPECT().Drop().Once().Return(nil)
	suite.mockDatabase.EXPECT().CheckCreated().Once().Return(false, nil)
	suite.mockDatabase.EXPECT().Create().Once().Return(nil)

	suite.mockLdap.EXPECT().Connect().Once().Return(nil)
	suite.mockHeimdall.EXPECT().UpdateFromServer().Once().Return(errors.New("failed"))

	api_handlers.LDAPRefresh = func(db database.Database, ldap ldap_handler.LdapHandler) error { return nil }

	db, ldap, heimdall, err := configure_server.NewConfigureServer().ConfigureAPI(api)
	require.ErrorContains(suite.T(), err, "failed")

	require.Nil(suite.T(), db)
	require.Nil(suite.T(), ldap)
	require.Nil(suite.T(), heimdall)
}

func (suite *ConfigureServerTestSuite) TestConfigureFlags() {
	api := &operations.NoodleAPI{}

	configure_server.NewConfigureServer().ConfigureFlags(api)

	require.Equal(suite.T(), 3, len(api.CommandLineOptionsGroups))

	require.Equal(suite.T(), "config", api.CommandLineOptionsGroups[0].ShortDescription)
	require.Equal(suite.T(), "Noodle Config", api.CommandLineOptionsGroups[0].LongDescription)

	require.Equal(suite.T(), "PostgreSQL", api.CommandLineOptionsGroups[1].ShortDescription)
	require.Equal(suite.T(), "PostgreSQL Options", api.CommandLineOptionsGroups[1].LongDescription)

	require.Equal(suite.T(), "LDAP", api.CommandLineOptionsGroups[2].ShortDescription)
	require.Equal(suite.T(), "LDAP Options", api.CommandLineOptionsGroups[2].LongDescription)

}

func TestAppTemplateHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigureServerTestSuite))
}
