package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const applicationTabTableCreate = `CREATE TABLE IF NOT EXISTS application_tabs (
  id SERIAL PRIMARY KEY,
  tabid int REFERENCES tabs(id) ON DELETE CASCADE,
  applicationid int REFERENCES applications(id) ON DELETE CASCADE,
  displayorder int
)`

const applicationTabTableInsertRow = `INSERT INTO application_tabs (tabid, applicationid, displayorder) VALUES ($1, $2, $3) RETURNING id`
const applicationTabTableDrop = `DROP TABLE application_tabs`
const applicationTabTableUpdateRow = `UPDATE application_tabs SET displayorder = $2 WHERE id = $1`
const applicationTabTableDeleteRow = `DELETE FROM application_tabs WHERE id = $1`
const applicationTabTableQueryAll = `SELECT at.id, at.displayorder, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM application_tabs at, applications app WHERE at.tabid = $1 AND app.id = at.applicationid ORDER BY at.displayorder`

var (
	NewApplicationTabTable = NewApplicationTabTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name ApplicationTabTable
type ApplicationTabTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(tab *models.ApplicationTab) error
	Update(tab models.ApplicationTab) error
	Delete(id int64) error

	GetTabApps(tabid int64) ([]*models.ApplicationTab, error)
}

type ApplicationTabTableImpl struct {
	database Database
}

// Drop implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableDrop)
	return err

}

// GetAll implements ApplicationTabTable
func (i *ApplicationTabTableImpl) GetTabApps(tabid int64) ([]*models.ApplicationTab, error) {
	rows, err := i.database.Pool().Query(context.Background(), applicationTabTableQueryAll, tabid)
	if err != nil {
		return nil, err
	}
	results := []*models.ApplicationTab{}
	var id, displayorder, applicationid int64
	var name, website, license, description, tilebackground, icon string
	var enhanced bool
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&displayorder,
		&applicationid,
		&name,
		&website,
		&license,
		&description,
		&enhanced,
		&tilebackground,
		&icon,
	}, func() error {

		results = append(results, &models.ApplicationTab{
			ID:            id,
			TabID:         tabid,
			ApplicationID: applicationid,
			DisplayOrder:  displayorder,
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

// Create implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableCreate)
	return err
}

// Delete implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Delete(id int64) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableDeleteRow, id)
	return err

}

// Insert implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Insert(tab *models.ApplicationTab) error {
	err := i.database.Pool().QueryRow(context.Background(), applicationTabTableInsertRow,
		tab.TabID,
		tab.ApplicationID,
		tab.DisplayOrder,
	).Scan(&tab.ID)
	return err
}

// Update implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Update(tab models.ApplicationTab) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableUpdateRow,
		tab.ID,
		tab.DisplayOrder,
	)
	return err
}

// Upgrade implements ApplicationTabTable
func (*ApplicationTabTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewApplicationTabTableImpl(database Database) ApplicationTabTable {
	return &ApplicationTabTableImpl{
		database: database,
	}
}
