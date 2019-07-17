package storage

import "github.com/gin-gonic/gin"

func GetStorageRoutes(r gin.Engine){
	r.GET("/storage", GetStorageListHandler)
	r.POST("/storage", AddStorageNodeHandler)
	r.DELETE("/storage/:id", DeleteStorageNodeHandler)
	//r.DELETE("/storage/:mip", DeleteStorageNodeByMipHandler)
}

