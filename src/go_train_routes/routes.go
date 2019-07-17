package go_train_routes

import (
	"path"
	"path/filepath"
	"storage"
	"github.com/gin-gonic/gin"
)

func GetAllRoutes(){
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	storage.GetStorageRoutes(*r)
	//compute.GetComputeRoutes()


	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}

}
