package errors

type APIErrorCode string

type IErrCode interface {
	GetMsgByErrCode() ErrorMessage
	GetErrorCode() APIErrorCode
}

const (
	HttpCode200OK                  = "200" //http.StatusOk
	HttpCode400BadRequest          = "400" //http.StatusBadRequest
	HttpCode401Unauthorized        = "401" //http.StatusUnauthorized
	HttpCode500InternalServerError = "500" //http.StatusInternalServerError

	ModuleCodeK8s           	   = "001" //k8s
)