package httpresponse

import (
	"time"

	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
)

type httpErrorResponse struct {
	RequestId string `json:"request_id"`
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func NewHttpErrorResponse(reqId, err string) *httpErrorResponse {
	return &httpErrorResponse{
		RequestId: reqId,
		IsSuccess: false,
		Message:   err,
		Timestamp: time.Now().Unix(),
	}
}

type httpSuccessResponse struct {
	RequestId    string      `json:"request_id"`
	IsSuccess    bool        `json:"is_success"`
	Message      string      `json:"message"`
	Timestamp    int64       `json:"timestamp"`
	ResponseData interface{} `json:"response_data"`
}

func NewHttpSuccessResponse(reqId string, res interface{}) *httpSuccessResponse {
	return &httpSuccessResponse{
		RequestId:    reqId,
		IsSuccess:    true,
		Message:      constants.SUCCESS,
		Timestamp:    time.Now().Unix(),
		ResponseData: res,
	}
}

func NewHttpSuccessResponseWithMessage(reqId string, res interface{}, msg string) *httpSuccessResponse {
	response := NewHttpSuccessResponse(reqId, res)
	response.Message = msg
	return response
}
