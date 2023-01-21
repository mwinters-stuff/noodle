package database

import (
	"context"

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
const userApplicationsTableQueryAll = `SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid`

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
}

type UserApplicationsTableImpl struct {
	database Database
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
	var name, website, license, description, tilebackground, icon string
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
	}, func() error {

		results = append(results, &models.UserApplications{
			ID:            id,
			UserID:        userid,
			ApplicationID: applicationid,
			Application: &models.Application{
				ID:             applicationid,
				Name:           name,
				Website:        website,
				License:        license,
				Description:    description,
				Enhanced:       enhanced,
				TileBackground: tilebackground,
				Icon:           icon,
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
