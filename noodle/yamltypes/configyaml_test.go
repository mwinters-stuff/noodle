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
	assert.Equal(t, "(&(objectClass=organizationalPerson)(uid=%s))", data.Ldap.UserFilter)
	// assert.Equal(t, "(&(objectClass=groupOfUniqueNames)(cn=%s))", data.Ldap.GroupFilter)
	assert.Equal(t, "(objectclass=organizationalPerson)", data.Ldap.AllUsersFilter)
	assert.Equal(t, "(objectclass=groupOfUniqueNames)", data.Ldap.AllGroupsFilter)
	assert.Equal(t, "(&(uniquemember={dn})(objectclass=groupOfUniqueNames))", data.Ldap.UserGroupsFilter)
	assert.Equal(t, "(&(objectClass=groupOfUniqueNames)(cn=%s))", data.Ldap.GroupUsersFilter)
	assert.Equal(t, "cn", data.Ldap.GroupNameAttribute)
	assert.Equal(t, "displayName", data.Ldap.UserDisplayNameAttribute)
	assert.Equal(t, "CN=readonly,DC=example,DC=com", data.Ldap.User)
	assert.Equal(t, "readonly", data.Ldap.Password)
	assert.Equal(t, "uniqueMember", data.Ldap.GroupMemberAttribute)
}
