package database_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	"github.com/mwinters-stuff/noodle/noodle/options"

	"github.com/stretchr/testify/require"
)

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
	// fmt.Printf("Expected => %s\nReceived => %s\n", expected, received)
	if expected != received {
		return fmt.Errorf("expected => %s\nreceived => %s", expected, received)
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
	// str, _ := json.Marshal(e.msg)
	// fmt.Printf("Sending  => %s\n", str)
	return backend.Send(e.msg)
}

func SendMessageX(msg pgproto3.BackendMessage) pgmock.Step {
	return &sendMessageStepX{msg}
}

type TestFunctions struct {
}

func (i *TestFunctions) TestStepsRunner(t *testing.T, script *pgmock.Script) (net.Listener, options.AllNoodleOptions) {

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
	port, _ := strconv.Atoi(parts[1])

	config := options.AllNoodleOptions{
		NoodleOptions: options.NoodleOptions{
			Debug:               false,
			Drop:                false,
			IconSavePath:        "",
			WebClientPath:       "",
			HeimdallListJsonURL: "",
			HeimdallIconBaseURL: "",
		},
		PostgresOptions: options.PostgresOptions{
			User:     "postgresuser",
			Password: "postgrespass",
			Database: "postgres",
			Hostname: host,
			Port:     port,
		},
		LDAPOptions: options.LDAPOptions{
			URL:                      "ldap://example.com",
			BaseDN:                   "dc=example,dc=com",
			User:                     "CN=readonly,DC=example,DC=com",
			Password:                 "readonly",
			UserFilter:               "(&({username_attribute}={input})(objectClass=person))",
			AllUsersFilter:           "(objectclass=organizationalPerson)",
			AllGroupsFilter:          "(objectclass=groupOfUniqueNames)",
			UserGroupsFilter:         "(&(uniquemember=%s)(objectclass=groupOfUniqueNames))",
			GroupUsersFilter:         "(&(objectClass=groupOfUniqueNames)(cn=%s))",
			UserNameAttribute:        "uid",
			GroupNameAttribute:       "cn",
			UserDisplayNameAttribute: "cn",
			GroupMemberAttribute:     "uniqueMember",
		},
	}

	return listener, config
}

type MsgType struct {
	Type string `json:"type"`
}

func (i *TestFunctions) LoadDatabaseSteps(t require.TestingT, script *pgmock.Script, steps []string) {
	for _, stepjson := range steps {
		frontend := stepjson[0] == 'F'
		stepjson = stepjson[2:]
		var msgType MsgType
		// fmt.Println(stepjson)
		err := json.Unmarshal([]byte(stepjson), &msgType)
		require.NoError(t, err)

		startedwith := len(script.Steps)

		if frontend {
			switch msgType.Type {
			case "StartupMessage":
				i.unMarshalFrontendMessage(stepjson, &pgproto3.StartupMessage{}, t, script)
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
				i.unMarshalFrontendMessage(stepjson, &pgproto3.Sync{}, t, script)
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
				// fmt.Printf("BIND => %#v\n", bind)
				script.Steps = append(script.Steps, ExpectMessageX(bind))

			case "Execute":
				i.unMarshalFrontendMessage(stepjson, &pgproto3.Execute{}, t, script)
			case "Query":
				i.unMarshalFrontendMessage(stepjson, &pgproto3.Query{}, t, script)
			default:
				log.Fatalf("Missing process for frontend %s", msgType.Type)
			}
		} else {
			switch msgType.Type {
			case "AuthenticationOK":
				i.unMarshalBackendMessage(stepjson, &pgproto3.AuthenticationOk{}, t, script)
			case "ParameterStatus":
				i.unMarshalBackendMessage(stepjson, &pgproto3.ParameterStatus{}, t, script)
			case "BackendKeyData":
				i.unMarshalBackendMessage(stepjson, &pgproto3.BackendKeyData{}, t, script)
			case "ParseComplete":
				i.unMarshalBackendMessage(stepjson, &pgproto3.ParseComplete{}, t, script)
			case "ParameterDescription":
				i.unMarshalBackendMessage(stepjson, &pgproto3.ParameterDescription{}, t, script)
			case "NoData":
				i.unMarshalBackendMessage(stepjson, &pgproto3.NoData{}, t, script)
			case "ReadyForQuery":
				i.unMarshalBackendMessage(stepjson, &pgproto3.ReadyForQuery{}, t, script)
			case "BindComplete":
				i.unMarshalBackendMessage(stepjson, &pgproto3.BindComplete{}, t, script)
			case "CommandComplete":
				i.unMarshalBackendMessage(stepjson, &pgproto3.CommandComplete{}, t, script)
			case "RowDescription":
				i.unMarshalBackendMessage(stepjson, &pgproto3.RowDescription{}, t, script)
			case "DataRow":
				i.unMarshalBackendMessage(stepjson, &pgproto3.DataRow{}, t, script)
			case "ErrorResponse":
				i.unMarshalBackendMessage(stepjson, &pgproto3.ErrorResponse{}, t, script)
			default:
				log.Fatalf("Missing process for backend %s", msgType.Type)
			}

		}

		require.NotEqual(t, startedwith, len(script.Steps), "%s was not unmarshaled!", msgType.Type)

	}
}

func (i *TestFunctions) unMarshalFrontendMessage(stepjson string, m pgproto3.FrontendMessage, t require.TestingT, script *pgmock.Script) {
	err := json.Unmarshal([]byte(stepjson), m)
	require.NoError(t, err)
	script.Steps = append(script.Steps, ExpectMessageX(m))
}

func (i *TestFunctions) unMarshalBackendMessage(stepjson string, m pgproto3.BackendMessage, t require.TestingT, script *pgmock.Script) {
	err := json.Unmarshal([]byte(stepjson), m)
	require.NoError(t, err)
	script.Steps = append(script.Steps, SendMessageX(m))
}

func (i *TestFunctions) SetupConnectionSteps(t require.TestingT, script *pgmock.Script) {
	i.LoadDatabaseSteps(t, script, []string{
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Parse","Name":"stmtcache_1","Query":"SELECT 1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_1"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"?column?","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_1","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"?column?","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
}
