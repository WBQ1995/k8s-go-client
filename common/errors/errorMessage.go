package errors

type ErrorMessage string

const (
	// common error
	ErrorMsgHttpCode500InternalServerError ErrorMessage = "Internal error, please try again."

	//k8s error
	ErrorMsgK8sListNodes        = "An error occurs while listing k8s nodes."
	ErrorMsgK8sListPods         = "An error occurs while listing k8s pods."
	ErrorMsgK8sCreateDeployment = "An error occurs while creating a k8s deployment."
	ErrorMsgK8sDeleteDeployment = "An error occurs while deleting a k8s deployment."
	ErrorMsgK8sCreateService    = "An error occurs while creating a k8s service."
	ErrorMsgK8sDeleteService    = "An error occurs while deleting a k8s service."
)
