package yamltypes

import "gopkg.in/yaml.v3"

func UnmarshalConfig(data []byte) (AppConfig, error) {
	var r AppConfig
	err := yaml.Unmarshal(data, &r)
	return r, err
}

type AppConfig struct {
	Postgres struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
	} `yaml:"postgres"`
	Ldap struct {
		URL                      string `yaml:"url"`
		BaseDn                   string `yaml:"base_dn"`
		UsernameAttribute        string `yaml:"username_attribute"`
		UserFilter               string `yaml:"user_filter"`
		AllUsersFilter           string `yaml:"all_users_filter"`
		AllGroupsFilter          string `yaml:"all_groups_filter"`
		UserGroupsFilter         string `yaml:"user_groups_filter"`
		GroupUsersFilter         string `yaml:"group_users_filter"`
		GroupNameAttribute       string `yaml:"group_name_attribute"`
		UserDisplayNameAttribute string `yaml:"user_display_name_attribute"`
		User                     string `yaml:"user"`
		Password                 string `yaml:"password"`
		GroupMemberAttribute     string `yaml:"group_member_attribute"`
	} `yaml:"ldap"`
}
