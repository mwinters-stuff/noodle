package ldap_handler

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/go-ldap/ldap/v3"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	ldap_shim "github.com/mwinters-stuff/noodle/package-shims/ldap"
	"github.com/mwinters-stuff/noodle/server/models"
)

var (
	NewLdapHandler = NewLdapHandlerImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name LdapHandler

type LdapHandler interface {
	Connect() error
	GetUsers() ([]models.User, error)
	GetGroups() ([]models.Group, error)

	GetUser(username string) (models.User, error)
	GetUserByDN(dn string) (models.User, error)
	GetUserGroups(models.User) ([]models.UserGroup, error)

	GetGroupUsers(models.Group) ([]models.UserGroup, error)

	AuthUser(userdn, password string) (bool, error)
}

type LdapHandlerImpl struct {
	appConfig yamltypes.AppConfig
	ldapShim  ldap_shim.LdapShim
}

// GetUserByDN implements LdapAuth
func (i *LdapHandlerImpl) GetUserByDN(dn string) (models.User, error) {
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
		return models.User{}, err
	}

	if len(sr.Entries) != 1 {
		Logger.Error().Msgf("User %s does not exist or too many entries returned", dn)
		return models.User{}, fmt.Errorf("user %s does not exist or too many entries returned", dn)
	}

	uid, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("uidNumber"))
	return models.User{
		Username:    sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
		DN:          sr.Entries[0].DN,
		DisplayName: sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
		GivenName:   sr.Entries[0].GetAttributeValue("givenName"),
		Surname:     sr.Entries[0].GetAttributeValue("sn"),
		UIDNumber:   int64(uid),
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

	// Reconnect with TLS
	err = i.ldapShim.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		i.ldapShim.CloseConn()
		Logger.Error().Err(err)
		return err
	}

	// First bind with a read only user
	err = i.ldapShim.Bind(i.appConfig.Ldap.User, i.appConfig.Ldap.Password)
	if err != nil {
		i.ldapShim.CloseConn()
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
		return false, nexterr
	}
	return success, err
}

// GetGroupUsers implements LdapAuth
func (i *LdapHandlerImpl) GetGroupUsers(group models.Group) ([]models.UserGroup, error) {
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
		return nil, fmt.Errorf("group %s does not exist or too many entries returned", group.DN)
	}

	users := sr.Entries[0].GetAttributeValues(i.appConfig.Ldap.GroupMemberAttribute)

	var results []models.UserGroup
	for _, e := range users {
		user, err := i.GetUserByDN(e)
		if err != nil {
			Logger.Error().Err(err)
			return nil, err
		}

		results = append(results, models.UserGroup{
			GroupDN:   group.DN,
			GroupName: group.Name,
			UserDN:    e,
			UserName:  user.Username,
		})
	}

	return results, nil
}

// GetGroups implements LdapAuth
func (i *LdapHandlerImpl) GetGroups() ([]models.Group, error) {
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

	var results []models.Group
	for _, e := range sr.Entries {
		results = append(results, models.Group{
			DN:   e.DN,
			Name: e.GetAttributeValue(i.appConfig.Ldap.GroupNameAttribute),
		})
	}

	return results, nil
}

// GetUser implements LdapAuth
func (i *LdapHandlerImpl) GetUser(username string) (models.User, error) {
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
		return models.User{}, err
	}

	if len(sr.Entries) != 1 {
		Logger.Error().Msgf("User %s does not exist or too many entries returned", username)
		return models.User{}, fmt.Errorf("user %s does not exist or too many entries returned", username)
	}

	uid, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("uidNumber"))
	return models.User{
		Username:    sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
		DN:          sr.Entries[0].DN,
		DisplayName: sr.Entries[0].GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
		GivenName:   sr.Entries[0].GetAttributeValue("givenName"),
		Surname:     sr.Entries[0].GetAttributeValue("sn"),
		UIDNumber:   int64(uid),
	}, nil

}

// GetUserGroups implements LdapAuth
func (i *LdapHandlerImpl) GetUserGroups(user models.User) ([]models.UserGroup, error) {
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

	var results []models.UserGroup
	for _, e := range sr.Entries {

		results = append(results, models.UserGroup{

			UserName:  user.DisplayName,
			UserDN:    user.DN,
			GroupDN:   e.DN,
			GroupName: e.GetAttributeValue(i.appConfig.Ldap.GroupNameAttribute),
		})
	}

	return results, nil

}

// GetUsers implements LdapAuth
func (i *LdapHandlerImpl) GetUsers() ([]models.User, error) {
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

	var results []models.User
	for _, e := range sr.Entries {
		uid, _ := strconv.Atoi(e.GetAttributeValue("uidNumber"))

		results = append(results, models.User{
			Username:    e.GetAttributeValue(i.appConfig.Ldap.UsernameAttribute),
			DN:          e.DN,
			DisplayName: e.GetAttributeValue(i.appConfig.Ldap.UserDisplayNameAttribute),
			GivenName:   e.GetAttributeValue("givenName"),
			Surname:     e.GetAttributeValue("sn"),
			UIDNumber:   int64(uid),
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
