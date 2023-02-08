package ldap_handler_test

import (
	"crypto/tls"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/jessevdk/go-flags"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/mwinters-stuff/noodle/package-shims/ldap/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
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
	ldapOptions options.LDAPOptions
	mockLdap    *mocks.LdapShim
	ldapHandler ldap_handler.LdapHandler
}

func (suite *LdapHandlerTestSuite) SetupSuite() {
	suite.loghook = ldapHandlerLogHook{}
	ldap_handler.Logger = log.Hook(&suite.loghook).Output(nil)

	os.Setenv("NOODLE_LDAP_URL", "ldap://example.nz")
	os.Setenv("NOODLE_LDAP_BASE_DN", "DC=example,DC=nz")
	os.Setenv("NOODLE_LDAP_USER", "CN=readonly,DC=example,DC=nz")
	os.Setenv("NOODLE_LDAP_PASSWORD", "readonly")

	os.Setenv("NOODLE_LDAP_USER_FILTER", "(&(objectClass=organizationalPerson)(uid=%s))")
	os.Setenv("NOODLE_LDAP_ALL_USERS_FILTER", "(objectclass=organizationalPerson)")
	os.Setenv("NOODLE_LDAP_ALL_GROUPS_FILTER", "(objectclass=groupOfUniqueNames)")
	os.Setenv("NOODLE_LDAP_USER_GROUPS_FILTER", "(&(uniquemember=%s)(objectclass=groupOfUniqueNames))")
	os.Setenv("NOODLE_LDAP_GROUP_USERS_FILTER", "(&(objectClass=groupOfUniqueNames)(cn=%s))")
	os.Setenv("NOODLE_LDAP_USERNAME_ATTRIBUTE", "uid")
	os.Setenv("NOODLE_LDAP_GROUP_NAME_ATTRIBUTE", "cn")
	os.Setenv("NOODLE_LDAP_USER_DISPLAY_NAME_ATTRIBUTE", "displayName")
	os.Setenv("NOODLE_LDAP_GROUP_MEMBER_ATTRIBUTE", "uniqueMember")

	parser := flags.NewParser(&suite.ldapOptions, flags.IgnoreUnknown)
	_, err := parser.Parse()

	require.NoError(suite.T(), err)

}

func (suite *LdapHandlerTestSuite) SetupTest() {
	suite.mockLdap = mocks.NewLdapShim(suite.T())
	suite.ldapHandler = ldap_handler.NewLdapHandler(suite.mockLdap, suite.ldapOptions)
}

func (suite *LdapHandlerTestSuite) TearDownTest() {
}

func (suite *LdapHandlerTestSuite) TestConnectFailedDial() {

	suite.mockLdap.EXPECT().DialURL("ldap://example.nz").Return(errors.New("it failed"))

	err := suite.ldapHandler.Connect()
	require.EqualError(suite.T(), err, "it failed")

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)

}

func (suite *LdapHandlerTestSuite) TestConnectFailedStartTLS() {

	suite.mockLdap.EXPECT().DialURL("ldap://example.nz").Return(nil)
	suite.mockLdap.EXPECT().StartTLS(&tls.Config{InsecureSkipVerify: true}).Return(errors.New("it failed"))
	suite.mockLdap.EXPECT().CloseConn().Return()

	err := suite.ldapHandler.Connect()
	require.EqualError(suite.T(), err, "it failed")

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "StartTLS", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "CloseConn", 1)
}

func (suite *LdapHandlerTestSuite) TestConnectFailedBind() {
	suite.mockLdap.EXPECT().DialURL("ldap://example.nz").Return(nil)
	suite.mockLdap.EXPECT().StartTLS(&tls.Config{InsecureSkipVerify: true}).Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=nz", "readonly").Return(errors.New("it failed"))
	suite.mockLdap.EXPECT().CloseConn().Return()

	err := suite.ldapHandler.Connect()
	require.EqualError(suite.T(), err, "it failed")

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "StartTLS", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "CloseConn", 1)

}

