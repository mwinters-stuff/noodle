package api_handlers_test

import (
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserHandlersTestSuite struct {
	suite.Suite
	mockDatabase  *mocks.Database
	mockTables    *mocks.Tables
	mockUserTable *mocks.UserTable
	api           *operations.NoodleAPI
}

func (suite *UserHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *UserHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterUserApiHandlers(suite.api, suite.mockDatabase)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUsersHandler)

}

func (suite *UserHandlersTestSuite) TearDownTest() {

}
func (suite *UserHandlersTestSuite) TearDownSuite() {

}

func (suite *UserHandlersTestSuite) TestHandlerUsersGetAll() {
	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
		{
			ID:          1,
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	pr := models.Principal("")

	response := suite.api.NoodleAPIGetNoodleUsersHandler.Handle(noodle_api.NewGetNoodleUsersParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	// args := mock.Arguments{}

	mockWriter.EXPECT().Write([]byte(`[{"DN":"CN=jack,DC=example,DC=nz","DisplayName":"Jack M","GivenName":"Jack","Id":2,"Surname":"M","UidNumber":1002,"Username":"jack"},{"DN":"CN=bob,DC=example,DC=nz","DisplayName":"bobextample","GivenName":"Bob","Id":1,"Surname":"Extample","UidNumber":1001,"Username":"bobe"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserHandlersTestSuite) TestHandlerUsersGetOne() {
	suite.mockUserTable.EXPECT().GetID(int64(2)).Return(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	},
		nil)

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUsersParams()
	var userid = int64(2)
	params.Userid = &userid

	response := suite.api.NoodleAPIGetNoodleUsersHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	// args := mock.Arguments{}

	mockWriter.EXPECT().Write([]byte(`[{"DN":"CN=jack,DC=example,DC=nz","DisplayName":"Jack M","GivenName":"Jack","Id":2,"Surname":"M","UidNumber":1002,"Username":"jack"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserHandlersTestSuite) TestHandlerUsersDBError() {
	suite.mockUserTable.EXPECT().GetID(int64(2)).Return(models.User{},
		errors.New("DB Failed"))

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUsersParams()
	var userid = int64(2)
	params.Userid = &userid

	response := suite.api.NoodleAPIGetNoodleUsersHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"DB Failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestUserHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlersTestSuite))
}
