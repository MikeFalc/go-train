package storage

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

func GetStorageListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, get())
}

func AddStorageNodeHandler(c *gin.Context) {
	storageItem, statusCode, err := convertHTTPBodyToStorage(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": add(storageItem.Mip, storageItem.Description)})
}

func DeleteStorageNodeHandler(c *gin.Context) {
	storageID := c.Param("id")
	if err := delete(storageID); err != nil {
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK, "")
}

//func DeleteStorageNodeByMipHandler(c *gin.Context) {
//	storageMip := c.Params("mip")
//	if err := deleteByMip(storageMip); err != nil {
//		c.JSON(http.StatusInternalServerError,err)
//	}
//	c.JSON(http.StatusOK, "")
//}

func convertHTTPBodyToStorage(httpBody io.ReadCloser) (StorageNode, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return StorageNode{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var storageItem StorageNode
	err = json.Unmarshal(body, &storageItem)
	if err != nil {
		return StorageNode{}, http.StatusBadRequest, err
	}
	return storageItem, http.StatusOK, nil
}