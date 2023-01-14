package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const userGroupTableCreate = `CREATE TABLE IF NOT EXISTS user_groups (
  id SERIAL PRIMARY KEY,
  groupid INTEGER REFERENCES groups(id) ON DELETE CASCADE,
  userid INTEGER  REFERENCES users(id) ON DELETE CASCADE
)`

const userGroupTableInsertRow = `INSERT INTO user_groups (groupid, userid) VALUES ($1, $2) RETURNING id`
const userGroupTableDrop = `DROP TABLE user_groups`
const userGroupTableDeleteRow = `DELETE FROM user_groups WHERE id = $1`
const userGroupTableQueryRowsGroup = `SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.groupid = $1 AND g.id = ug.groupid AND u.id = ug.userid`
const userGroupTableQueryRowsUser = `SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.userid = $1 AND g.id = ug.groupid AND u.id = ug.userid`
const userGroupTableQueryAll = `SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE g.id = ug.groupid AND u.id = ug.userid ORDER BY g.name`
const userGroupTableQueryExists = `SELECT COUNT(*) FROM user_groups WHERE groupid = $1 AND userid = $2`

var (
	NewUserGroupsTable = NewUserGroupsTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name UserGroupsTable
type UserGroupsTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(user *models.UserGroup) error
	Delete(user models.UserGroup) error

	GetGroup(groupid int64) ([]models.UserGroup, error)
	GetUser(userid int64) ([]models.UserGroup, error)
	GetAll() ([]models.UserGroup, error)
	Exists(groupid, userid int64) (bool, error)
}

type UserGroupsTableImpl struct {
	database Database
	cache    TableCache[models.UserGroup]
}

// Create implements UserGroupsTable
func (i *UserGroupsTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), userGroupTableCreate)
	return err
}

// Delete implements UserGroupsTable
func (i *UserGroupsTableImpl) Delete(usergroup models.UserGroup) error {
	i.cache.DeleteIndex(usergroup.ID)
	_, err := i.database.Pool().Exec(context.Background(), userGroupTableDeleteRow, usergroup.ID)
	return err

}

// Drop implements UserGroupsTable
func (i *UserGroupsTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), userGroupTableDrop)
	return err
}

// Exists implements UserGroupsTable
func (i *UserGroupsTableImpl) Exists(groupid int64, userid int64) (bool, error) {
	ok, _ := i.cache.Find((func(index int64, value models.UserGroup) bool {
		return value.GroupID == groupid && value.UserID == userid
	}))
	if ok {
		return true, nil
	}

	var found int
	err := i.database.Pool().QueryRow(context.Background(), userGroupTableQueryExists, groupid, userid).Scan(&found)
	return found > 0, err

}

func (i *UserGroupsTableImpl) getQuery(query string, value any) ([]models.UserGroup, error) {
	var rows pgx.Rows
	var err error
	if value == nil {
		rows, err = i.database.Pool().Query(context.Background(), query)
	} else {
		rows, err = i.database.Pool().Query(context.Background(), query, value)
	}
	if err != nil {
		return nil, err
	}
	results := []models.UserGroup{}

	var id, groupid, userid int64
	var groupdn, groupname, userdn, username string
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&groupid,
		&userid,
		&groupdn,
		&groupname,
		&userdn,
		&username,
	}, func() error {

		usergroup := models.UserGroup{
			ID:        id,
			GroupID:   groupid,
			UserID:    userid,
			GroupDN:   groupdn,
			GroupName: groupname,
			UserDN:    userdn,
			UserName:  username,
		}
		results = append(results, usergroup)
		i.cache.Add(id, usergroup)
		return nil
	})

	return results, err
}

// GetAll implements UserGroupsTable
func (i *UserGroupsTableImpl) GetAll() ([]models.UserGroup, error) {
	return i.getQuery(userGroupTableQueryAll, nil)
}

// GetGroup implements UserGroupsTable
func (i *UserGroupsTableImpl) GetGroup(groupid int64) ([]models.UserGroup, error) {
	ok, usergroups := i.cache.FindAll((func(index int64, value models.UserGroup) bool {
		return value.GroupID == groupid
	}))
	if ok {
		return usergroups, nil
	}

	return i.getQuery(userGroupTableQueryRowsGroup, groupid)
}

// GetUser implements UserGroupsTable
func (i *UserGroupsTableImpl) GetUser(userid int64) ([]models.UserGroup, error) {
	ok, usergroups := i.cache.FindAll((func(index int64, value models.UserGroup) bool {
		return value.UserID == userid
	}))
	if ok {
		return usergroups, nil
	}

	return i.getQuery(userGroupTableQueryRowsUser, userid)
}

// Insert implements UserGroupsTable
func (i *UserGroupsTableImpl) Insert(usergroup *models.UserGroup) error {
	err := i.database.Pool().QueryRow(context.Background(), userGroupTableInsertRow,
		usergroup.GroupID,
		usergroup.UserID,
	).Scan(&usergroup.ID)
	i.cache.Add(usergroup.ID, *usergroup)
	return err
}

// Upgrade implements UserGroupsTable
func (*UserGroupsTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewUserGroupsTableImpl(database Database, cache TableCache[models.UserGroup]) UserGroupsTable {
	return &UserGroupsTableImpl{
		database: database,
		cache:    cache,
	}
}
