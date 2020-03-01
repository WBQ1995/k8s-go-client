package client

import (
	"github.com/gin-gonic/gin"
	"k8s-go-client/common/constants"
	"k8s-go-client/common/errors"
	"k8s-go-client/common/response"
	"k8s-go-client/logs"
	"k8s-go-client/vo"
	"net/http"
)

func ClientManager(router *gin.RouterGroup)  {
	router.GET(constants.TEST, testRouter)

	router.GET(constants.GET_PODS, getPodList)
	router.GET(constants.GET_NODES, getNodeList)

	router.GET(constants.DEPLOYMENT_DEMO, createDeploymentDemo)
	router.DELETE(constants.DEPLOYMENT_DEMO, deleteDeploymentDemo)

	router.POST(constants.CREATE_DEPLOYMENT, createDeployment)
	router.DELETE(constants.DELETE_DEPLOYMENT, deleteDeployment)

	router.POST(constants.CREATE_SERVICE, createService)
	router.DELETE(constants.DELETE_SERVICE, deleteService)
}

func testRouter(c *gin.Context)  {

	c.JSON(http.StatusOK, "k8s go client at your service.")
}

func getNodeList(c *gin.Context)  {

	nodes, err := getK8sNodeList()

	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sListNodes)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sListNodes))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse(nodes))
}

func getPodList(c *gin.Context)  {

	pods, err := getK8sPodList()

	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sListPods)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sListPods))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse(pods))
}

func createDeployment(c *gin.Context) {

	var deploymentConfig vo.DeploymentConfig
	err := c.BindJSON(&deploymentConfig)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, response.CreateErrorResponse(errors.ErrorCodeK8sCreateDeployment))
		return
	}

	result, err := createK8sDeployment(deploymentConfig)
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sCreateDeployment)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sCreateDeployment))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse(result))
}

func deleteDeployment(c *gin.Context)  {

	name := c.Param("name")
	err := deleteK8sDeployment(name)
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sDeleteDeployment)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sDeleteDeployment))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse("Delete deployment: " + name +  " successfully."))
}

func createService(c *gin.Context) {

	var serviceConfig vo.ServiceConfig
	err := c.BindJSON(&serviceConfig)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, response.CreateErrorResponse(errors.ErrorCodeK8sCreateService))
		return
	}

	result, err := createK8sService(serviceConfig)
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sCreateService)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sCreateService))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse(result))
}

func deleteService(c *gin.Context)  {

	name := c.Param("name")
	err := deleteK8sService(name)
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sDeleteService)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sDeleteService))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse("Delete service: " + name +  " successfully."))
}

func createDeploymentDemo(c *gin.Context)  {

	var config = vo.DeploymentConfig{
		Name:          "demo-deployment",
		Labels:        map[string]string {"app": "demo",},
		Replicas:      2,
		ContainerName: "demo-web",
		ImageName:     "nginx:1.12",
		ContainerPort: 80,
	}

	result, err := createK8sDeployment(config)
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sCreateDeployment)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sCreateDeployment))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse(result))
}

func deleteDeploymentDemo(c *gin.Context)  {

	err := deleteK8sDeployment("demo-deployment")
	if err != nil {
		logs.GetLogger().Error(errors.ErrorMsgK8sDeleteDeployment)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.CreateErrorResponse(errors.ErrorCodeK8sDeleteDeployment))
		return
	}
	c.JSON(http.StatusOK, response.CreateSuccessResponse("Delete deployment: demo-deployment successfully."))
}
