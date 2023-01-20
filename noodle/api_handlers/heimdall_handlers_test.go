package api_handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	heimdall_mocks "github.com/mwinters-stuff/noodle/noodle/heimdall/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HiemdallHandlersTest struct {
	suite.Suite
	mockHiemdall *heimdall_mocks.Heimdall
	mockDatabase *mocks.Database
}

func (suite *HiemdallHandlersTest) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
	heimdall.Logger = log.Output(nil)
	api_handlers.Logger = log.Output(nil)
}

func (suite *HiemdallHandlersTest) SetupTest() {
	suite.mockHiemdall = heimdall_mocks.NewHeimdall(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())

}

func (suite *HiemdallHandlersTest) TearDownTest() {

}
func (suite *HiemdallHandlersTest) TearSuite() {

}

func (suite *HiemdallHandlersTest) TestHandlerRefreshNoError() {
	suite.mockHiemdall.EXPECT().UpdateFromServer().Once().Return(nil)

	pr := models.Principal("")

	response := api_handlers.HandleHeimdallRefresh(suite.mockDatabase, suite.mockHiemdall, noodle_api.NewGetNoodleHeimdallReloadParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())

	mockWriter.EXPECT().Header().Once().Return(http.Header{})
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *HiemdallHandlersTest) TestHandlerRefreshError() {
	suite.mockHiemdall.EXPECT().UpdateFromServer().Once().Return(errors.New("failed"))

	pr := models.Principal("")

	response := api_handlers.HandleHeimdallRefresh(suite.mockDatabase, suite.mockHiemdall, noodle_api.NewGetNoodleHeimdallReloadParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())

	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestHiemdallHandlersTest(t *testing.T) {
	suite.Run(t, new(HiemdallHandlersTest))
}
