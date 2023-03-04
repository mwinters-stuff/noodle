package api_handlers_test

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	heimdall_mocks "github.com/mwinters-stuff/noodle/noodle/heimdall/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UploadIconHandlersTestSuite struct {
	suite.Suite
	mockHiemdall *heimdall_mocks.Heimdall
	api          *operations.NoodleAPI
	tempFile     *os.File
}

func (suite *UploadIconHandlersTestSuite) SetupSuite() {
	heimdall.Logger = log.Output(nil)
}

func (suite *UploadIconHandlersTestSuite) SetupTest() {

	suite.mockHiemdall = heimdall_mocks.NewHeimdall(suite.T())

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterUploadIconApiHandlers(suite.api, suite.mockHiemdall)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUploadIconHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleUploadIconHandler)
	suite.tempFile = nil
}

func (suite *UploadIconHandlersTestSuite) TearDownTest() {
	if suite.tempFile != nil {
		os.Remove(suite.tempFile.Name())
	}

}

func (suite *UploadIconHandlersTestSuite) TearDownSuite() {

}

func (suite *UploadIconHandlersTestSuite) TestHandlerUploadIconsGet() {
	suite.mockHiemdall.EXPECT().ListIcons().Once().Return([]string{
		"file1",
		"file2",
		"file3",
	}, nil)

	pr := models.Principal("")

	response := suite.api.NoodleAPIGetNoodleUploadIconHandler.Handle(noodle_api.NewGetNoodleUploadIconParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`["file1","file2","file3"]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UploadIconHandlersTestSuite) TestHandlerUploadIconsGetFailed() {
	suite.mockHiemdall.EXPECT().ListIcons().Once().Return([]string{}, errors.New("failed"))

	pr := models.Principal("")

	response := suite.api.NoodleAPIGetNoodleUploadIconHandler.Handle(noodle_api.NewGetNoodleUploadIconParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UploadIconHandlersTestSuite) TestHandlerUploadIconsPost() {
	suite.tempFile, _ = os.CreateTemp("/tmp", "afile")

	pr := models.Principal("")
	params := noodle_api.NewPostNoodleUploadIconParams()
	params.Upfile = io.NopCloser(suite.tempFile)

	params.HTTPRequest = &http.Request{
		MultipartForm: &multipart.Form{
			File: make(map[string][]*multipart.FileHeader),
		},
	}
	params.HTTPRequest.MultipartForm.File["upfile"] = append(params.HTTPRequest.MultipartForm.File["upfile"], &multipart.FileHeader{Filename: "afile"})

	suite.mockHiemdall.EXPECT().UploadIcon("afile", params.Upfile).Once().Return(nil)

	response := suite.api.NoodleAPIPostNoodleUploadIconHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UploadIconHandlersTestSuite) TestHandlerUploadIconsPostError() {
	suite.tempFile, _ = os.CreateTemp("/tmp", "afile")

	pr := models.Principal("")
	params := noodle_api.NewPostNoodleUploadIconParams()
	params.Upfile = io.NopCloser(suite.tempFile)

	params.HTTPRequest = &http.Request{
		MultipartForm: &multipart.Form{
			File: make(map[string][]*multipart.FileHeader),
		},
	}
	params.HTTPRequest.MultipartForm.File["upfile"] = append(params.HTTPRequest.MultipartForm.File["upfile"], &multipart.FileHeader{Filename: "afile"})

	suite.mockHiemdall.EXPECT().UploadIcon("afile", params.Upfile).Once().Return(errors.New("failed"))

	response := suite.api.NoodleAPIPostNoodleUploadIconHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestUploadIconHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UploadIconHandlersTestSuite))
}
