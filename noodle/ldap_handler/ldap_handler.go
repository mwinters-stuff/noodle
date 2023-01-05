package ldap_handler

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/go-ldap/ldap/v3"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	ldap_shim "github.com/mwinters-stuff/noodle/package-shims/ldap"
)

var (
	NewLdapHandler = NewLdapHandlerImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name LdapHandler

type LdapHandler interface {
	Connect() error
	GetUsers() ([]database.User, error)
	GetGroups() ([]database.Group, error)

	GetUser(username string) (database.User, error)
	GetUserByDN(dn string) (database.User, error)
	GetUserGroups(database.User) ([]database.UserGroup, error)

	GetGroupUsers(database.Group) ([]database.UserGroup, error)

	AuthUser(userdn, password string) (bool, error)
}

type LdapHandlerImpl struct {
	appConfig yamltypes.AppConfig
	ldapShim  ldap_shim.LdapShim
}

// GetUserByDN implements LdapAuth
func (i *LdapHandlerImpl) GetUserByDN(dn string) (database.User, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		dn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		i.appConfig.Ldap.AllUsersFilter,
		[]string{"dn", i.appConfig.Ldap.UserDisplayNameAttribute, "givenName", "sn", "uidNumber", i.appConfig.Ldap.UsernameAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return database.User{}, err
	}

	if len(sr.Entries) != 1 {
		Logger.Error().Msgf("User %s does not exist or too many entries returned", dn)
		return database.User{}, fmt.Errorf("user %s does not exist or too many entries returned", dn)
	}

	uid, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("uidNumber"))
	return database.User{
		Username:    sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
		DN:          sr.Entries[0].DN,
		DisplayName: sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
		GivenName:   sr.Entries[0].GetAttributeValue("givenName"),
		Surname:     sr.Entries[0].GetAttributeValue("sn"),
		UidNumber:   uid,
	}, nil

}

// Conn	ect implements LdapAuth
func (i *LdapHandlerImpl) Connect() error {
	var err error
	err = i.ldapShim.DialURL(i.appConfig.Ldap.URL)
	if err != nil {
		Logger.Error().Err(err)
		return err
	}
	defer i.ldapShim.CloseConn()

	// Reconnect with TLS
	err = i.ldapShim.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		Logger.Error().Err(err)
		return err
	}

	// First bind with a read only user
	err = i.ldapShim.Bind(i.appConfig.Ldap.User, i.appConfig.Ldap.Password)
	if err != nil {
		Logger.Error().Err(err)
		return err
	}

	Logger.Info().Msg("Connected to LDAP Server")
	return err

}

// AuthUser implements LdapAuth
func (i *LdapHandlerImpl) AuthUser(userdn string, password string) (bool, error) {
	err := i.ldapShim.Bind(userdn, password)
	success := err == nil
	if err != nil {
		Logger.Error().Err(err)
	}

	nexterr := i.ldapShim.Bind(i.appConfig.Ldap.User, i.appConfig.Ldap.Password)
	if nexterr != nil {
		Logger.Error().Err(nexterr)
	}
	return success, err
}

// GetGroupUsers implements LdapAuth
func (i *LdapHandlerImpl) GetGroupUsers(group database.Group) ([]database.UserGroup, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		group.DN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		i.appConfig.Ldap.AllGroupsFilter,
		[]string{"dn", i.appConfig.Ldap.GroupMemberAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return nil, err
	}

	if len(sr.Entries) != 1 {
		Logger.Error().Msgf("Group %s does not exist or too many entries returned", group.DN)
		return nil, fmt.Errorf("user %s does not exist or too many entries returned", group.DN)
	}

	users := sr.Entries[0].GetAttributeValues(i.appConfig.Ldap.GroupMemberAttribute)

	var results []database.UserGroup
	for _, e := range users {
		user, err := i.GetUserByDN(e)
		if err != nil {
			Logger.Error().Err(err)
			return nil, err
		}

		results = append(results, database.UserGroup{
			GroupDN:   group.DN,
			GroupName: group.DisplayName,
			UserDN:    e,
			UserName:  user.Username,
		})
	}

	return results, nil
}

