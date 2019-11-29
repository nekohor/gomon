package gomon

import "C"
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RunHttpServer(app *Application) {


	router := gin.Default()

	router.GET("/pond", func(c *gin.Context) {

		req := c.Query("req")
		log.Println(req)
		c.JSON(http.StatusOK, app.RespondCoils(req))

	})

	router.Run(app.Ctx.Cfg.GetPort())
}