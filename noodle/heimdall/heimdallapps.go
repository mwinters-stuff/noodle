package heimdall

import (
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
)

type Heimdall interface {
	UpdateFromServer()
	FindApp(search string) jsontypes.App
}

type HeimdallImpl struct {
	database *database.Database
}

// FindApp implements Heimdall
func (*HeimdallImpl) FindApp(search string) jsontypes.App {
	panic("unimplemented")
}

// UpdateFromServer implements Heimdall
func (*HeimdallImpl) UpdateFromServer() {
	panic("unimplemented")
}

func NewHeimdall(database *database.Database) Heimdall {
	return &HeimdallImpl{
		database: database,
	}
}
