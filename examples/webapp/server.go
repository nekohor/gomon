package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nekohor/gomon"
	"net/http"
	// "log"
	"fmt"
	"strings"
)

func main() {

	// curDir := `C:\Users\Crystal\go\src\github.com\nekohor\gomonitor\app\sample`
	// coilList := "M18122793Y"
	gm := gomon.NewGoMonitor()

	router := gin.Default()
	router.Use(Cors())
	router.LoadHTMLGlob(gomon.GetExeDir() + "/templates/*")
	router.Static("/static", gomon.GetExeDir()+"/static")

	// router
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Pond Show",
		})
		c.Abort()
	})

	router.GET("/api/curCoilIds", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"coilIdList": gm.Context.CoilIds,
			"seriesList": gm.Context.FactorConf.GetSeriesNames(),
		})
		c.Abort()
	})

	router.GET("/api/coil/:coilId", func(c *gin.Context) {
		coilId := c.Param("coilId")
		if coilId == "default" {
			coilId = gm.Context.CoilIds[0]
		}
		c.JSON(http.StatusOK, gin.H{
			"curCoilId": coilId,
			"coil":      gm.GetCoil(coilId),
		})
		c.Abort()
	})

	router.GET("/api/coils/:coilIdMessage", func(c *gin.Context) {
		coilIdMessage := c.Param("coilIdMessage")
		coilIdList := strings.Split(coilIdMessage, "&")
		c.JSON(http.StatusOK, gin.H{
			"coilIdList": coilIdList,
			"coils":      gm.GetCoils,
		})
		c.Abort()
	})

	router.Run(":9000") // listen and serve
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}
