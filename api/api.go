package api

import (
	"websitehimaif/config"
	"websitehimaif/router"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	db := config.ConnectDB()

	router.Router(r, db)
	r.Run()
}
