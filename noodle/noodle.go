package noodle

import (
	"io/ioutil"
	"os"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/mwinters-stuff/noodle/server/models"
)

// const pgport = 15432

// const pgport = 5432

type Noodle interface {
	Run()
}

type NoodleImpl struct {
}

// Run implements Noodle
func (*NoodleImpl) Run() {
	Logger.Info().Msg("Noodle Start")

	yfile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	config, err := yamltypes.UnmarshalConfig(yfile)

	db := database.NewDatabase(config)
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	err = db.Connect()
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	db.Drop()

	err = db.Create()
	if err != nil {
		Logger.Fatal().Msg(err.Error())
	}

	table8 := database.NewGroupApplicationsTable(db)
	table8.Drop()

	table9 := database.NewUserApplicationsTable(db)
	table9.Drop()

	table7 := database.NewApplicationTabTable(db)
	table7.Drop()

	table5 := database.NewApplicationsTable(db)
	table5.Drop()

	table1 := database.NewAppTemplateTable(db)
	table1.Drop()
	table1.Create()

	appTemplate1 := models.ApplicationTemplate{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}
	table1.Insert(appTemplate1)

	appTemplate2 := models.ApplicationTemplate{
		Appid:          "653caf8bdf55d6a99d77ceacd79f622353cd821a",
		Name:           "Adminer",
		Website:        "https://www.adminer.org",
		License:        "Apache License 2.0",
		Description:    "Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB.",
		Enhanced:       false,
		TileBackground: "light",
		Icon:           "adminer.svg",
		SHA:            "28ab6a65c96ce05b9c6aaaa73c643a83b87ad1e5",
	}
	table1.Insert(appTemplate2)

	table4 := database.NewUserGroupsTable(db)
	table4.Drop()

	table2 := database.NewUserTable(db, database.NewTableCache())
	table2.Drop()
	table2.Create()

	table3 := database.NewGroupTable(db)
	table3.Drop()
	table3.Create()

	table4.Create()

	user1 := models.User{
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}

	user2 := models.User{
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}

	table2.Insert(&user1)
	table2.Insert(&user2)

	group1 := database.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}
	group2 := database.Group{
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}
	table3.Insert(&group1)
	table3.Insert(&group2)

	usergroup1 := database.UserGroup{
		GroupId: group1.Id,
		UserId:  user1.ID,
	}
	usergroup2 := database.UserGroup{
		GroupId: group1.Id,
		UserId:  user2.ID,
	}
	usergroup3 := database.UserGroup{
		GroupId: group2.Id,
		UserId:  user1.ID,
	}

	table4.Insert(&usergroup1)
	table4.Insert(&usergroup2)
	table4.Insert(&usergroup3)

	//	table4.Delete(usergroup2)

	table4.GetUser(user2.ID)

	table5.Create()

	application1 := database.Application{
		TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	application2 := database.Application{
		Name:           "Adminer",
		Website:        "https://www.adminer.org",
		License:        "Apache License 2.0",
		Description:    "Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB.",
		Enhanced:       false,
		TileBackground: "light",
		Icon:           "adminer.svg",
	}

	table5.Insert(&application1)
	table5.Insert(&application2)

	table6 := database.NewTabTable(db)

	table6.Drop()
	table6.Create()

	tab1 := database.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}
	tab2 := database.Tab{
		Label:        "Apps",
		DisplayOrder: 2,
	}

	table6.Insert(&tab1)
	table6.Insert(&tab2)

	// table6.Update(tab1)
	// table6.Delete(tab1)
	table6.GetAll()

	table7.Create()

	at1 := database.ApplicationTab{
		TabId:         tab1.Id,
		ApplicationId: application1.Id,
		DisplayOrder:  1,
	}
	at2 := database.ApplicationTab{
		TabId:         tab1.Id,
		ApplicationId: application2.Id,
		DisplayOrder:  2,
	}
	at3 := database.ApplicationTab{
		TabId:         tab2.Id,
		ApplicationId: application2.Id,
		DisplayOrder:  1,
	}

	table7.Insert(&at1)
	table7.Insert(&at2)
	table7.Insert(&at3)

	// table7.Update(at3)

	// table7.GetTabApps(tab1.Id)

	table8.Create()

	ga1 := database.GroupApplications{
		GroupId:       group2.Id,
		ApplicationId: application1.Id,
	}

	ga2 := database.GroupApplications{
		GroupId:       group2.Id,
		ApplicationId: application2.Id,
	}

	table8.Insert(&ga1)
	table8.Insert(&ga2)

	// table8.GetGroupApps(group2.Id)

	// table8.Delete(ga1)

	table9.Create()

	ua1 := database.UserApplications{
		UserId:        user1.ID,
		ApplicationId: application1.Id,
	}

	ua2 := database.UserApplications{
		UserId:        user1.ID,
		ApplicationId: application2.Id,
	}

	table9.Insert(&ua1)
	table9.Insert(&ua2)

	table9.GetUserApps(1)
	table9.Delete(ua2)

}

func NewNoodle() Noodle {
	noodle := &NoodleImpl{}
	return noodle
}
