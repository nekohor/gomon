package gomon

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

		v1.GET("/ponds/stats/:coilId", func(c *gin.Context) {

			req := &StatsRequest{}

			req.CoilInfo.CoilId = c.Param("coilId")
			req.CoilInfo.CurDir = c.DefaultQuery("curDir", "")
			req.CoilInfo.FactorName = c.DefaultQuery("factorName", "")

			req.StatsOption.FunctionName = c.DefaultQuery("functionName", "")
			req.StatsOption.Aim = StringToFloat32(c.DefaultQuery("aim", ""))
			req.StatsOption.Tolerance = StringToFloat32(c.DefaultQuery("tolerance", ""))
			req.StatsOption.Upper = StringToFloat32(c.DefaultQuery("upper", ""))
			req.StatsOption.Lower = StringToFloat32(c.DefaultQuery("lower", ""))
			req.StatsOption.Unit = c.DefaultQuery("unit", "")

			req.LengthDivision.LengthName = c.DefaultQuery("lengthName", "")

			var err error
			req.LengthDivision.HeadLen, err = strconv.Atoi(c.DefaultQuery("headLen", ""))
			CheckError(err)
			req.LengthDivision.TailLen, err = strconv.Atoi(c.DefaultQuery("tailLen", ""))
			CheckError(err)

			req.LengthDivision.HeadPerc = StringToFloat32(c.DefaultQuery("headPerc", "-1"))
			req.LengthDivision.TailPerc = StringToFloat32(c.DefaultQuery("tailPerc", "-1"))

			req.LengthDivision.HeadCut, err = strconv.Atoi(c.DefaultQuery("headCut", ""))
			CheckError(err)
			req.LengthDivision.TailCut, err = strconv.Atoi(c.DefaultQuery("tailCut", ""))
			CheckError(err)

			if req.CoilInfo.CurDir == "" || req.CoilInfo.FactorName == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Invalid query"})
			}

			c.JSON(
				http.StatusOK,
				gin.H{
					"stats": app.Stat(req),
				})

		})
	}

	err := router.Run(app.Ctx.Cfg.GetPort())
	CheckError(err)
}
