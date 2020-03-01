package errors

const (
	ErrorCodeK8sListNodes			K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "001"
	ErrorCodeK8sListPods			K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "002"
	ErrorCodeK8sCreateDeployment	K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "003"
	ErrorCodeK8sDeleteDeployment	K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "004"
	ErrorCodeK8sCreateService		K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "005"
	ErrorCodeK8sDeleteService		K8sErrorCode = HttpCode500InternalServerError + ModuleCodeK8s + "006"
)

type K8sErrorCode  APIErrorCode

func (e K8sErrorCode) GetMsgByErrCode() ErrorMessage {
	apiErr, ok := k8sErrorCodes[e]
	if !ok {
		return k8sErrorCodes[HttpCode500InternalServerError]
	}
	return apiErr
}

func (e K8sErrorCode) GetErrorCode() APIErrorCode  {
	return APIErrorCode(e)
}

type k8sErrorCodeMap map[K8sErrorCode]ErrorMessage

var k8sErrorCodes  = k8sErrorCodeMap {
	HttpCode500InternalServerError: 	ErrorMsgHttpCode500InternalServerError,
	ErrorCodeK8sListPods:				ErrorMsgK8sListPods,
	ErrorMsgK8sListNodes:				ErrorMsgK8sListNodes,
	ErrorMsgK8sCreateDeployment:		ErrorMsgK8sCreateDeployment,
	ErrorMsgK8sDeleteDeployment:		ErrorMsgK8sDeleteDeployment,
	ErrorMsgK8sCreateService:			ErrorMsgK8sCreateService,
	ErrorMsgK8sDeleteService:			ErrorMsgK8sDeleteService,
}
