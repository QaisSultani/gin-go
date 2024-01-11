package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/QaisSultani/gin-go/logger"
	"github.com/QaisSultani/gin-go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Default() will attach logger and recovery handler to the router by default
	// where as below will not attach any default handler. logger is attached menually below.
	// router := gin.New()
	// router.Use(gin.Logger())
	// router.Use(middleware.Authenticate) // apply to all routes

	// custom specific format log
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // displaying on file and standard output

	router.Use(gin.LoggerWithFormatter(logger.FormatLogsJson))

	router.LoadHTMLGlob("templates/*")
	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
		// "user2": "pass2",
		// "user3": "pass3",
	})

	router.GET("/getData", middleware.Authenticate, middleware.AddHeader, getData)
	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)
	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData3", getData3)
	}
	client := router.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)
	}
	router.GET("/getUrlData/:name/:age", getUrlDataInHTML)
	router.POST("/getDataPost", getDataPost)

	// custom server configuration
	server := &http.Server{
		Addr:         ":3001",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
	// router.Run(":3001")
	// http.ListenAndServe(":3001", router) // same as above
}

func getData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hi I am GetData Method",
	})
}

func getData1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hi I am GetData1 Method",
	})
}

func getData2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hi I am GetData2 Method",
	})
}

func getData3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hi I am GetData3 Method",
	})
}

// http://localhost:3001/client/getQueryString?name=Qais&age=23
func getQueryString(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hi I am GET Method in GIN Framework with Query String",
		"name": name,
		"age":  age,
	})
}

// http://localhost:3001/getUrlData/Qais/30
func getUrlDataInHTML(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")

	// for HTML response
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"name": name,
		"age":  age,
	})

	// for JSON response
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": "Hi I am GET Method in GIN Framework with Query String",
	// 	"name": name,
	// 	"age":  age,
	// })
}

func getDataPost(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := io.ReadAll(body)

	ctx.JSON(http.StatusOK, gin.H{
		"data":     "Hi I am POST method from GIN Framework",
		"bodyData": string(value),
	})
}