func (suite *LdapHandlerTestSuite) TestConnectSuccess() {
	suite.mockLdap.EXPECT().DialURL("ldap://example.nz").Return(nil)
	suite.mockLdap.EXPECT().StartTLS(&tls.Config{InsecureSkipVerify: true}).Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=nz", "readonly").Return(nil)
	// suite.mockLdap.EXPECT().CloseConn().Return()

	err := suite.ldapHandler.Connect()
	require.NoError(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "Connected to LDAP Server"
	}, time.Second*3, time.Millisecond*100)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "DialURL", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "StartTLS", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 1)
	// suite.mockLdap.AssertNumberOfCalls(suite.T(), "CloseConn", 1)

}

func (suite *LdapHandlerTestSuite) TestAuthUserSuccess() {
	suite.mockLdap.EXPECT().Bind("CN=bob,DC=example,DC=nz", "pass").Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=nz", "readonly").Return(nil)

	success, err := suite.ldapHandler.AuthUser("CN=bob,DC=example,DC=nz", "pass")
	require.NoError(suite.T(), err)
	require.True(suite.T(), success)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 2)

}

func (suite *LdapHandlerTestSuite) TestAuthUserFail() {
	suite.mockLdap.EXPECT().Bind("CN=bob,DC=example,DC=nz", "pass").Return(errors.New("Bad Auth"))
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=nz", "readonly").Return(nil)

	success, err := suite.ldapHandler.AuthUser("CN=bob,DC=example,DC=nz", "pass")
	require.EqualError(suite.T(), err, "Bad Auth")
	require.False(suite.T(), success)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 2)
}

func (suite *LdapHandlerTestSuite) TestAuthUserReauthError() {
	suite.mockLdap.EXPECT().Bind("CN=bob,DC=example,DC=nz", "pass").Return(nil)
	suite.mockLdap.EXPECT().Bind("CN=readonly,DC=example,DC=nz", "readonly").Return(errors.New("Bad Auth"))

	success, err := suite.ldapHandler.AuthUser("CN=bob,DC=example,DC=nz", "pass")
	require.EqualError(suite.T(), err, "Bad Auth")
	require.False(suite.T(), success)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Bind", 2)
}

func (suite *LdapHandlerTestSuite) TestGetUserByDN() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"CN=bob,DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("CN=bob,DC=example,DC=nz", map[string][]string{
				"displayName": {"bobextample"},
				"givenName":   {"Bob"},
				"sn":          {"Extample"},
				"uidNumber":   {"1001"},
				"uid":         {"bobe"},
			})}}, nil)

	user, err := suite.ldapHandler.GetUserByDN("CN=bob,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), user)

	require.Equal(suite.T(), models.User{
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserByDNLDAPError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"CN=bob,DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	user, err := suite.ldapHandler.GetUserByDN("CN=bob,DC=example,DC=nz")
	require.EqualError(suite.T(), err, "failed")
	require.Equal(suite.T(), models.User{}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserByDNErrorMoreThanOne() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"CN=bob,DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "CN=bob,DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("CN=bob,DC=example,DC=nz", map[string][]string{
				"displayName": {"bobextample"},
				"givenName":   {"Bob"},
				"sn":          {"Extample"},
				"uidNumber":   {"1001"},
				"uid":         {"bobe"},
			}),
			ldap.NewEntry("CN=bob,DC=example,DC=nz", map[string][]string{
				"displayName": {"bobextample"},
				"givenName":   {"Bob"},
				"sn":          {"Extample"},
				"uidNumber":   {"1001"},
				"uid":         {"bobe"},
			})}}, nil)

	user, err := suite.ldapHandler.GetUserByDN("CN=bob,DC=example,DC=nz")
	require.EqualError(suite.T(), err, "user CN=bob,DC=example,DC=nz does not exist or too many entries returned")
	require.Equal(suite.T(), models.User{}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) userByDnSteps(dn string, username string, displayname string, surname string, givenname string, uidnumber int, err error) {
	suite.mockLdap.EXPECT().NewSearchRequest(
		dn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       dn,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       dn,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry(dn, map[string][]string{
				"displayName": {displayname},
				"givenName":   {givenname},
				"sn":          {surname},
				"uidNumber":   {strconv.Itoa(uidnumber)},
				"uid":         {username},
			})}}, err)
}

