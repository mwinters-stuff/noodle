package database_test

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"

	"github.com/stretchr/testify/require"
)

func TestStepsRunner(t *testing.T, script *pgmock.Script) (net.Listener, yamltypes.AppConfig) {

	listener, err := net.Listen("tcp", "127.0.0.1:")
	require.NoError(t, err)

	serverErrChan := make(chan error, 1)
	go func() {
		defer close(serverErrChan)

		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			serverErrChan <- err
			return
		}
		defer conn.Close()

		err = conn.SetDeadline(time.Now().Add(time.Second))
		if err != nil {
			serverErrChan <- err
			return
		}

		err = script.Run(pgproto3.NewBackend(pgproto3.NewChunkReader(conn), conn))
		if err != nil {
			t.Log(err)
			serverErrChan <- err
			return
		}
	}()

	parts := strings.Split(listener.Addr().String(), ":")
	host := parts[0]
	port := parts[1]

	yamltext := fmt.Sprintf(`
postgres:
  user: postgresuser
  password: postgrespass
  db: postgres
  hostname: %s
  port: %s
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
`, host, port)

	config, _ := yamltypes.UnmarshalConfig([]byte(yamltext))

	return listener, config
}

type MsgType struct {
	Type string `json:"type"`
}

func LoadDatabaseSteps(t require.TestingT, script *pgmock.Script, steps []string) {
	for _, stepjson := range steps {
		frontend := stepjson[0] == 'F'
		stepjson = stepjson[2:]
		var msgType MsgType
		err := json.Unmarshal([]byte(stepjson), &msgType)
		require.NoError(t, err)

		startedwith := len(script.Steps)

		if frontend {
			switch msgType.Type {
			case "StartupMessage":
				unMarshalFrontendMessage(stepjson, &pgproto3.StartupMessage{}, t, script)
			case "Parse":
				parse := &pgproto3.Parse{}

				err := json.Unmarshal([]byte(stepjson), parse)
				require.NoError(t, err)

				if strings.HasPrefix(parse.Name, "stmtcache_") {
					parse.Name = "stmtcache_?"
				}
				script.Steps = append(script.Steps, ExpectMessageX(parse))
			case "Describe":
				describe := &pgproto3.Describe{}

				err := json.Unmarshal([]byte(stepjson), describe)
				require.NoError(t, err)

				if strings.HasPrefix(describe.Name, "stmtcache_") {
					describe.Name = "stmtcache_?"
				}
				script.Steps = append(script.Steps, ExpectMessageX(describe))
			case "Sync":
				unMarshalFrontendMessage(stepjson, &pgproto3.Sync{}, t, script)
			case "Bind":
				bind := &pgproto3.Bind{}

				err := json.Unmarshal([]byte(stepjson), bind)
				require.NoError(t, err)

				if strings.HasPrefix(bind.PreparedStatement, "stmtcache_") {
					bind.PreparedStatement = "stmtcache_?"
				}
				if bind.Parameters == nil || len(bind.Parameters) == 0 {
					bind.Parameters = [][]uint8{}
				}
				if bind.ParameterFormatCodes == nil || len(bind.ParameterFormatCodes) == 0 {
					bind.ParameterFormatCodes = []int16{}
				}
				fmt.Printf("BIND => %#v\n", bind)
				script.Steps = append(script.Steps, ExpectMessageX(bind))

			case "Execute":
				unMarshalFrontendMessage(stepjson, &pgproto3.Execute{}, t, script)
			}
		} else {
			switch msgType.Type {
			case "AuthenticationOK":
				unMarshalBackendMessage(stepjson, &pgproto3.AuthenticationOk{}, t, script)
			case "ParameterStatus":
				unMarshalBackendMessage(stepjson, &pgproto3.ParameterStatus{}, t, script)
			case "BackendKeyData":
				unMarshalBackendMessage(stepjson, &pgproto3.BackendKeyData{}, t, script)
			case "ParseComplete":
				unMarshalBackendMessage(stepjson, &pgproto3.ParseComplete{}, t, script)
			case "ParameterDescription":
				unMarshalBackendMessage(stepjson, &pgproto3.ParameterDescription{}, t, script)
			case "NoData":
				unMarshalBackendMessage(stepjson, &pgproto3.NoData{}, t, script)
			case "ReadyForQuery":
				unMarshalBackendMessage(stepjson, &pgproto3.ReadyForQuery{}, t, script)
			case "BindComplete":
				unMarshalBackendMessage(stepjson, &pgproto3.BindComplete{}, t, script)
			case "CommandComplete":
				unMarshalBackendMessage(stepjson, &pgproto3.CommandComplete{}, t, script)
			case "RowDescription":
				unMarshalBackendMessage(stepjson, &pgproto3.RowDescription{}, t, script)
			case "DataRow":
				unMarshalBackendMessage(stepjson, &pgproto3.DataRow{}, t, script)
			}
		}

		require.NotEqual(t, startedwith, script.Steps, "%s was not unmarshaled!")

	}
}

func unMarshalFrontendMessage(stepjson string, m pgproto3.FrontendMessage, t require.TestingT, script *pgmock.Script) {
	err := json.Unmarshal([]byte(stepjson), m)
	require.NoError(t, err)
	script.Steps = append(script.Steps, ExpectMessageX(m))
}

