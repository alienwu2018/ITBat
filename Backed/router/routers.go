package router

import (
	v1 "ITBat/controller/api/v1"
	"ITBat/controller/api/v1/page"
	"ITBat/controller/api/v1/score"
	"ITBat/controller/api/v1/search"
	"ITBat/middleware"
	"ITBat/pkg/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(setting.RUNMODE)
	r := gin.New()
	//全局中间件
	r.Use(gin.Logger())
	r.Use(middleware.LoggerToFile())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/index", v1.Index)
		apiv1.GET("/search", search.BookSearchCtr)
		book := apiv1.Group("/book/category")
		{
			book.GET("/:bigCategory", page.BigCategoryCtr)
			book.POST("/:bigCategory/:smallCategory", score.BookScore)
			book.GET("/:bigCategory/:smallCategory", page.SmallCategoryCtr)

		}

	}
	return r
}
