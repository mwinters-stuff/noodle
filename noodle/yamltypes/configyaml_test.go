package yamltypes_test

import (
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeString(t *testing.T) {
	yamltext := `
postgres:
  user: postgresuser
  password: postgrespass
  db: postgres
  hostname: localhost
  port: 5432
ldap:
  url: ldap://example.com
  base_dn: dc=example,dc=com
  username_attribute: uid
  additional_users_dn: ou=people
  users_filter: (&({username_attribute}={input})(objectClass=person))
  additional_groups_dn: ou=groups
  groups_filter: (&(uniquemember={dn})(objectclass=groupOfUniqueNames))
  group_name_attribute: cn
  display_name_attribute: displayName
  user: CN=readonly,DC=example,DC=com
  password: readonly
`
	data, err := yamltypes.UnmarshalConfig([]byte(yamltext))
	require.NoError(t, err)

	assert.Equal(t, "postgresuser", data.Postgres.User)
	assert.Equal(t, "postgrespass", data.Postgres.Password)
	assert.Equal(t, "postgres", data.Postgres.Db)
	assert.Equal(t, "localhost", data.Postgres.Hostname)
	assert.Equal(t, 5432, data.Postgres.Port)

	assert.Equal(t, "ldap://example.com", data.Ldap.URL)
	assert.Equal(t, "dc=example,dc=com", data.Ldap.BaseDn)
	assert.Equal(t, "uid", data.Ldap.UsernameAttribute)
	assert.Equal(t, "ou=people", data.Ldap.AdditionalUsersDn)
	assert.Equal(t, "(&({username_attribute}={input})(objectClass=person))", data.Ldap.UsersFilter)
	assert.Equal(t, "ou=groups", data.Ldap.AdditionalGroupsDn)
	assert.Equal(t, "(&(uniquemember={dn})(objectclass=groupOfUniqueNames))", data.Ldap.GroupsFilter)
	assert.Equal(t, "cn", data.Ldap.GroupNameAttribute)
	assert.Equal(t, "displayName", data.Ldap.DisplayNameAttribute)
	assert.Equal(t, "CN=readonly,DC=example,DC=com", data.Ldap.User)
	assert.Equal(t, "readonly", data.Ldap.Password)
}
