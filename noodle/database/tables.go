package database

var (
	NewTables = NewTablesImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name Tables
type Tables interface {
	InitTables(db Database)
	Create() error
	Drop() error

	AppTemplateTable() AppTemplateTable
	ApplicationTabTable() ApplicationTabTable
	ApplicationsTable() ApplicationsTable
	GroupApplicationsTable() GroupApplicationsTable
	GroupTable() GroupTable
	TabTable() TabTable
	UserApplicationsTable() UserApplicationsTable
	UserGroupsTable() UserGroupsTable
	UserTable() UserTable
}

type TablesImpl struct {
	appTemplateTable       AppTemplateTable
	applicationTabTable    ApplicationTabTable
	applicationsTable      ApplicationsTable
	groupApplicationsTable GroupApplicationsTable
	groupTable             GroupTable
	tabTable               TabTable
	userApplicationsTable  UserApplicationsTable
	userGroupsTable        UserGroupsTable
	userTable              UserTable
}

func (i *TablesImpl) InitTables(db Database) {
	i.appTemplateTable = NewAppTemplateTable(db)
	i.applicationTabTable = NewApplicationTabTable(db)
	i.applicationsTable = NewApplicationsTable(db)
	i.groupApplicationsTable = NewGroupApplicationsTable(db)
	i.groupTable = NewGroupTable(db)
	i.tabTable = NewTabTable(db)
	i.userApplicationsTable = NewUserApplicationsTable(db)
	i.userGroupsTable = NewUserGroupsTable(db)
	i.userTable = NewUserTable(db)
}

func (i *TablesImpl) Create() error {
	var err error
	if err = i.appTemplateTable.Create(); err != nil {
		return err
	}

	if err = i.userTable.Create(); err != nil {
		return err
	}

	if err = i.groupTable.Create(); err != nil {
		return err
	}

	if err = i.applicationsTable.Create(); err != nil {
		return err
	}

	if err = i.tabTable.Create(); err != nil {
		return err
	}

	if err = i.applicationTabTable.Create(); err != nil {
		return err
	}

	if err = i.groupApplicationsTable.Create(); err != nil {
		return err
	}

	if err = i.userApplicationsTable.Create(); err != nil {
		return err
	}

	if err = i.userGroupsTable.Create(); err != nil {
		return err
	}
	return nil
}

func (i *TablesImpl) Drop() error {
	var err error
	if err = i.applicationTabTable.Drop(); err != nil {
		return err
	}

	if err = i.groupApplicationsTable.Drop(); err != nil {
		return err
	}

	if err = i.userApplicationsTable.Drop(); err != nil {
		return err
	}

	if err = i.userGroupsTable.Drop(); err != nil {
		return err
	}

	if err = i.userTable.Drop(); err != nil {
		return err
	}

	if err = i.tabTable.Drop(); err != nil {
		return err
	}

	if err = i.groupTable.Drop(); err != nil {
		return err
	}

	if err = i.applicationsTable.Drop(); err != nil {
		return err
	}

	if err = i.appTemplateTable.Drop(); err != nil {
		return err
	}

	return nil
}

func (i *TablesImpl) AppTemplateTable() AppTemplateTable {
	return i.appTemplateTable
}

// ApplicationTabTable implements Database
func (i *TablesImpl) ApplicationTabTable() ApplicationTabTable {
	return i.applicationTabTable
}

// ApplicationsTable implements Database
func (i *TablesImpl) ApplicationsTable() ApplicationsTable {
	return i.applicationsTable
}

// GroupApplicationsTable implements Database
func (i *TablesImpl) GroupApplicationsTable() GroupApplicationsTable {
	return i.groupApplicationsTable
}

// GroupTable implements Database
func (i *TablesImpl) GroupTable() GroupTable {
	return i.groupTable
}

// TabTable implements Database
func (i *TablesImpl) TabTable() TabTable {
	return i.tabTable
}

// UserApplicationsTable implements Database
func (i *TablesImpl) UserApplicationsTable() UserApplicationsTable {
	return i.userApplicationsTable
}

// UserGroupsTable implements Database
func (i *TablesImpl) UserGroupsTable() UserGroupsTable {
	return i.userGroupsTable
}

// UserTable implements Database
func (i *TablesImpl) UserTable() UserTable {
	return i.userTable
}

func NewTablesImpl() Tables {
	return &TablesImpl{}
}
