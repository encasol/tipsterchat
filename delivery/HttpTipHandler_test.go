package delivery_test

import (
	"github.com/encasol/tipsterchat/delivery"
	"github.com/encasol/tipsterchat/model"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"

	"testing"
)

// Mock section
type MockedTipService struct {
	mock.Mock
}

func (mock MockedTipService) ListTips() ([]model.Tip, error) {
	args := mock.Called()
	return args.Get(0).([]model.Tip), args.Error(1)
}

func (mock MockedTipService) AddTip(tip model.Tip) error {
	args := mock.Called(tip)
	return args.Error(0)
}

type MockedResponseWriter struct {
	mock.Mock
}

func (mock MockedResponseWriter) Header() http.Header {
	args := mock.Called()
	return args.Get(0).(http.Header)
}

func (mock MockedResponseWriter) Write(b []byte) (int, error) {
	args := mock.Called(b)
	return args.Int(0), args.Error(1)
}

func (mock MockedResponseWriter) WriteHeader(statusCode int) {
	mock.Called(statusCode)
}

type MockedJsonProxy struct {
	mock.Mock
}

func (mock MockedJsonProxy) DecodeJson(Body io.ReadCloser) (model.Tip, error) {
	args := mock.Called(Body)
	return args.Get(0).(model.Tip), args.Error(1)
}

func (mock MockedJsonProxy) EncodeJson(tips []model.Tip) ([]byte, error) {
	args := mock.Called(tips)
	return args.Get(0).([]byte), args.Error(1)
}

// Test section
func TestListTipsOnlyGet(t *testing.T) {
	svc := new(MockedTipService)
	jsonMock := new(MockedJsonProxy)
	writer := new(MockedResponseWriter)

	handler := delivery.HttpTipHandler{
		TipService: svc,
		Json:       jsonMock,
	}

	bytesJs := []byte{1, 2, 3}
	emptyTips := []model.Tip{}
	header := make(http.Header)
	request := http.Request{
		Method: "GET",
	}

	svc.On("ListTips").Return(emptyTips, nil)
	jsonMock.On("EncodeJson", emptyTips).Return(bytesJs, nil)
	writer.On("Header").Return(header)
	writer.On("Write", bytesJs).Return(0, nil)

	handler.ListTipHandler(writer, &request)
}

func TestListTipsPost(t *testing.T) {
	svc := new(MockedTipService)
	jsonMock := new(MockedJsonProxy)
	writer := new(MockedResponseWriter)

	handler := delivery.HttpTipHandler{
		TipService: svc,
		Json:       jsonMock,
	}

	emptyTips := []model.Tip{}
	emptyTip := model.Tip{}
	bytesJs := []byte{1, 2, 3}
	header := make(http.Header)
	request := http.Request{
		Method: "POST",
	}

	svc.On("AddTip", emptyTip).Return(nil)
	svc.On("ListTips").Return(emptyTips, nil)
	jsonMock.On("DecodeJson", nil).Return(emptyTip, nil)
	jsonMock.On("EncodeJson", emptyTips).Return(bytesJs, nil)
	writer.On("Header").Return(header)
	writer.On("Write", bytesJs).Return(0, nil)

	handler.ListTipHandler(writer, &request)
}
