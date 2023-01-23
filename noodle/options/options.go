package options

import "gopkg.in/yaml.v3"

type NoodleOptions struct {
	Config string `short:"c" long:"config" description:"Noodle Configuration File"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`
	Drop   bool   `long:"drop" description:"Drop Database"`
}

type PostgresOptions struct {
	User     string `long:"pg-user" description:"Postgres User" yaml:"user"`
	Password string `long:"pg-password" description:"Postgres Password" yaml:"password"`
	Database string `long:"pg-database" description:"Postgres Database" yaml:"db"`
	Hostname string `long:"pg-hostname" description:"Postgres Hostname" yaml:"hostname"`
	Port     int    `long:"pg-port" description:"Postgres Port" default:"5432" yaml:"port"`
}

type LDAPOptions struct {
	URL                      string `long:"ldap-url" description:"LDAP Server URL" yaml:"url"`
	BaseDN                   string `long:"ldap-base-url" description:"LDAP BaseDN" yaml:"base_dn"`
	User                     string `long:"ldap-user" description:"LDAP User (ReadOnly)" `
	Password                 string `long:"ldap-password" description:"LDAP Password (ReadOnly)"`
	UserFilter               string `long:"ldap-user-filter" description:"LDAP User Filter" default:"(&(objectClass=organizationalPerson)(uid=%s))" yaml:"user_filter"`
	AllUsersFilter           string `long:"ldap-all-user-filter" description:"LDAP All User Filter" default:"(objectclass=organizationalPerson)" yaml:"all_users_filter"`
	AllGroupsFilter          string `long:"ldap-all-groups-filter" description:"LDAP All Groups Filter" default:"(objectclass=groupOfUniqueNames)" yaml:"all_groups_filter"`
	UserGroupsFilter         string `long:"ldap-user-groups-filter" description:"LDAP User Groups Filter" default:"(&(uniquemember=%s)(objectclass=groupOfUniqueNames))" yaml:"user_groups_filter"`
	GroupUsersFilter         string `long:"ldap-group-userss-filter" description:"LDAP Group Users Filter" default:"(&(objectClass=groupOfUniqueNames)(cn=%s))" yaml:"group_users_filter"`
	UserNameAttribute        string `long:"ldap-user-name-attribute" description:"LDAP User Name Attribute" default:"uid" yaml:"username_attribute"`
	GroupNameAttribute       string `long:"ldap-group-name-attribute" description:"LDAP Group Name Attribute" default:"cn" yaml:"group_name_attribute"`
	UserDisplayNameAttribute string `long:"ldap-user-display-name-attribute" description:"LDAP User Display Name Attribute" default:"cn" yaml:"user_display_name_attribute"`
	GroupMemberAttribute     string `long:"ldap-group-member-attribute" description:"LDAP Group Nember Attribute" default:"uniqueMember" yaml:"group_member_attribute"`
}

type AllNoodleOptions struct {
	NoodleOptions   NoodleOptions
	PostgresOptions PostgresOptions `yaml:"postgres"`
	LDAPOptions     LDAPOptions     `yaml:"ldap"`
}

func UnmarshalOptions(data []byte) (AllNoodleOptions, error) {
	var r AllNoodleOptions
	err := yaml.Unmarshal(data, &r)
	return r, err
}