func (suite *LdapHandlerTestSuite) TestGetGroupUsers() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"cn=admins,ou=groups,dc=example,dc=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "uniqueMember"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("cn=admins,ou=groups,dc=example,dc=nz", map[string][]string{
				"uniqueMember": {"uid=testuser1,ou=people,dc=example,dc=nz",
					"uid=testuser2,ou=people,dc=example,dc=nz",
				},
			})}}, nil)

	suite.userByDnSteps("uid=testuser1,ou=people,dc=example,dc=nz", "TestUser1", "TestUser1", "test", "user1", 1001, nil)
	suite.userByDnSteps("uid=testuser2,ou=people,dc=example,dc=nz", "TestUser2", "TestUser2", "test", "user2", 1002, nil)

	group := models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}

	usergroups, err := suite.ldapHandler.GetGroupUsers(group)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), usergroups)

	require.ElementsMatch(suite.T(), []models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			UserDN:    "uid=testuser1,ou=people,dc=example,dc=nz",
			GroupName: "Admins",
			UserName:  "TestUser1",
		},
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			UserDN:    "uid=testuser2,ou=people,dc=example,dc=nz",
			GroupName: "Admins",
			UserName:  "TestUser2",
		},
	}, usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 3)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 3)
}

func (suite *LdapHandlerTestSuite) TestGetGroupUsersLDAPSearchError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"cn=admins,ou=groups,dc=example,dc=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "uniqueMember"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	// suite.userByDnSteps("uid=testuser1,ou=people,dc=example,dc=nz", "TestUser1", "TestUser1", "test", "user1", 1001)
	// suite.userByDnSteps("uid=testuser2,ou=people,dc=example,dc=nz", "TestUser2", "TestUser2", "test", "user2", 1002)

	group := models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}

	usergroups, err := suite.ldapHandler.GetGroupUsers(group)
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetGroupUsersErrorNoResult() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"cn=admins,ou=groups,dc=example,dc=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "uniqueMember"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, nil)

	// suite.userByDnSteps("uid=testuser1,ou=people,dc=example,dc=nz", "TestUser1", "TestUser1", "test", "user1", 1001,errors.New("user failed"))
	// suite.userByDnSteps("uid=testuser2,ou=people,dc=example,dc=nz", "TestUser2", "TestUser2", "test", "user2", 1002)

	group := models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}

	usergroups, err := suite.ldapHandler.GetGroupUsers(group)
	require.EqualError(suite.T(), err, "group cn=admins,ou=groups,dc=example,dc=nz does not exist or too many entries returned")
	require.Nil(suite.T(), usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetGroupUsersErrorUserSearchFailed() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"cn=admins,ou=groups,dc=example,dc=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "uniqueMember"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "cn=admins,ou=groups,dc=example,dc=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "uniqueMember"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("cn=admins,ou=groups,dc=example,dc=nz", map[string][]string{
				"uniqueMember": {"uid=testuser1,ou=people,dc=example,dc=nz",
					"uid=testuser2,ou=people,dc=example,dc=nz",
				},
			})}}, nil)

	suite.userByDnSteps("uid=testuser1,ou=people,dc=example,dc=nz", "TestUser1", "TestUser1", "test", "user1", 1001, errors.New("user failed"))
	// suite.userByDnSteps("uid=testuser2,ou=people,dc=example,dc=nz", "TestUser2", "TestUser2", "test", "user2", 1002)

	group := models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}

	usergroups, err := suite.ldapHandler.GetGroupUsers(group)
	require.EqualError(suite.T(), err, "user failed")
	require.Nil(suite.T(), usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 2)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 2)
}

