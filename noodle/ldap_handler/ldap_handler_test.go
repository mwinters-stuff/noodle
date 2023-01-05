package ldap_handler_test

import (
	"crypto/tls"
	"errors"
	"testing"
	"time"

	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/mwinters-stuff/noodle/package-shims/ldap/mocks"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ldapHandlerLogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *ldapHandlerLogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type LdapHandlerTestSuite struct {
	suite.Suite
	loghook     ldapHandlerLogHook
	appConfig   yamltypes.AppConfig
	mockLdap    *mocks.LdapShim
	ldapHandler ldap_handler.LdapHandler
}

func (suite *LdapHandlerTestSuite) SetupSuite() {
	suite.loghook = ldapHandlerLogHook{}
	ldap_handler.Logger = log.Hook(&suite.loghook)

	yamltext := `
postgres:
  user: postgresuser
  password: postgrespass
  db: postgres
  hostname: localhost
  port: 5432
ldap:
  url: ldap://example.com
  base_dn: dc=example,dc=com
  username_attribute: uid
  user_filter: (&(objectClass=organizationalPerson)(uid=%s))
  all_users_filter: (objectclass=organizationalPerson)
  all_groups_filter: (objectclass=groupOfUniqueNames)
  user_groups_filter: (&(uniquemember={dn})(objectclass=groupOfUniqueNames))
  group_users_filter: (&(objectClass=groupOfUniqueNames)(cn=%s))
  group_name_attribute: cn
  user_display_name_attribute: displayName
  user: CN=readonly,DC=example,DC=com
  password: readonly
  group_member_attribute: uniqueMember
`
	var err error

	suite.appConfig, err = yamltypes.UnmarshalConfig([]byte(yamltext))
	require.NoError(suite.T(), err)

}
func (suite *LdapHandlerTestSuite) SetupTest() {
	suite.mockLdap = mocks.NewLdapShim(suite.T())
	suite.ldapHandler = ldap_handler.NewLdapHandler(suite.mockLdap, suite.appConfig)
}

func (suite *LdapHandlerTestSuite) TearDownTest() {
}

func (suite *LdapHandlerTestSuite) TestConnectFailed() {

	suite.mockLdap.EXPECT().DialURL("ldap://example.com").Return(errors.New("it failed"))

	err := suite.ldapHandler.Connect()
	require.Error(suite.T(), err, "it failed")
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)

}

func (suite *LdapHandlerTestSuite) TestConnectSuccess() {
	suite.mockLdap.EXPECT().DialURL("ldap://example.com").Return(nil)
	suite.mockLdap.EXPECT().StartTLS(&tls.Config{InsecureSkipVerify: true}).Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=com", "readonly").Return(nil)
	suite.mockLdap.EXPECT().CloseConn().Return()

	err := suite.ldapHandler.Connect()
	require.NoError(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "Connected to LDAP Server"
	}, time.Second*3, time.Millisecond*100)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "StartTLS", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "CloseConn", 1)

}

func (suite *LdapHandlerTestSuite) TestAuthUserSuccess() {
	suite.mockLdap.EXPECT().Bind("CN=bob,DC=example,DC=com", "pass").Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=com", "readonly").Return(nil)

	success, err := suite.ldapHandler.AuthUser("CN=bob,DC=example,DC=com", "pass")
	require.NoError(suite.T(), err)
	require.True(suite.T(), success)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 2)

}

func (suite *LdapHandlerTestSuite) TestAuthUserFail() {
	suite.mockLdap.EXPECT().Bind("CN=bob,DC=example,DC=com", "pass").Return(errors.New("Bad Auth"))
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=com", "readonly").Return(nil)

	success, err := suite.ldapHandler.AuthUser("CN=bob,DC=example,DC=com", "pass")
	require.Error(suite.T(), err, "Bad Auth")
	require.False(suite.T(), success)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 2)

}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlerTestSuite))
}

// func TestCheck(t *testing.T) {
// 	// The username and password we want to check
// 	// username := "jessica"
// 	// password := "harperismydog"

// 	bindusername := "cn=readonly,dc=winters,dc=nz"
// 	bindpassword := "readonly"

// 	l, err := ldap.DialURL("ldap://192.168.30.23:389")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer l.Close()

// 	// Reconnect with TLS
// 	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// First bind with a read only user
// 	err = l.Bind(bindusername, bindpassword)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Search for the given username
// 	searchRequest := ldap.NewSearchRequest(
// 		"cn=jenkins,ou=groups,dc=winters,dc=nz",
// 		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
// 		, // ldap.EscapeFilter("cn=jenkins,ou=groups,dc=winters,dc=nz")),
// 		[]string{"dn", "cn", "uniqueMember"},
// 		nil,
// 	)

// 	sr, err := l.Search(searchRequest)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, e := range sr.Entries {
// 		e.PrettyPrint(2)
// 	}

// 	// sr.Entries[0].PrettyPrint(2)
// 	// log.Print(sr.Entries[0], userdn)
// 	// // Bind as the user to verify their password
// 	// err = l.Bind(userdn, password)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// // Rebind as the read only user for any further queries
// 	// err = l.Bind(bindusername, bindpassword)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// searchRequest = ldap.NewSearchRequest(
// 	// 	"dc=winters,dc=nz", // The base dn to search
// 	// 	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
// 	// 	fmt.Sprintf("(&(uniquemember=%s)(objectclass=groupOfUniqueNames))", "*"), // The filter to apply
// 	// 	[]string{"dn", "cn", "uniqueMember"},                                     // A list attributes to retrieve
// 	// 	nil,
// 	// )

// 	// sr, err = l.Search(searchRequest)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// for _, e := range sr.Entries {
// 	// 	e.PrettyPrint(2)
// 	// }

// }
