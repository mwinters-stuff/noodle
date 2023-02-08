package options

type NoodleOptions struct {
	Debug               bool   `short:"d" long:"debug" description:"Debug Information"  env:"NOODLE_DEBUG"`
	Drop                bool   `long:"drop" description:"Drop Database"`
	IconSavePath        string `long:"image-save-path" description:"Location to save images to"  env:"NOODLE_ICON_SAVE_PATH"`
	WebClientPath       string `long:"web-client-path" description:"Location to find web client"  env:"NOODLE_WEB_CLIENT_PATH"`
	HeimdallListJsonURL string `long:"heimdall-list-json-url" description:"Location for list.json" default:"https://appslist.heimdall.site/list.json" env:"NOODLE_HEIMDALL_LIST_JSON_URL"`
	HeimdallIconBaseURL string `long:"heimdall-icon-base-url" description:"Base Location for Icons" default:"https://appslist.heimdall.site/icons" env:"NOODLE_HEIMDALL_ICON_BASE_URL"`
}

type PostgresOptions struct {
	User     string `long:"pg-user" description:"Postgres User"  env:"NOODLE_POSTGRES_USER"`
	Password string `long:"pg-password" description:"Postgres Password"  env:"NOODLE_POSTGRES_PASSWORD"`
	Database string `long:"pg-database" description:"Postgres Database"  env:"NOODLE_POSTGRES_DB"`
	Hostname string `long:"pg-hostname" description:"Postgres Hostname"  env:"NOODLE_POSTGRES_HOSTNAME"`
	Port     int    `long:"pg-port" description:"Postgres Port" default:"5432"  env:"NOODLE_POSTGRES_PORT"`
}

type LDAPOptions struct {
	URL                      string `long:"ldap-url" description:"LDAP Server URL"  env:"NOODLE_LDAP_URL"`
	BaseDN                   string `long:"ldap-base-url" description:"LDAP BaseDN"  env:"NOODLE_LDAP_BASE_DN"`
	User                     string `long:"ldap-user" description:"LDAP User (ReadOnly)" env:"NOODLE_LDAP_USER"`
	Password                 string `long:"ldap-password" description:"LDAP Password (ReadOnly)" env:"NOODLE_LDAP_PASSWORD"`
	UserFilter               string `long:"ldap-user-filter" description:"LDAP User Filter" default:"(&(objectClass=organizationalPerson)(uid=%s))"  env:"NOODLE_LDAP_USER_FILTER"`
	AllUsersFilter           string `long:"ldap-all-user-filter" description:"LDAP All User Filter" default:"(objectclass=organizationalPerson)"  env:"NOODLE_LDAP_ALL_USERS_FILTER"`
	AllGroupsFilter          string `long:"ldap-all-groups-filter" description:"LDAP All Groups Filter" default:"(objectclass=groupOfUniqueNames)"   env:"NOODLE_LDAP_ALL_GROUPS_FILTER"`
	UserGroupsFilter         string `long:"ldap-user-groups-filter" description:"LDAP User Groups Filter" default:"(&(uniquemember=%s)(objectclass=groupOfUniqueNames))"  env:"NOODLE_LDAP_USER_GROUPS_FILTER"`
	GroupUsersFilter         string `long:"ldap-group-userss-filter" description:"LDAP Group Users Filter" default:"(&(objectClass=groupOfUniqueNames)(cn=%s))"  env:"NOODLE_LDAP_GROUP_USERS_FILTER"`
	UserNameAttribute        string `long:"ldap-user-name-attribute" description:"LDAP User Name Attribute" default:"uid"  env:"NOODLE_LDAP_USERNAME_ATTRIBUTE"`
	GroupNameAttribute       string `long:"ldap-group-name-attribute" description:"LDAP Group Name Attribute" default:"cn"  env:"NOODLE_LDAP_GROUP_NAME_ATTRIBUTE"`
	UserDisplayNameAttribute string `long:"ldap-user-display-name-attribute" description:"LDAP User Display Name Attribute" default:"cn"  env:"NOODLE_LDAP_USER_DISPLAY_NAME_ATTRIBUTE"`
	GroupMemberAttribute     string `long:"ldap-group-member-attribute" description:"LDAP Group Nember Attribute" default:"uniqueMember"  env:"NOODLE_LDAP_GROUP_MEMBER_ATTRIBUTE"`
}

type AllNoodleOptions struct {
	NoodleOptions   NoodleOptions   `group:"Noodle Options"`
	PostgresOptions PostgresOptions `group:"Postgres"`
	LDAPOptions     LDAPOptions     `group:"LDAP"`
}