func (suite *LdapHandlerTestSuite) TestGetGroups() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "cn"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("cn=admins,ou=groups,dc=example,dc=nz", map[string][]string{
				"cn": {"Admins"},
			}),
			ldap.NewEntry("cn=users,ou=groups,dc=example,dc=nz", map[string][]string{
				"cn": {"Users"},
			}),
		}}, nil)

	groups, err := suite.ldapHandler.GetGroups()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), groups)

	require.ElementsMatch(suite.T(), []models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
		{
			DN:   "cn=users,ou=groups,dc=example,dc=nz",
			Name: "Users",
		},
	}, groups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetGroupsError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=groupOfUniqueNames)",
		[]string{"dn", "cn"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=groupOfUniqueNames)",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	groups, err := suite.ldapHandler.GetGroups()
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), groups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUser() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson)(uid=bob))",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("CN=bob,DC=example,DC=nz", map[string][]string{
				"displayName": {"bobextample"},
				"givenName":   {"Bob"},
				"sn":          {"Extample"},
				"uidNumber":   {"1001"},
				"uid":         {"bobe"},
			})}}, nil)

	user, err := suite.ldapHandler.GetUser("bob")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), user)

	require.Equal(suite.T(), models.User{
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson)(uid=bob))",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	user, err := suite.ldapHandler.GetUser("bob")
	require.EqualError(suite.T(), err, "failed")
	require.Equal(suite.T(), models.User{}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserErrorNone() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson)(uid=bob))",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(objectClass=organizationalPerson)(uid=bob))",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, nil)

	user, err := suite.ldapHandler.GetUser("bob")
	require.EqualError(suite.T(), err, "user bob does not exist or too many entries returned")
	require.Equal(suite.T(), models.User{}, user)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserGroups() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		[]string{"dn", "cn"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("cn=admins,ou=groups,dc=example,dc=nz", map[string][]string{
				"cn": {"Admins"},
			},
			),
			ldap.NewEntry("cn=users,ou=groups,dc=example,dc=nz", map[string][]string{
				"cn": {"Users"},
			},
			)}}, nil)

	usergroups, err := suite.ldapHandler.GetUserGroups(models.User{DN: "uid=testuser1,ou=people,dc=example,dc=nz", DisplayName: "TestUser1"})
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), usergroups)

	require.ElementsMatch(suite.T(), []models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			UserDN:    "uid=testuser1,ou=people,dc=example,dc=nz",
			GroupName: "Admins",
			UserName:  "TestUser1",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			UserDN:    "uid=testuser1,ou=people,dc=example,dc=nz",
			GroupName: "Users",
			UserName:  "TestUser1",
		},
	}, usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUserGroupsError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		[]string{"dn", "cn"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(&(uniquemember=uid=testuser1,ou=people,dc=example,dc=nz)(objectclass=groupOfUniqueNames))",
		Attributes:   []string{"dn", "cn"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	usergroups, err := suite.ldapHandler.GetUserGroups(models.User{DN: "uid=testuser1,ou=people,dc=example,dc=nz", DisplayName: "TestUser1"})
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), usergroups)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUsers() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{
			ldap.NewEntry("CN=bob,DC=example,DC=nz", map[string][]string{
				"displayName": {"bobextample"},
				"givenName":   {"Bob"},
				"sn":          {"Extample"},
				"uidNumber":   {"1001"},
				"uid":         {"bobe"},
			}),
			ldap.NewEntry("CN=jill,DC=example,DC=nz", map[string][]string{
				"displayName": {"jilly"},
				"givenName":   {"Jill"},
				"sn":          {"Frill"},
				"uidNumber":   {"1002"},
				"uid":         {"jillie"},
			}),
		}}, nil)

	users, err := suite.ldapHandler.GetUsers()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), users)

	require.ElementsMatch(suite.T(), []models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
		{
			DN:          "CN=jill,DC=example,DC=nz",
			Username:    "jillie",
			DisplayName: "jilly",
			Surname:     "Frill",
			GivenName:   "Jill",
			UIDNumber:   1002,
		},
	}, users)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func (suite *LdapHandlerTestSuite) TestGetUsersError() {
	suite.mockLdap.EXPECT().NewSearchRequest(
		"DC=example,DC=nz",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectclass=organizationalPerson)",
		[]string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		[]ldap.Control(nil),
	).Return(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	})

	suite.mockLdap.EXPECT().Search(&ldap.SearchRequest{
		BaseDN:       "DC=example,DC=nz",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       "(objectclass=organizationalPerson)",
		Attributes:   []string{"dn", "displayName", "givenName", "sn", "uidNumber", "uid"},
		Controls:     nil,
	}).Return(&ldap.SearchResult{
		Entries: []*ldap.Entry{}}, errors.New("failed"))

	users, err := suite.ldapHandler.GetUsers()
	require.EqualError(suite.T(), err, "failed")
	require.Nil(suite.T(), users)

	suite.mockLdap.AssertNumberOfCalls(suite.T(), "NewSearchRequest", 1)
	suite.mockLdap.AssertNumberOfCalls(suite.T(), "Search", 1)
}

func TestLDAPHandlerSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlerTestSuite))
}
