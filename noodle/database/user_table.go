package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/server/models"
)

const userTableCreate = `CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE,
  dn VARCHAR(200) UNIQUE,
  displayname VARCHAR(100),
  givenname VARCHAR(100),
  surname VARCHAR(100),
  uidnumber INTEGER
)`

const userTableInsertRow = `INSERT INTO users (username, dn, displayname, givenname, surname, uidnumber) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
const userTableDrop = `DROP TABLE users`
const userTableUpdateRow = `UPDATE users SET username = $2,dn = $3,displayname = $4,givenname = $5,surname = $6,uidnumber = $7 WHERE id = $1`
const userTableDeleteRow = `DELETE FROM users WHERE id = $1`
const userTableQueryRowsDN = `SELECT * FROM users WHERE dn = $1`
const userTableQueryRowsID = `SELECT * FROM users WHERE id = $1`
const userTableQueryAll = `SELECT * FROM users ORDER BY username`
const userTableQueryExistsDN = `SELECT COUNT(*) FROM users WHERE dn = $1`
const userTableQueryExistsUsername = `SELECT COUNT(*) FROM users WHERE username = $1`

var (
	NewUserTable = NewUserTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name UserTable
type UserTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(user *models.User) error
	Update(user models.User) error
	Delete(user models.User) error

	GetDN(dn string) (models.User, error)
	GetID(id int) (models.User, error)
	GetAll() ([]models.User, error)
	ExistsDN(dn string) (bool, error)
	ExistsUsername(username string) (bool, error)
}

type UserTableImpl struct {
	database Database
	cache    TableCache[models.User]
}

// Drop implements UserTable
func (i *UserTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), userTableDrop)
	return err
}

// GetAll implements UserTable
func (i *UserTableImpl) getQuery(query string, value any) ([]models.User, error) {
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
	results := []models.User{}
	var username, dn, displayname, givenname, surname string
	var id, uidnumber int64
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&username,
		&dn,
		&displayname,
		&givenname,
		&surname,
		&uidnumber,
	}, func() error {

		user := models.User{
			ID:          id,
			Username:    username,
			DN:          dn,
			DisplayName: displayname,
			GivenName:   givenname,
			Surname:     surname,
			UIDNumber:   uidnumber,
		}
		results = append(results, user)
		i.cache.Add(id, user)
		return nil
	})

	return results, err
}

// GetAll implements UserTable
func (i *UserTableImpl) GetAll() ([]models.User, error) {
	return i.getQuery(userTableQueryAll, nil)
}

// GetDN implements UserTable
func (i *UserTableImpl) GetDN(dn string) (models.User, error) {
	ok, user := i.cache.Find((func(index int64, value models.User) bool {
		return value.DN == dn
	}))
	if ok {
		return *user, nil
	}

	rows, err := i.getQuery(userTableQueryRowsDN, dn)
	if err == nil {
		i.cache.Add(rows[0].ID, rows[0])
		return rows[0], nil
	}
	return models.User{}, err
}

// GetID implements UserTable
func (i *UserTableImpl) GetID(id int) (models.User, error) {
	found, user := i.cache.GetID(int64(id))
	if found {
		return user, nil
	}

	rows, err := i.getQuery(userTableQueryRowsID, id)
	if err == nil {
		return rows[0], nil
	}
	return models.User{}, err

}

// Create implements UserTable
func (i *UserTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), userTableCreate)
	return err
}

// Delete implements UserTable
func (i *UserTableImpl) Delete(user models.User) error {
	_, err := i.database.Pool().Exec(context.Background(), userTableDeleteRow, user.ID)
	if err == nil {
		i.cache.DeleteIndex(user.ID)
	}
	return err
}

// Exists implements UserTable
func (i *UserTableImpl) ExistsDN(dn string) (bool, error) {
	ok, _ := i.cache.Find((func(index int64, value models.User) bool {
		return value.DN == dn
	}))
	if ok {
		return true, nil
	}

	var found int
	err := i.database.Pool().QueryRow(context.Background(), userTableQueryExistsDN, dn).Scan(&found)
	return found > 0, err
}

func (i *UserTableImpl) ExistsUsername(username string) (bool, error) {
	ok, _ := i.cache.Find((func(index int64, value models.User) bool {
		return value.Username == username
	}))
	if ok {
		return true, nil
	}
	var found int
	err := i.database.Pool().QueryRow(context.Background(), userTableQueryExistsUsername, username).Scan(&found)
	return found > 0, err
}

// Insert implements UserTable
func (i *UserTableImpl) Insert(user *models.User) error {
	err := i.database.Pool().QueryRow(context.Background(), userTableInsertRow,
		user.Username,
		user.DN,
		user.DisplayName,
		user.GivenName,
		user.Surname,
		user.UIDNumber,
	).Scan(&user.ID)
	if err == nil {
		i.cache.Add(user.ID, *user)
	}
	return err
}

// Update implements UserTable
func (i *UserTableImpl) Update(user models.User) error {
	_, err := i.database.Pool().Exec(context.Background(), userTableUpdateRow,
		user.ID,
		user.Username,
		user.DN,
		user.DisplayName,
		user.GivenName,
		user.Surname,
		user.UIDNumber,
	)
	if err == nil {
		i.cache.Update(user.ID, user)
	}
	return err
}

// Upgrade implements UserTable
func (*UserTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewUserTableImpl(database Database, cache TableCache[models.User]) UserTable {
	return &UserTableImpl{
		database: database,
		cache:    cache,
	}
}
