package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const tabTableCreate = `CREATE TABLE IF NOT EXISTS tabs (
  id SERIAL PRIMARY KEY,
  Label VARCHAR(200) UNIQUE,
  DisplayOrder int
)`

const tabTableInsertRow = `INSERT INTO tabs (label,displayorder) VALUES ($1, $2) RETURNING id`
const tabTableDrop = `DROP TABLE tabs`
const tabTableUpdateRow = `UPDATE tabs SET label = $2,displayorder = $3 WHERE id = $1`
const tabTableDeleteRow = `DELETE FROM tabs WHERE id = $1`
const tabTableQueryAll = `SELECT * FROM tabs ORDER BY displayorder`

var (
	NewTabTable = NewTabTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name TabTable
type TabTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(tab *models.Tab) error
	Update(tab models.Tab) error
	Delete(id int64) error

	GetAll() ([]*models.Tab, error)
}

type TabTableImpl struct {
	database Database
}

// Drop implements TabTable
func (i *TabTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), tabTableDrop)
	return err

}

// GetAll implements TabTable
func (i *TabTableImpl) GetAll() ([]*models.Tab, error) {
	rows, err := i.database.Pool().Query(context.Background(), tabTableQueryAll)
	if err != nil {
		return nil, err
	}
	results := []*models.Tab{}
	var label string
	var id, displayorder int64
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&label,
		&displayorder,
	}, func() error {

		results = append(results, &models.Tab{
			ID:           id,
			Label:        label,
			DisplayOrder: displayorder,
		})
		return nil
	})

	return results, err

}

// Create implements TabTable
func (i *TabTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), tabTableCreate)
	return err
}

// Delete implements TabTable
func (i *TabTableImpl) Delete(id int64) error {
	_, err := i.database.Pool().Exec(context.Background(), tabTableDeleteRow, id)
	return err

}

// Insert implements TabTable
func (i *TabTableImpl) Insert(tab *models.Tab) error {
	err := i.database.Pool().QueryRow(context.Background(), tabTableInsertRow,
		tab.Label,
		tab.DisplayOrder,
	).Scan(&tab.ID)
	return err
}

// Update implements TabTable
func (i *TabTableImpl) Update(tab models.Tab) error {
	_, err := i.database.Pool().Exec(context.Background(), tabTableUpdateRow,
		tab.ID,
		tab.Label,
		tab.DisplayOrder,
	)
	return err
}

// Upgrade implements TabTable
func (*TabTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewTabTableImpl(database Database) TabTable {
	return &TabTableImpl{
		database: database,
	}
}
