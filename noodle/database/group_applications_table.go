package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const groupApplicationsTableCreate = `CREATE TABLE IF NOT EXISTS group_applications (
  id SERIAL PRIMARY KEY,
  groupid int REFERENCES groups(id) ON DELETE CASCADE,
  applicationid int REFERENCES applications(id) ON DELETE CASCADE
)`

const groupApplicationsTableInsertRow = `INSERT INTO group_applications (groupid, applicationid) VALUES ($1, $2) RETURNING id`
const groupApplicationsTableDrop = `DROP TABLE group_applications`
const groupApplicationsTableDeleteRow = `DELETE FROM group_applications WHERE id = $1`
const groupApplicationsTableQueryAll = `SELECT ga.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM group_applications ga, applications app WHERE ga.groupid = $1 AND app.id = ga.applicationid`

var (
	NewGroupApplicationsTable = NewGroupApplicationsTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name GroupApplicationsTable
type GroupApplicationsTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(app *models.GroupApplications) error
	Delete(id int64) error

	GetGroupApps(groupid int64) ([]*models.GroupApplications, error)
}

type GroupApplicationsTableImpl struct {
	database Database
}

// Drop implements GroupApplicationsTable
func (i *GroupApplicationsTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), groupApplicationsTableDrop)
	return err

}

// GetAll implements GroupApplicationsTable
func (i *GroupApplicationsTableImpl) GetGroupApps(groupid int64) ([]*models.GroupApplications, error) {
	rows, err := i.database.Pool().Query(context.Background(), groupApplicationsTableQueryAll, groupid)
	if err != nil {
		return nil, err
	}
	results := []*models.GroupApplications{}
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

		results = append(results, &models.GroupApplications{
			ID:            id,
			GroupID:       groupid,
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

// Create implements GroupApplicationsTable
func (i *GroupApplicationsTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), groupApplicationsTableCreate)
	return err
}

// Delete implements GroupApplicationsTable
func (i *GroupApplicationsTableImpl) Delete(id int64) error {
	_, err := i.database.Pool().Exec(context.Background(), groupApplicationsTableDeleteRow, id)
	return err

}

// Insert implements GroupApplicationsTable
func (i *GroupApplicationsTableImpl) Insert(app *models.GroupApplications) error {
	err := i.database.Pool().QueryRow(context.Background(), groupApplicationsTableInsertRow,
		app.GroupID,
		app.ApplicationID,
	).Scan(&app.ID)
	return err
}

// Upgrade implements GroupApplicationsTable
func (*GroupApplicationsTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewGroupApplicationsTableImpl(database Database) GroupApplicationsTable {
	return &GroupApplicationsTableImpl{
		database: database,
	}
}
