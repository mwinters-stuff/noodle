package options_test

import (
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/options"
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
  user_filter: (&(objectClass=organizationalPerson)(uid=%s))
  all_users_filter: (objectclass=organizationalPerson)
  all_groups_filter: (objectclass=groupOfUniqueNames)
  user_groups_filter: (&(uniquemember={dn})(objectclass=groupOfUniqueNames))
  group_users_filter: (&(objectClass=groupOfUniqueNames)(cn=%s))
  group_name_attribute: cn
  user_display_name_attribute: displayName
  user: CN=readonly,DC=example,DC=com
  password: readonly
  group_member_attribute: uniqueMember
`
	data, err := options.UnmarshalOptions([]byte(yamltext))
	require.NoError(t, err)

	assert.Equal(t, "postgresuser", data.PostgresOptions.User)
	assert.Equal(t, "postgrespass", data.PostgresOptions.Password)
	assert.Equal(t, "postgres", data.PostgresOptions.Database)
	assert.Equal(t, "localhost", data.PostgresOptions.Hostname)
	assert.Equal(t, 5432, data.PostgresOptions.Port)

	assert.Equal(t, "ldap://example.com", data.LDAPOptions.URL)
	assert.Equal(t, "dc=example,dc=com", data.LDAPOptions.BaseDN)
	assert.Equal(t, "uid", data.LDAPOptions.UserNameAttribute)
	assert.Equal(t, "(&(objectClass=organizationalPerson)(uid=%s))", data.LDAPOptions.UserFilter)
	assert.Equal(t, "(objectclass=organizationalPerson)", data.LDAPOptions.AllUsersFilter)
	assert.Equal(t, "(objectclass=groupOfUniqueNames)", data.LDAPOptions.AllGroupsFilter)
	assert.Equal(t, "(&(uniquemember={dn})(objectclass=groupOfUniqueNames))", data.LDAPOptions.UserGroupsFilter)
	assert.Equal(t, "(&(objectClass=groupOfUniqueNames)(cn=%s))", data.LDAPOptions.GroupUsersFilter)
	assert.Equal(t, "cn", data.LDAPOptions.GroupNameAttribute)
	assert.Equal(t, "displayName", data.LDAPOptions.UserDisplayNameAttribute)
	assert.Equal(t, "CN=readonly,DC=example,DC=com", data.LDAPOptions.User)
	assert.Equal(t, "readonly", data.LDAPOptions.Password)
	assert.Equal(t, "uniqueMember", data.LDAPOptions.GroupMemberAttribute)
}
