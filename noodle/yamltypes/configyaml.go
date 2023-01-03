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
		URL                  string `yaml:"url"`
		BaseDn               string `yaml:"base_dn"`
		UsernameAttribute    string `yaml:"username_attribute"`
		AdditionalUsersDn    string `yaml:"additional_users_dn"`
		UsersFilter          string `yaml:"users_filter"`
		AdditionalGroupsDn   string `yaml:"additional_groups_dn"`
		GroupsFilter         string `yaml:"groups_filter"`
		GroupNameAttribute   string `yaml:"group_name_attribute"`
		DisplayNameAttribute string `yaml:"display_name_attribute"`
		User                 string `yaml:"user"`
		Password             string `yaml:"password"`
	} `yaml:"ldap"`
}
