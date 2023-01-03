package ldap

import "github.com/mwinters-stuff/noodle/noodle/yamltypes"

type LdapAuth interface {
}

type LdapAuthImpl struct {
	appConfig yamltypes.AppConfig
}

func NewLdapAuthImpl(appConfig yamltypes.AppConfig) LdapAuth {
	return &LdapAuthImpl{
		appConfig: appConfig,
	}
}
