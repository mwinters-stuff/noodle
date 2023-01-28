package database

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mwinters-stuff/noodle/server/models"
)

const UserSessionTableCreate = `CREATE TABLE IF NOT EXISTS user_sessions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  token VARCHAR(100) UNIQUE NOT NULL,
  issued TIMESTAMP NOT NULL DEFAULT NOW(),
  expires TIMESTAMP NOT NULL DEFAULT NOW() + interval '1 month'
)`

const UserSessionTableInsertRow = `INSERT INTO user_sessions (user_id, token) VALUES ($1, $2) RETURNING id, issued, expires`
const UserSessionTableDrop = `DROP TABLE user_sessions`
const UserSessionTableDeleteRow = `DELETE FROM user_sessions WHERE id = $1`
const UserSessionTableQueryRowsUserID = `SELECT * FROM user_sessions WHERE user_id = $1`
const UserSessionTableQueryRowsToken = `SELECT * FROM user_sessions WHERE token = $1 AND expires > NOW()`
const UserSessionTableDeleteExpired = `DELETE FROM user_sessions WHERE expires < NOW()`

var (
	NewUserSessionTable = NewUserSessionTableImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name UserSessionTable
type UserSessionTable interface {
	Create() error
	Upgrade(old_version, new_verison int) error
	Drop() error

	Insert(UserSession *models.UserSession) error
	Delete(id int64) error
	DeleteExpired() error

	GetUser(user_id int64) ([]*models.UserSession, error)
	GetToken(token string) (models.UserSession, error)
}

type UserSessionTableImpl struct {
	database Database
}

// Drop implements UserSessionTable
func (i *UserSessionTableImpl) Drop() error {
	_, err := i.database.Pool().Exec(context.Background(), UserSessionTableDrop)
	return err

}

// GetAll implements UserSessionTable
func (i *UserSessionTableImpl) getQuery(query string, value any) ([]*models.UserSession, error) {
	rows, err := i.database.Pool().Query(context.Background(), query, value)

	if err != nil {
		return nil, err
	}
	results := []*models.UserSession{}
	var token string
	var id, user_id int64
	var issued, expires pgtype.Timestamptz
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&user_id,
		&token,
		&issued,
		&expires,
	}, func() error {

		results = append(results, &models.UserSession{
			ID:      id,
			UserID:  user_id,
			Token:   token,
			Issued:  strfmt.DateTime(issued.Time),
			Expires: strfmt.DateTime(expires.Time),
		})
		return nil
	})

	return results, err
}

// GetAll implements UserSessionTable
func (i *UserSessionTableImpl) GetUser(user_id int64) ([]*models.UserSession, error) {
	return i.getQuery(UserSessionTableQueryRowsUserID, user_id)
}

// GetDN implements UserSessionTable
func (i *UserSessionTableImpl) GetToken(token string) (models.UserSession, error) {
	rows, err := i.getQuery(UserSessionTableQueryRowsToken, token)
	if err == nil {
		return *rows[0], nil
	}
	return models.UserSession{}, err
}

// Create implements UserSessionTable
func (i *UserSessionTableImpl) Create() error {
	_, err := i.database.Pool().Exec(context.Background(), UserSessionTableCreate)
	return err
}

// Delete implements UserSessionTable
func (i *UserSessionTableImpl) Delete(id int64) error {
	_, err := i.database.Pool().Exec(context.Background(), UserSessionTableDeleteRow, id)
	return err

}

func (i *UserSessionTableImpl) DeleteExpired() error {
	_, err := i.database.Pool().Exec(context.Background(), UserSessionTableDeleteExpired)
	return err

}

// Insert implements UserSessionTable
func (i *UserSessionTableImpl) Insert(userSession *models.UserSession) error {
	var issued, expires pgtype.Timestamptz

	err := i.database.Pool().QueryRow(context.Background(), UserSessionTableInsertRow,
		userSession.UserID,
		userSession.Token,
	).Scan(&userSession.ID, &issued, &expires)
	userSession.Issued = strfmt.DateTime(issued.Time)
	userSession.Expires = strfmt.DateTime(expires.Time)
	return err
}

// Upgrade implements UserSessionTable
func (*UserSessionTableImpl) Upgrade(old_version int, new_verison int) error {
	panic("unimplemented")
}

func NewUserSessionTableImpl(database Database) UserSessionTable {
	return &UserSessionTableImpl{
		database: database,
	}
}
