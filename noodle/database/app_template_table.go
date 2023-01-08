package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
)

const appTemplateTableCreate = `CREATE TABLE IF NOT EXISTS application_template (
  appid CHAR(40) PRIMARY KEY,
  name VARCHAR(20) UNIQUE,
  website VARCHAR(100) UNIQUE,
  license VARCHAR(100),
  description VARCHAR(1000),
  enhanced BOOL,
  tilebackground VARCHAR(256),
  icon VARCHAR(256), 
  sha CHAR(40)
)`
const appTemplateTableDrop = `DROP TABLE application_template`
const appTemplateTableIndexCreate = `CREATE INDEX IF NOT EXISTS application_template_idx1 ON application_template(name)`
const appTemplateTableInsertRow = `INSERT INTO application_template (appid,name,website,license,description,enhanced,tilebackground,icon,sha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
const appTemplateTableUpdateRow = `UPDATE application_template SET name = $2,website = $3,license = $4,description = $5,enhanced = $6,tilebackground = $7,icon = $8,sha = $9 WHERE appid = $1`
const appTemplateTableDeleteRow = `DELETE FROM application_template WHERE appid = $1`
const appTemplateTableQueryRows = `SELECT * FROM application_template WHERE name LIKE $1`
const appTemplateTableQueryExists = `SELECT COUNT(*) FROM application_template WHERE appid = $1`

var (
	NewAppTemplateTable = NewAppTemplateTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name AppTemplateTable
type AppTemplateTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(app jsontypes.App) error
	Update(app jsontypes.App) error
	Delete(app jsontypes.App) error

	Search(search string) ([]jsontypes.App, error)
	Exists(appid string) (bool, error)
}

type AppTemplateTableImpl struct {
	database Database
}

// Drop implements AppTemplateTable
func (i *AppTemplateTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), appTemplateTableDrop)
	return err
}

// Exists implements AppTemplateTable
func (i *AppTemplateTableImpl) Exists(appid string) (bool, error) {
	var found int
	err := i.database.Pool().QueryRow(context.Background(), appTemplateTableQueryExists, appid).Scan(&found)
	return found > 0, err
}

// Create implements AppTemplateTable
func (i *AppTemplateTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), appTemplateTableCreate)
	if err == nil {
		_, err = i.database.Pool().Exec(context.Background(), appTemplateTableIndexCreate)
	}
	return err
}

// Delete implements AppTemplateTable
func (i *AppTemplateTableImpl) Delete(app jsontypes.App) error {
	_, err := i.database.Pool().Exec(context.Background(), appTemplateTableDeleteRow, app.Appid)
	return err
}

// Insert implements AppTemplateTable
func (i *AppTemplateTableImpl) Insert(app jsontypes.App) error {
	_, err := i.database.Pool().Exec(context.Background(), appTemplateTableInsertRow,
		app.Appid,
		app.Name,
		app.Website,
		app.License,
		app.Description,
		app.Enhanced,
		app.TileBackground,
		app.Icon,
		app.SHA)
	return err
}

// Search implements AppTemplateTable
func (i *AppTemplateTableImpl) Search(search string) ([]jsontypes.App, error) {
	rows, err := i.database.Pool().Query(context.Background(), appTemplateTableQueryRows, fmt.Sprintf("%%%s%%", search))
	if err != nil {
		return nil, err
	}
	results := []jsontypes.App{}
	var appid, sha pgtype.Text
	var name, website, license, description, tilebackground, icon string
	var enhanced bool
	_, err = pgx.ForEachRow(rows, []any{&appid,
		&name,
		&website,
		&license,
		&description,
		&enhanced,
		&tilebackground,
		&icon,
		&sha}, func() error {

		results = append(results, jsontypes.App{
			Appid:          appid.String,
			Name:           name,
			Website:        website,
			License:        license,
			Description:    description,
			Enhanced:       enhanced,
			TileBackground: tilebackground,
			Icon:           icon,
			SHA:            sha.String,
		})
		return nil
	})

	return results, err
}

// Update implements AppTemplateTable
func (i *AppTemplateTableImpl) Update(app jsontypes.App) error {
	_, err := i.database.Pool().Exec(context.Background(), appTemplateTableUpdateRow,
		app.Appid,
		app.Name,
		app.Website,
		app.License,
		app.Description,
		app.Enhanced,
		app.TileBackground,
		app.Icon,
		app.SHA)
	return err
}

// Upgrade implements AppTemplateTable
func (*AppTemplateTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewAppTemplateTableImpl(database Database) AppTemplateTable {
	return &AppTemplateTableImpl{
		database: database,
	}
}
