package ldap_test

import (
	"crypto/tls"
	"fmt"
	"log"
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestCheck(t *testing.T) {
	// The username and password we want to check
	username := "jessica"
	// password := "harperismydog"

	bindusername := "cn=readonly,dc=winters,dc=nz"
	bindpassword := "readonly"

	l, err := ldap.DialURL("ldap://192.168.30.23:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	// First bind with a read only user
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		"dc=winters,dc=nz",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN
	log.Print(sr.Entries[0])
	// // Bind as the user to verify their password
	// err = l.Bind(userdn, password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Rebind as the read only user for any further queries
	// err = l.Bind(bindusername, bindpassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	searchRequest = ldap.NewSearchRequest(
		"dc=winters,dc=nz", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(uniquemember=%s)(objectclass=groupOfUniqueNames))", ldap.EscapeFilter(userdn)), // The filter to apply
		[]string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err = l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range sr.Entries {
		log.Print(e)
	}

}
