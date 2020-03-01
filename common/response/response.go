package response

import (
	"fmt"
	"k8s-go-client/common/constants"
	"k8s-go-client/common/errors"
	"net/http"
)

type BasicResponse struct {
	Status   string              `json:"status"`
	Code     errors.APIErrorCode `json:"code"`
	Data     interface{}         `json:"data,omitempty"`
	Message  errors.ErrorMessage `json:"message,omitempty"`
	PageInfo *PageInfo           `json:"page_info,omitempty"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

type MixedResponse struct {
	BasicResponse
	MixData struct {
		Success interface{} `json:"success"`
		Fail    interface{} `json:"fail"`
	} `json:"mix_data"`
}

func NewSuccessResponseWithPageInfo(_data interface{}, _page *PageInfo) BasicResponse {
	return BasicResponse{
		Status:   constants.HttpStatusSuccess,
		Code:     errors.APIErrorCode(fmt.Sprint(http.StatusOK)),
		Data:     _data,
		PageInfo: _page,
	}
}

func CreateErrorResponse(_err errors.IErrCode) BasicResponse {
	errMessage := _err.GetMsgByErrCode()
	return BasicResponse{
		Status:  constants.HTTP_STATUS_ERROR,
		Code:    _err.GetErrorCode(),
		Message: errMessage,
	}
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: constants.HttpStatusSuccess,
		Code:   errors.HttpCode200OK,
		Data:   _data,
	}
}

func NewMixedResponse(_success, _fail interface{}) MixedResponse {

	return MixedResponse{
		BasicResponse: BasicResponse{
			Status: constants.HttpStatusMix,
		},
		MixData: struct {
			Success interface{} `json:"success"`
			Fail    interface{} `json:"fail"`
		}{_success, _fail},
	}
}

func NewUnauthorizedResponse() (int, BasicResponse) {
	return http.StatusUnauthorized, BasicResponse{
		Status:  constants.HTTP_STATUS_ERROR,
		Code:    errors.HttpCode401Unauthorized,
		Message: "Authorization failed.",
	}
}