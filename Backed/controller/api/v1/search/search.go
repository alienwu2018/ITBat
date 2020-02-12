package search

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/setting"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

//GET  api/v1/search
func BookSearchCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bookName := c.Query("query")
	books, err := book.QueryBookByName(util.GetPage(c), setting.PAGE_SIZE, bookName)
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
}
