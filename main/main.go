package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"k8s-go-client/client"
	"k8s-go-client/common/constants"
	"k8s-go-client/configs"
	"k8s-go-client/logs"
	"time"
)

func main()  {

	conf := config.GetConfig()
	client.K8sInit()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		ValidateHeaders: false,
		Origins:         "*",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "GET, PUT, POST, DELETE",
		Methods:         "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
	}))

	v1 := r.Group("/api/v1")
	client.ClientManager(v1.Group(constants.K8S_CLIENT_PREFIX))

	err := r.Run(":" + conf.Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}