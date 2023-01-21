package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const groupTableCreate = `CREATE TABLE IF NOT EXISTS groups (
  id SERIAL PRIMARY KEY,
  dn VARCHAR(200) UNIQUE,
  name VARCHAR(100)
)`

const groupTableInsertRow = `INSERT INTO groups (dn, name) VALUES ($1, $2) RETURNING id`
const groupTableDrop = `DROP TABLE groups`
const groupTableUpdateRow = `UPDATE groups SET dn = $2,name = $3 WHERE id = $1`
const groupTableDeleteRow = `DELETE FROM groups WHERE id = $1`
const groupTableQueryRowsDN = `SELECT * FROM groups WHERE dn = $1`
const groupTableQueryRowsID = `SELECT * FROM groups WHERE id = $1`
const groupTableQueryAll = `SELECT * FROM groups ORDER BY name`
const groupTableQueryExistsDN = `SELECT COUNT(*) FROM groups WHERE dn = $1`
const groupTableQueryExistsName = `SELECT COUNT(*) FROM groups WHERE name = $1`

var (
	NewGroupTable = NewGroupTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name GroupTable
type GroupTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(group *models.Group) error
	Update(group models.Group) error
	Delete(group models.Group) error

	GetDN(dn string) (models.Group, error)
	GetID(id int64) (models.Group, error)
	GetAll() ([]*models.Group, error)
	ExistsDN(dn string) (bool, error)
	ExistsName(groupname string) (bool, error)
}

type GroupTableImpl struct {
	database Database
}

// Drop implements GroupTable
func (i *GroupTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), groupTableDrop)
	return err

}

// GetAll implements GroupTable
func (i *GroupTableImpl) getQuery(query string, value any) ([]*models.Group, error) {
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
	results := []*models.Group{}
	var dn, name string
	var id int64
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&dn,
		&name,
	}, func() error {

		results = append(results, &models.Group{
			ID:   id,
			DN:   dn,
			Name: name,
		})
		return nil
	})

	return results, err
}

// GetAll implements GroupTable
func (i *GroupTableImpl) GetAll() ([]*models.Group, error) {
	return i.getQuery(groupTableQueryAll, nil)
}

// GetDN implements GroupTable
func (i *GroupTableImpl) GetDN(dn string) (models.Group, error) {
	rows, err := i.getQuery(groupTableQueryRowsDN, dn)
	if err == nil {
		return *rows[0], nil
	}
	return models.Group{}, err
}

// GetID implements GroupTable
func (i *GroupTableImpl) GetID(id int64) (models.Group, error) {
	rows, err := i.getQuery(groupTableQueryRowsID, id)
	if err == nil {
		return *rows[0], nil
	}
	return models.Group{}, err

}

// Create implements GroupTable
func (i *GroupTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), groupTableCreate)
	return err
}

// Delete implements GroupTable
func (i *GroupTableImpl) Delete(group models.Group) error {
	_, err := i.database.Pool().Exec(context.Background(), groupTableDeleteRow, group.ID)
	return err

}

// Exists implements GroupTable
func (i *GroupTableImpl) ExistsDN(dn string) (bool, error) {
	var found int
	err := i.database.Pool().QueryRow(context.Background(), groupTableQueryExistsDN, dn).Scan(&found)
	return found > 0, err
}

func (i *GroupTableImpl) ExistsName(name string) (bool, error) {
	var found int
	err := i.database.Pool().QueryRow(context.Background(), groupTableQueryExistsName, name).Scan(&found)
	return found > 0, err
}

// Insert implements GroupTable
func (i *GroupTableImpl) Insert(group *models.Group) error {
	err := i.database.Pool().QueryRow(context.Background(), groupTableInsertRow,
		group.DN,
		group.Name,
	).Scan(&group.ID)
	return err
}

// Update implements GroupTable
func (i *GroupTableImpl) Update(group models.Group) error {
	_, err := i.database.Pool().Exec(context.Background(), groupTableUpdateRow,
		group.ID,
		group.DN,
		group.Name,
	)
	return err
}

// Upgrade implements GroupTable
func (*GroupTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewGroupTableImpl(database Database) GroupTable {
	return &GroupTableImpl{
		database: database,
	}
}
