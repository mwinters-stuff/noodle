package noodle

import (
	"io/ioutil"
	"os"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
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

	table1 := database.NewAppTemplateTable(db)
	table1.Drop()
	table1.Create()

	appTemplate := jsontypes.App{
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
	table1.Insert(appTemplate)

	table4 := database.NewUserGroupsTable(db)
	table4.Drop()

	table2 := database.NewUserTable(db)
	table2.Drop()
	table2.Create()

	table3 := database.NewGroupTable(db)
	table3.Drop()
	table3.Create()

	table4.Create()

	// user1 := database.User{
	// 	DN:          "CN=bob,DC=example,DC=nz",
	// 	Username:    "bobe",
	// 	DisplayName: "bobextample",
	// 	Surname:     "Extample",
	// 	GivenName:   "Bob",
	// 	UidNumber:   1001,
	// }

	// user2 := database.User{
	// 	DN:          "CN=jack,DC=example,DC=nz",
	// 	Username:    "jack",
	// 	DisplayName: "Jack M",
	// 	Surname:     "M",
	// 	GivenName:   "Jack",
	// 	UidNumber:   1002,
	// }

	// table2.Insert(&user1)
	// table2.Insert(&user2)

	// group1 := database.Group{
	// 	DN:   "cn=admins,ou=groups,dc=example,dc=nz",
	// 	Name: "Admins",
	// }
	// group2 := database.Group{
	// 	DN:   "cn=users,ou=groups,dc=example,dc=nz",
	// 	Name: "Users",
	// }
	// table3.Insert(&group1)
	// table3.Insert(&group2)

	// usergroup1 := database.UserGroup{
	// 	GroupId: group1.Id,
	// 	UserId:  user1.Id,
	// }
	// usergroup2 := database.UserGroup{
	// 	GroupId: group1.Id,
	// 	UserId:  user2.Id,
	// }
	// usergroup3 := database.UserGroup{
	// 	GroupId: group2.Id,
	// 	UserId:  user1.Id,
	// }

	// table4.Insert(&usergroup1)
	// table4.Insert(&usergroup2)
	// table4.Insert(&usergroup3)

	// //	table4.Delete(usergroup2)

	// table4.GetUser(user2.Id)

	table5 := database.NewApplicationsTable(db)
	table5.Drop()
	table5.Create()

	application := database.Application{
		TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	table5.Insert(&application)

	// fmt.Printf("%#v\n", application)

	// table5.Update(application)

	// table5.Delete(application)
	table5.GetTemplateID("140902edbcc424c09736af28ab2de604c3bde936")
}

func NewNoodle() Noodle {
	noodle := &NoodleImpl{}
	return noodle
}
