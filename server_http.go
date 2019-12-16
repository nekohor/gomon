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
				c.JSON(
					http.StatusNotFound,
					gin.H{
						"code": 1230,
						"error": "Invalid query",
					})
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": 1231,
						"exports": app.RespondCoil(req),
					})
			}

		})

		v1.GET("/ponds/stats/:coilId", func(c *gin.Context) {

			req := &StatsRequest{}

			req.CoilInfo.CoilId = c.Param("coilId")
			req.CoilInfo.CurDir = c.DefaultQuery("curDir", "")
			req.CoilInfo.FactorName = c.DefaultQuery("factorName", "")

			req.StatsOption.FunctionName = c.DefaultQuery("functionName", "")

			// stats options
			aim, err := strconv.ParseFloat(c.DefaultQuery("aim", "0"),32)
			CheckError(err)
			req.StatsOption.Aim = DataType(aim)

			tol, err := strconv.ParseFloat(c.DefaultQuery("tolerance", "0"), 32)
			CheckError(err)
			req.StatsOption.Tolerance = DataType(tol)

			upper, err := strconv.ParseFloat(c.DefaultQuery("upper", "0"), 32)
			CheckError(err)
			req.StatsOption.Upper = DataType(upper)

			lower, err := strconv.ParseFloat(c.DefaultQuery("lower", "0"), 32)
			CheckError(err)
			req.StatsOption.Lower = DataType(lower)

			req.StatsOption.Unit = c.DefaultQuery("unit", "")


			// length division
			req.LengthDivision.LengthName = c.DefaultQuery("lengthName", "")


			req.LengthDivision.HeadLen, err = strconv.Atoi(c.DefaultQuery("headLen", "0"))
			CheckError(err)
			req.LengthDivision.TailLen, err = strconv.Atoi(c.DefaultQuery("tailLen", "0"))
			CheckError(err)

			headPerc, err := strconv.ParseFloat(c.DefaultQuery("headPerc", "-100000000"), 32)
			CheckError(err)
			req.LengthDivision.HeadPerc = DataType(headPerc)

			tailPerc, err := strconv.ParseFloat(c.DefaultQuery("tailPerc", "-100000000"), 32)
			CheckError(err)
			req.LengthDivision.TailPerc = DataType(tailPerc)

			req.LengthDivision.HeadCut, err = strconv.Atoi(c.DefaultQuery("headCut", "5"))
			CheckError(err)
			req.LengthDivision.TailCut, err = strconv.Atoi(c.DefaultQuery("tailCut", "5"))
			CheckError(err)

			if req.CoilInfo.CurDir == "" || req.CoilInfo.FactorName == "" {
				c.JSON(
					http.StatusNotFound,
					gin.H{
						"code": 1230,
						"msg": "Invalid query",
					})
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": 1231,
						"msg": "query successfully",
						"stats": app.Stat(req),
					})
			}
		})
	}

	err := router.Run(app.Ctx.Cfg.GetPort())
	CheckError(err)
}
