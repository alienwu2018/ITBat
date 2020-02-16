package category

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

// GET /api/v1/categories
func DoCategoryCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	books, err := book.QueryAllCategory()
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
	categories, trueName := util.DoCategory(books)
	data = map[string]interface{}{"categories": categories, "trueName": trueName}
	msg = "success"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
