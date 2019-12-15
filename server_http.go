package gomon

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RunHttpServer(app *Application) {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ponds/exports/:coilId", func(c *gin.Context) {

			req := &coilRequest{}
			req.CoilId = c.Param("coilId")
			req.CurDir = c.DefaultQuery("curDir", "")

			factorsQuery := c.DefaultQuery("factorNames","")
			req.FatcorNames = strings.Split(factorsQuery, ",")

			if req.CurDir == "" || factorsQuery == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Invalid query"})
			}

			c.JSON(
				http.StatusOK,
				gin.H{
					"exports": app.RespondCoil(req),
				})
		})

		v1.GET("/ponds/stats/:coilId/:factor", func(c *gin.Context) {

			//
			//if curDir == "" || factorsQuery == "" {
			//	c.JSON(http.StatusNotFound, gin.H{"error": "Invalid query"})
			//}
			//
			//c.JSON(
			//	http.StatusOK,
			//	gin.H{
			//		"stats": app.RespondCoil(coilId, curDir, factorNames),
			//	})
		})
	}

	router.Run(app.Ctx.Cfg.GetPort())
}
