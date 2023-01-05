package ldap_shim

import (
	"crypto/tls"

	ldap "github.com/go-ldap/ldap/v3"
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name LdapShim

type LdapShim interface {
	DialURL(addr string, opts ...ldap.DialOpt) error
	EscapeFilter(arg1 string) string
	NewSearchRequest(BaseDN string,
		Scope, DerefAliases, SizeLimit, TimeLimit int,
		TypesOnly bool,
		Filter string,
		Attributes []string,
		Controls []ldap.Control) *ldap.SearchRequest

	CloseConn()
	StartTLS(config *tls.Config) error
	Bind(username, password string) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
}

type LdapShimImpl struct {
	conn *ldap.Conn
}

// Bind implements LdapShim
func (s LdapShimImpl) Bind(username string, password string) error {
	return s.conn.Bind(username, password)
}

// CloseConn implements LdapShim
func (s LdapShimImpl) CloseConn() {
	s.conn.Close()
}

// Search implements LdapShim
func (s LdapShimImpl) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return s.conn.Search(searchRequest)
}

// StartTLS implements LdapShim
func (s LdapShimImpl) StartTLS(config *tls.Config) error {
	return s.conn.StartTLS(config)
}

func (s *LdapShimImpl) DialURL(addr string, opts ...ldap.DialOpt) error {
	var err error
	s.conn, err = ldap.DialURL(addr, opts...)
	return err
}

func (p *LdapShimImpl) EscapeFilter(arg1 string) string {
	return ldap.EscapeFilter(arg1)
}
func (p *LdapShimImpl) NewSearchRequest(BaseDN string,
	Scope, DerefAliases, SizeLimit, TimeLimit int,
	TypesOnly bool,
	Filter string,
	Attributes []string,
	Controls []ldap.Control) *ldap.SearchRequest {
	return ldap.NewSearchRequest(BaseDN, Scope, DerefAliases, SizeLimit, TimeLimit, TypesOnly, Filter, Attributes, Controls)
}

func NewLdapShim() LdapShim {
	return &LdapShimImpl{}
}