// GetGroups implements LdapAuth
func (i *LdapHandlerImpl) GetGroups() ([]database.Group, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		i.appConfig.Ldap.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		i.appConfig.Ldap.AllGroupsFilter,
		[]string{"dn", i.appConfig.Ldap.GroupNameAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return nil, err
	}

	var results []database.Group
	for _, e := range sr.Entries {
		results = append(results, database.Group{
			DN:          e.DN,
			DisplayName: e.GetAttributeValue(i.appConfig.Ldap.GroupNameAttribute),
		})
	}

	return results, nil
}

// GetUser implements LdapAuth
func (i *LdapHandlerImpl) GetUser(username string) (database.User, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		i.appConfig.Ldap.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(i.appConfig.Ldap.UserFilter, ldap.EscapeFilter(username)),
		[]string{"dn", i.appConfig.Ldap.UserDisplayNameAttribute, "givenName", "sn", "uidNumber", i.appConfig.Ldap.UsernameAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return database.User{}, err
	}

	if len(sr.Entries) != 1 {
		Logger.Error().Msgf("User %s does not exist or too many entries returned", username)
		return database.User{}, fmt.Errorf("user %s does not exist or too many entries returned", username)
	}

	uid, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("uidNumber"))
	return database.User{
		Username:    sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
		DN:          sr.Entries[0].DN,
		DisplayName: sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
		GivenName:   sr.Entries[0].GetAttributeValue("givenName"),
		Surname:     sr.Entries[0].GetAttributeValue("sn"),
		UidNumber:   uid,
	}, nil

}

// GetUserGroups implements LdapAuth
func (i *LdapHandlerImpl) GetUserGroups(user database.User) ([]database.UserGroup, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		i.appConfig.Ldap.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(i.appConfig.Ldap.UserGroupsFilter, ldap.EscapeFilter(user.DN)),
		[]string{"dn", i.appConfig.Ldap.GroupNameAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return nil, err
	}

	var results []database.UserGroup
	for _, e := range sr.Entries {

		results = append(results, database.UserGroup{

			UserName:  user.DisplayName,
			UserDN:    user.DN,
			GroupDN:   e.DN,
			GroupName: e.GetAttributeValue(i.appConfig.Ldap.GroupNameAttribute),
		})
	}

	return results, nil

}

// GetUsers implements LdapAuth
func (i *LdapHandlerImpl) GetUsers() ([]database.User, error) {
	searchRequest := i.ldapShim.NewSearchRequest(
		i.appConfig.Ldap.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		i.appConfig.Ldap.AllUsersFilter,
		[]string{"dn", i.appConfig.Ldap.UserDisplayNameAttribute, "givenName", "sn", "uidNumber", i.appConfig.Ldap.UsernameAttribute},
		nil,
	)

	sr, err := i.ldapShim.Search(searchRequest)
	if err != nil {
		Logger.Error().Err(err)
		return nil, err
	}

	var results []database.User
	for _, e := range sr.Entries {
		uid, _ := strconv.Atoi(e.GetAttributeValue("uidNumber"))

		results = append(results, database.User{
			Username:    e.GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
			DN:          e.DN,
			DisplayName: e.GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
			GivenName:   e.GetAttributeValue("givenName"),
			Surname:     e.GetAttributeValue("sn"),
			UidNumber:   uid,
		})
	}

	return results, nil
}

func NewLdapHandlerImpl(ldapShim ldap_shim.LdapShim, appConfig yamltypes.AppConfig) LdapHandler {
	return &LdapHandlerImpl{
		appConfig: appConfig,
		ldapShim:  ldapShim,
	}
}
