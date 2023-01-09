package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Application struct {
	Id             int
	TemplateAppid  string
	Name           string
	Website        string
	License        string
	Description    string
	Enhanced       bool
	TileBackground string
	Icon           string
}

const applicationsTableCreate = `CREATE TABLE IF NOT EXISTS applications (
  id SERIAL PRIMARY KEY,
  template_appid CHAR(40) REFERENCES application_template(appid) ON DELETE SET NULL,
  name VARCHAR(20),
  website VARCHAR(100),
  license VARCHAR(100),
  description VARCHAR(1000),
  enhanced BOOL,
  tilebackground VARCHAR(256),
  icon VARCHAR(256)
)`
const applicationsTableDrop = `DROP TABLE applications`
const applicationsTableInsertRow1 = `INSERT INTO applications (template_appid,name,website,license,description,enhanced,tilebackground,icon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
const applicationsTableInsertRow2 = `INSERT INTO applications (name,website,license,description,enhanced,tilebackground,icon) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
const applicationsTableUpdateRow = `UPDATE applications SET template_appid = $2, name = $3, website = $4, license = $5, description = $6, enhanced = $7,tilebackground = $8,icon = $9 WHERE id = $1`
const applicationsTableDeleteRow = `DELETE FROM applications WHERE id = $1`
const applicationsTableQueryID = `SELECT * FROM applications WHERE id = $1`
const applicationsTableQueryTemplateID = `SELECT * FROM applications WHERE template_appid = $1`

var (
	NewApplicationsTable = NewApplicationsTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name ApplicationsTable
type ApplicationsTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(app *Application) error
	Update(app Application) error
	Delete(app Application) error

	GetID(id int) (Application, error)
	GetTemplateID(appid string) ([]Application, error)
}

type ApplicationsTableImpl struct {
	database Database
}

func (i *ApplicationsTableImpl) getQuery(query string, value any) ([]Application, error) {
	rows, err := i.database.Pool().Query(context.Background(), query, value)
	if err != nil {
		return nil, err
	}
	results := []Application{}

	var id int
	var templateappid pgtype.Text
	var name, website, license, description, tilebackground, icon string
	var enhanced bool
	_, err = pgx.ForEachRow(rows, []any{&id,
		&templateappid,
		&name,
		&website,
		&license,
		&description,
		&enhanced,
		&tilebackground,
		&icon,
	}, func() error {

		results = append(results, Application{
			Id:             id,
			TemplateAppid:  templateappid.String,
			Name:           name,
			Website:        website,
			License:        license,
			Description:    description,
			Enhanced:       enhanced,
			TileBackground: tilebackground,
			Icon:           icon,
		})
		return nil
	})
	return results, err
}

// GetID implements ApplicationsTable
func (i *ApplicationsTableImpl) GetID(id int) (Application, error) {
	result, err := i.getQuery(applicationsTableQueryID, id)
	if err != nil {
		return Application{}, err
	}
	return result[0], nil
}

// GetTemplateID implements ApplicationsTable
func (i *ApplicationsTableImpl) GetTemplateID(appid string) ([]Application, error) {
	return i.getQuery(applicationsTableQueryTemplateID, appid)
}

// Drop implements ApplicationsTable
func (i *ApplicationsTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), applicationsTableDrop)
	return err
}

// Create implements ApplicationsTable
func (i *ApplicationsTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), applicationsTableCreate)
	return err
}

// Delete implements ApplicationsTable
func (i *ApplicationsTableImpl) Delete(app Application) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationsTableDeleteRow, app.Id)
	return err
}

// Insert implements ApplicationsTable
func (i *ApplicationsTableImpl) Insert(app *Application) error {
	if app.TemplateAppid != "" {
		return i.database.Pool().QueryRow(context.Background(), applicationsTableInsertRow1,
			app.TemplateAppid,
			app.Name,
			app.Website,
			app.License,
			app.Description,
			app.Enhanced,
			app.TileBackground,
			app.Icon,
		).Scan(&app.Id)

	} else {
		return i.database.Pool().QueryRow(context.Background(), applicationsTableInsertRow2,
			app.Name,
			app.Website,
			app.License,
			app.Description,
			app.Enhanced,
			app.TileBackground,
			app.Icon,
		).Scan(&app.Id)
	}

}

// Update implements ApplicationsTable
func (i *ApplicationsTableImpl) Update(app Application) error {
	_, err := i.database.Pool().Exec(context.Background(), applicationsTableUpdateRow,
		app.Id,
		app.TemplateAppid,
		app.Name,
		app.Website,
		app.License,
		app.Description,
		app.Enhanced,
		app.TileBackground,
		app.Icon)
	return err
}

// Upgrade implements ApplicationsTable
func (*ApplicationsTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewApplicationsTableImpl(database Database) ApplicationsTable {
	return &ApplicationsTableImpl{
		database: database,
	}
}