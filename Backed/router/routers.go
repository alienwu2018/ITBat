package router

import (
	v1 "ITBat/controller/api/v1"
	"ITBat/controller/api/v1/book"
	"ITBat/controller/api/v1/download"
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
	gin.SetMode(setting.ServerCfg.RUNMODE)
	r := gin.New()
	//全局中间件
	//r.Use(gin.Logger())
	r.Use(middleware.LoggerToFile())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/index", v1.Index)
		apiv1.GET("/search", search.BookSearchCtr)
		b := apiv1.Group("/book")
		{
			b.GET("/:bid", book.BooksCtr)
			b.GET("/:bid/download/:bid", download.DownBookCtr)
		}
		bs := apiv1.Group("/books")
		{
			bs.GET("/category/:bigCategory", page.BigCategoryCtr)
			bs.POST("/category/:bigCategory/:smallCategory", score.BookScoreCtr)
			bs.GET("/category/:bigCategory/:smallCategory", page.SmallCategoryCtr)
		}
	}
	return r
}
