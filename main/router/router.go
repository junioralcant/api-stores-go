package router

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/main/router/userR"
)

func Initialize() {
	apiPrefix := "/api/v1"

	r := gin.Default()

	userR.InitUserRoutes(r, apiPrefix)

	r.Run(":3333")
}
