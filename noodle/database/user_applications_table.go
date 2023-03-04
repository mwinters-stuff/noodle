package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const userApplicationsTableCreate = `CREATE TABLE IF NOT EXISTS user_applications (
  id SERIAL PRIMARY KEY,
  userid int REFERENCES users(id) ON DELETE CASCADE,
  applicationid int REFERENCES applications(id) ON DELETE CASCADE
)`

const userApplicationsTableInsertRow = `INSERT INTO user_applications (userid, applicationid) VALUES ($1, $2) RETURNING id`
const userApplicationsTableDrop = `DROP TABLE user_applications`
const userApplicationsTableDeleteRow = `DELETE FROM user_applications WHERE id = $1`
const userApplicationsTableQueryAll = `SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon,app.template_appid FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid`

const userAllowedQuery = `SELECT d.tabid, d.displayorder, a.id as application_id,a.name,a.website,a.license,a.description,a.enhanced,a.tilebackground,a.icon,a.template_appid
FROM applications a ,
(
  SELECT ua.applicationid, at.tabid, at.displayorder FROM user_applications ua, application_tabs at WHERE at.applicationid = ua.applicationid AND userid = $1
  UNION
  SELECT ga.applicationid, at.tabid, at.displayorder FROM user_groups ug, group_applications ga, application_tabs at WHERE at.applicationid = ga.applicationid AND ga.groupid = ug.groupid AND userid = $1
  UNION
  SELECT applicationid, tabid, displayorder FROM application_tabs at WHERE at.applicationid NOT IN (SELECT applicationid  FROM user_applications UNION select applicationid FROM group_applications)
) as d
WHERE a.id = d.applicationid
ORDER BY d.tabid, d.displayorder;
`

var (
	NewUserApplicationsTable = NewUserApplicationsTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name UserApplicationsTable
type UserApplicationsTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(app *models.UserApplications) error
	Delete(id int64) error

	GetUserApps(userid int64) ([]*models.UserApplications, error)

	GetUserAllowdApplications(userid int64) (models.UsersApplications, error)
}

type UserApplicationsTableImpl struct {
	database Database
}

// GetUserAllowdApplications implements UserApplicationsTable
func (i *UserApplicationsTableImpl) GetUserAllowdApplications(userid int64) (models.UsersApplications, error) {
	rows, err := i.database.Pool().Query(context.Background(), userAllowedQuery, userid)
	if err != nil {
		return nil, err
	}
	applist := models.UsersApplications{}
	var applicationid, tabid, displayorder int64
	var name, website, license, description, tilebackground, icon, templateappid string
	var enhanced bool
	_, err = pgx.ForEachRow(rows, []any{
		&tabid,
		&displayorder,
		&applicationid,
		&name,
		&website,
		&license,
		&description,
		&enhanced,
		&tilebackground,
		&icon,
		&templateappid,
	}, func() error {

		applist = append(applist, &models.UsersApplicationItem{
			Application: &models.Application{
				ID:             applicationid,
				TemplateAppid:  templateappid,
				Name:           name,
				Website:        website,
				License:        license,
				Description:    description,
				Enhanced:       enhanced,
				TileBackground: tilebackground,
				Icon:           icon},
			DisplayOrder: displayorder,
			TabID:        tabid,
		})
		return nil
	})

	return applist, err

}

// Drop implements UserApplicationsTable
func (i *UserApplicationsTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), userApplicationsTableDrop)
	return err

}

// GetAll implements UserApplicationsTable
func (i *UserApplicationsTableImpl) GetUserApps(userid int64) ([]*models.UserApplications, error) {
	rows, err := i.database.Pool().Query(context.Background(), userApplicationsTableQueryAll, userid)
	if err != nil {
		return nil, err
	}
	results := []*models.UserApplications{}
	var id, applicationid int64
	var name, website, license, description, tilebackground, icon, templateappid sql.NullString
	var enhanced bool
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&applicationid,
		&name,
		&website,
		&license,
		&description,
		&enhanced,
		&tilebackground,
		&icon,
		&templateappid,
	}, func() error {

		results = append(results, &models.UserApplications{
			ID:            id,
			UserID:        userid,
			ApplicationID: applicationid,
			Application: &models.Application{
				ID:             applicationid,
				TemplateAppid:  templateappid.String,
				Name:           name.String,
				Website:        website.String,
				License:        license.String,
				Description:    description.String,
				Enhanced:       enhanced,
				TileBackground: tilebackground.String,
				Icon:           icon.String,
			},
		})
		return nil
	})

	return results, err

}

// Create implements UserApplicationsTable
func (i *UserApplicationsTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), userApplicationsTableCreate)
	return err
}

// Delete implements UserApplicationsTable
func (i *UserApplicationsTableImpl) Delete(id int64) error {
	_, err := i.database.Pool().Exec(context.Background(), userApplicationsTableDeleteRow, id)
	return err

}

// Insert implements UserApplicationsTable
func (i *UserApplicationsTableImpl) Insert(app *models.UserApplications) error {
	err := i.database.Pool().QueryRow(context.Background(), userApplicationsTableInsertRow,
		app.UserID,
		app.ApplicationID,
	).Scan(&app.ID)
	return err
}

// Upgrade implements UserApplicationsTable
func (*UserApplicationsTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewUserApplicationsTableImpl(database Database) UserApplicationsTable {
	return &UserApplicationsTableImpl{
		database: database,
	}
}
