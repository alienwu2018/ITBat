package search

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/setting"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

//GET  api/v1/search
func BookSearchV1Ctr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bookName := c.Query("query")
	if len(bookName) > 0 {
		books, err := book.QueryBookByName(util.GetPage(c), setting.AppCfg.PAGE_SIZE, bookName, false)
		row, _ := book.BooksNameRow(bookName)
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
	} else {
		code = e.BAD_REQUEST
		data = map[string]interface{}{"error": "bad request"}
		msg = "bad requests"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
}

//PUT api/v1/search
func BookSearchV2Ctr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bookName := c.Query("query")
	if len(bookName) > 0 {
		books, err := book.QueryBookByName(util.GetPage(c), setting.AppCfg.PAGE_SIZE, bookName, true)
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
	} else {
		code = e.BAD_REQUEST
		data = map[string]interface{}{"error": "bad request"}
		msg = "bad requests"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
}
