package v1

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/setting"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

//GET api/v1/index
func Index(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	books, err := book.QueryIndex(util.GetPage(c), setting.AppCfg.PAGE_SIZE, false)
	row, _ := book.BooksRow()
	if err != nil {
		data = map[string]interface{}{"error": "server error"}
		msg = "server error"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data = map[string]interface{}{"success": "request success", "books": books, "row": row}
	msg = "success"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//GET api/v1/index/score
func IndexScore(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	books, err := book.QueryIndex(util.GetPage(c), setting.AppCfg.PAGE_SIZE, true)
	if err != nil {
		data = map[string]interface{}{"error": "server error"}
		msg = "server error"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data = map[string]interface{}{"success": "request success", "books": books}
	msg = "success"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}
