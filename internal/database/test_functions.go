package database_test

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"

	"github.com/stretchr/testify/require"
)

func TestStepsRunner(t *testing.T, script *pgmock.Script) (net.Listener, string) {

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
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s", host, port)
	return listener, connStr
}

func SetupConnectionSteps(script *pgmock.Script) {
	QueryMock(script, "SELECT 1",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: nil,
			Parameters:           nil,
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

	fmt.Printf("msg => %#v\n", msg)

	switch p := msg.(type) {
	case *pgproto3.Parse:
		if strings.HasPrefix(p.Name, "stmtcache_") {
			p.Name = "stmtcache_?"
		}
	case *pgproto3.Describe:
		if strings.HasPrefix(p.Name, "stmtcache_") {
			p.Name = "stmtcache_?"
		}
	case *pgproto3.Bind:
		if strings.HasPrefix(p.PreparedStatement, "stmtcache_") {
			p.PreparedStatement = "stmtcache_?"
		}
	}

	if !reflect.DeepEqual(msg, e.want) {
		return fmt.Errorf("msg => %#v, e.want => %#v", msg, e.want)
	}

	return nil
}

func ExpectMessageX(want pgproto3.FrontendMessage) pgmock.Step {
	return &expectMessageStepX{want}
}

func QueryMock(script *pgmock.Script, statement string, bind pgproto3.Bind, fields []pgproto3.FieldDescription, row [][]byte) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ParameterDescription{ParameterOIDs: []uint32{}}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.RowDescription{Fields: fields}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{ObjectType: 'P', Name: ""}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.BindComplete{}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.DataRow{Values: row}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte("SELECT 1")}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}

func InsertUpdateDeleteMock(script *pgmock.Script, statement string, parameterDescription pgproto3.ParameterDescription, bind pgproto3.Bind) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&parameterDescription))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.NoData{}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "", ObjectType: 'P'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.BindComplete{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.NoData{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte("INSERT 0 1")}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}

func SelectMock(script *pgmock.Script, statement string, parameterDescription pgproto3.ParameterDescription, fields []pgproto3.FieldDescription, bind pgproto3.Bind, values [][]byte) {
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Parse{Name: "stmtcache_?", Query: statement, ParameterOIDs: nil}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "stmtcache_?", ObjectType: 'S'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ParseComplete{}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&parameterDescription))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	script.Steps = append(script.Steps, ExpectMessageX(&bind))

	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Describe{Name: "", ObjectType: 'P'}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Execute{Portal: "", MaxRows: 0}))
	script.Steps = append(script.Steps, ExpectMessageX(&pgproto3.Sync{}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.BindComplete{}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.RowDescription{Fields: fields}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.DataRow{Values: values}))

	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte(fmt.Sprintf("SELECT %d", len(values)))}))
	script.Steps = append(script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

}
