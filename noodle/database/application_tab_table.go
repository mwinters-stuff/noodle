package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ApplicationTab struct {
	Id            int
	ApplicationId int
	TabId         int
	DisplayOrder  int
	Application   Application
}

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

	Insert(tab *ApplicationTab) error
	Update(tab ApplicationTab) error
	Delete(tab ApplicationTab) error

	GetTabApps(tabid int) ([]ApplicationTab, error)
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
func (i *ApplicationTabTableImpl) GetTabApps(tabid int) ([]ApplicationTab, error) {
	rows, err := i.database.Pool().Query(context.Background(), applicationTabTableQueryAll, tabid)
	if err != nil {
		return nil, err
	}
	results := []ApplicationTab{}
	var id, displayorder, applicationid int
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

		results = append(results, ApplicationTab{
			Id:            id,
			TabId:         tabid,
			ApplicationId: applicationid,
			DisplayOrder:  displayorder,
			Application: Application{
				Id:             applicationid,
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
func (i *ApplicationTabTableImpl) Delete(tab ApplicationTab) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableDeleteRow, tab.Id)
	return err

}

// Insert implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Insert(tab *ApplicationTab) error {
	err := i.database.Pool().QueryRow(context.Background(), applicationTabTableInsertRow,
		tab.TabId,
		tab.ApplicationId,
		tab.DisplayOrder,
	).Scan(&tab.Id)
	return err
}

// Update implements ApplicationTabTable
func (i *ApplicationTabTableImpl) Update(tab ApplicationTab) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationTabTableUpdateRow,
		tab.Id,
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
