package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	RegisterLdapApiHandlers = RegisterLdapApiHandlersImpl
	LDAPRefresh             = LDAPRefreshImpl
)

func SyncLDAPUsers(db database.Database, ldap ldap_handler.LdapHandler) error {
	users, err := ldap.GetUsers()
	if err != nil {
		Logger.Error().Err(err).Msg("ldap.GetUsers")
		return err
	}

	var dbusers []*models.User
	if dbusers, err = db.Tables().UserTable().GetAll(); err != nil {
		Logger.Error().Err(err).Msg("userTable.GetAll")
		return err
	}

	var insertusers []models.User
	for _, user := range users {
		exists, err := db.Tables().UserTable().ExistsDN(user.DN)
		if err != nil {
			Logger.Error().Err(err).Msg("userTable.ExistsDN")
			return err
		}
		if exists {
			i := IndexUser(dbusers, user)
			if i > -1 {
				user.ID = dbusers[i].ID
				dbusers = append(dbusers[:i], dbusers[i+1:]...)
				Logger.Info().Msgf("Updating LDAP User %s", user.DisplayName)

				if err = db.Tables().UserTable().Update(user); err != nil {
					Logger.Error().Err(err).Msg("UserTable.Update")
					return err
				}
			}
		} else {
			insertusers = append(insertusers, user)
		}
	}

	for _, dbuser := range dbusers {
		Logger.Info().Msgf("Deleting Database User %s", dbuser.DisplayName)
		if err = db.Tables().UserTable().Delete(*dbuser); err != nil {
			Logger.Error().Err(err).Msg("UserTable.Delete")
			return err
		}
	}

	for _, user := range insertusers {
		Logger.Info().Msgf("Inserting LDAP User %s", user.DisplayName)
		if err = db.Tables().UserTable().Insert(&user); err != nil {
			Logger.Error().Err(err).Msg("UserTable.Insert")
			return err
		}
	}
	return nil
}

func SyncLDAPGroups(db database.Database, ldap ldap_handler.LdapHandler) error {
	groups, err := ldap.GetGroups()
	if err != nil {
		Logger.Error().Err(err).Msg("ldap.GetGroups")
		return err
	}

	var dbgroups []*models.Group
	if dbgroups, err = db.Tables().GroupTable().GetAll(); err != nil {
		Logger.Error().Err(err).Msg("groupTable.GetAll")
		return err
	}

	var insertgroups []models.Group
	for _, group := range groups {
		exists, err := db.Tables().GroupTable().ExistsDN(group.DN)
		if err != nil {
			Logger.Error().Err(err).Msg("GroupTable.ExistsDN")
			return err
		}
		if exists {
			i := IndexGroup(dbgroups, group)
			if i > -1 {
				group.ID = dbgroups[i].ID
				dbgroups = append(dbgroups[:i], dbgroups[i+1:]...)
				Logger.Info().Msgf("Updating LDAP Group %s", group.Name)
				if err = db.Tables().GroupTable().Update(group); err != nil {
					Logger.Error().Err(err).Msg("GroupTable.Update")
					return err
				}
			}
		} else {
			insertgroups = append(insertgroups, group)
		}
	}

	for _, dbgroup := range dbgroups {
		Logger.Info().Msgf("Deleting Database Group %s", dbgroup.Name)
		if err = db.Tables().GroupTable().Delete(*dbgroup); err != nil {
			Logger.Error().Err(err).Msg("GroupTable.Delete")
			return err
		}
	}

	for _, group := range insertgroups {
		Logger.Info().Msgf("Inserting LDAP Group %s", group.Name)
		if err = db.Tables().GroupTable().Insert(&group); err != nil {
			Logger.Error().Err(err).Msg("GroupTable.Insert")
			return err
		}
	}

	return nil

}

func SyncLDAPUserGroups(db database.Database, ldap ldap_handler.LdapHandler) error {

	dbusers, err := db.Tables().UserTable().GetAll()
	if err != nil {
		Logger.Error().Err(err).Msg("UserTable.GetAll")
		return err
	}

	for _, user := range dbusers {
		usergroups, err := ldap.GetUserGroups(*user)
		if err != nil {
			Logger.Error().Err(err).Msg("ldap.GetUserGroups")
			return err
		}
		dbusergroups, err := db.Tables().UserGroupsTable().GetUser(user.ID)
		if err != nil {
			Logger.Error().Err(err).Msg("UserGroupsTable.GetUser")
			return err
		}

		for _, usergroup := range usergroups {
			usergroup.UserID = user.ID
			usergroup.UserName = user.Username

			group, err := db.Tables().GroupTable().GetDN(usergroup.GroupDN)
			if err != nil {
				Logger.Error().Err(err).Msg("GroupTable.GetDN")
				return err
			}
			usergroup.GroupID = group.ID
			usergroup.GroupName = group.Name

			i := IndexUserGroup(dbusergroups, usergroup)
			if i < 0 {
				Logger.Info().Msgf("Updating User Group Mapping %s => %s", usergroup.UserName, usergroup.GroupName)

				err := db.Tables().UserGroupsTable().Insert(&usergroup)
				if err != nil {
					Logger.Error().Err(err).Msg("userGroupsTable.Insert")
					return err
				}
			} else {
				usergroup.ID = dbusergroups[i].ID
				dbusergroups = append(dbusergroups[:i], dbusergroups[i+1:]...)
			}
		}

		for _, dbusergroup := range dbusergroups {
			Logger.Info().Msgf("Deleting Database User Group Mapping %s => %s", dbusergroup.UserName, dbusergroup.GroupName)
			err := db.Tables().UserGroupsTable().Delete(*dbusergroup)
			if err != nil {
				Logger.Error().Err(err).Msg("UserGroupsTable.delete")
				return err
			}
		}
	}
	return nil
}

func LDAPRefreshImpl(db database.Database, ldap ldap_handler.LdapHandler) error {

	if err := SyncLDAPUsers(db, ldap); err != nil {
		return err
	}

	if err := SyncLDAPGroups(db, ldap); err != nil {
		return err
	}

	return SyncLDAPUserGroups(db, ldap)
}

func HandleLDAPRefresh(db database.Database, ldap ldap_handler.LdapHandler, params noodle_api.GetNoodleLdapReloadParams, principal *models.Principal) middleware.Responder {
	Logger.Info().Msg("Starting LDAP Refresh")

	err := LDAPRefresh(db, ldap)
	if err != nil {
		return noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	Logger.Info().Msg("Finished LDAP Refresh")
	return noodle_api.NewGetNoodleLdapReloadOK()
}

func IndexUserGroup(s []*models.UserGroup, v models.UserGroup) int {
	for i := range s {
		if v.UserID == s[i].UserID && v.GroupID == s[i].GroupID {
			return i
		}
	}
	return -1
}

func IndexUser(s []*models.User, v models.User) int {
	for i := range s {
		if v.DN == s[i].DN {
			return i
		}
	}
	return -1
}

func IndexGroup(s []*models.Group, v models.Group) int {
	for i := range s {
		if v.DN == s[i].DN {
			return i
		}
	}
	return -1
}

func RegisterLdapApiHandlersImpl(api *operations.NoodleAPI, db database.Database, ldap ldap_handler.LdapHandler) {
	api.NoodleAPIGetNoodleLdapReloadHandler = noodle_api.GetNoodleLdapReloadHandlerFunc(func(params noodle_api.GetNoodleLdapReloadParams, principal *models.Principal) middleware.Responder {
		return HandleLDAPRefresh(db, ldap, params, principal)
	})
}