func unMarshalBackendMessage(stepjson string, m pgproto3.BackendMessage, t require.TestingT, script *pgmock.Script) {
	err := json.Unmarshal([]byte(stepjson), m)
	require.NoError(t, err)
	script.Steps = append(script.Steps, SendMessageX(m))
}

func SetupConnectionSteps(t require.TestingT, script *pgmock.Script) {
	QueryMock(script, "SELECT 1",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{},
			Parameters:           [][]byte{},
			ResultFormatCodes:    []int16{1},
		},
		[]pgproto3.FieldDescription{
			{
				Name:                 []byte("?column?"),
				TableOID:             0,
				TableAttributeNumber: 1,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
		[][]byte{[]byte("1")})

}

type expectMessageStepX struct {
	want pgproto3.FrontendMessage
}

func (e *expectMessageStepX) Step(backend *pgproto3.Backend) error {
	msg, err := backend.Receive()
	if err != nil {
		return err
	}

	received := fmt.Sprintf("%#v", msg)
	expected := fmt.Sprintf("%#v", e.want)

	switch p := msg.(type) {
	case *pgproto3.Parse:
		if strings.HasPrefix(p.Name, "stmtcache_") {
			p.Name = "stmtcache_?"
		}
		received = fmt.Sprintf("%#v", p)
	case *pgproto3.Describe:
		if strings.HasPrefix(p.Name, "stmtcache_") {
			p.Name = "stmtcache_?"
		}
		received = fmt.Sprintf("%#v", p)
	case *pgproto3.Bind:
		if strings.HasPrefix(p.PreparedStatement, "stmtcache_") {
			p.PreparedStatement = "stmtcache_?"
		}
		if p.Parameters == nil || len(p.Parameters) == 0 {
			p.Parameters = [][]uint8{}
		}
		if p.ParameterFormatCodes == nil || len(p.ParameterFormatCodes) == 0 {
			p.ParameterFormatCodes = []int16{}
		}
		received = fmt.Sprintf("%#v", p)

	}

	fmt.Printf("Expected => %s\nReceived => %s\n", expected, received)
	if expected != received {
		return fmt.Errorf("not equal")
	}

	return nil
}

func ExpectMessageX(want pgproto3.FrontendMessage) pgmock.Step {
	return &expectMessageStepX{want}
}

type sendMessageStepX struct {
	msg pgproto3.BackendMessage
}

func (e *sendMessageStepX) Step(backend *pgproto3.Backend) error {
	str, _ := json.Marshal(e.msg)
	fmt.Printf("Sending  => %s\n", str)
	return backend.Send(e.msg)
}

func SendMessageX(msg pgproto3.BackendMessage) pgmock.Step {
	return &sendMessageStepX{msg}
}

func QueryMock(script *pgmock.Script, statement string, bind pgproto3.Bind, fields []pgproto3.FieldDescription, row [][]byte) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ParameterDescription{ParameterOIDs: []uint32{}}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.RowDescription{Fields: fields}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{ObjectType: 'P', Name: ""}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.BindComplete{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.DataRow{Values: row}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.CommandComplete{CommandTag: []byte("SELECT 1")}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}

func InsertUpdateDeleteMock(script *pgmock.Script, statement string, parameterOIDs []uint32, bind pgproto3.Bind) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ParameterDescription{ParameterOIDs: parameterOIDs}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.NoData{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "", ObjectType: 'P'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.BindComplete{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.NoData{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.CommandComplete{CommandTag: []byte("INSERT 0 1")}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}

func SelectMock(script *pgmock.Script, statement string, parameterDescription pgproto3.ParameterDescription, fields []pgproto3.FieldDescription, bind pgproto3.Bind, values [][]byte) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, SendMessageX(&parameterDescription))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))

	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "", ObjectType: 'P'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.BindComplete{}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.DataRow{Values: values}))

	script.Steps = append(script.Steps, SendMessageX(&pgproto3.CommandComplete{CommandTag: []byte(fmt.Sprintf("SELECT %d", len(values)))}))
	script.Steps = append(script.Steps, SendMessageX(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}

func CreateAppTemplateTableSteps(t *testing.T, script *pgmock.Script) {
	LoadDatabaseSteps(t, script, []string{
		`F {"Type": "Query", "String": "CREATE TABLE IF NOT EXISTS application_template (\n  appid CHAR(40) PRIMARY KEY,\n  name VARCHAR(20) UNIQUE,\n  website VARCHAR(100) UNIQUE,\n  license VARCHAR(100),\n  description VARCHAR(1000),\n  enhanced BOOL,\n  tilebackground VARCHAR(256),\n  icon VARCHAR(256), \n  sha CHAR(40)\n)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE TABLE"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
		`F {"Type": "Query", "String": "CREATE INDEX IF NOT EXISTS application_template_idx1 ON application_template(name)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE INDEX"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
	})
}

func CreateUserTableSteps(t *testing.T, script *pgmock.Script) {
	LoadDatabaseSteps(t, script, []string{
		`F {"Type": "Query", "String": "CREATE TABLE IF NOT EXISTS users (\n  id SERIAL PRIMARY KEY,\n  username VARCHAR(50) UNIQUE,\n  dn VARCHAR(200) UNIQUE,\n  displayname VARCHAR(100),\n  givenname VARCHAR(100),\n  surname VARCHAR(100),\n  uidnumber INTEGER\n)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE TABLE"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
	})
}
