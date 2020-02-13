package score

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/setting"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

// POST api/v1/book/category/:bigCategory/:smallCategory
func BookScoreCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bigCategory := c.Param("bigCategory")
	smallCategory := c.Param("smallCategory")
	flag := 0
	for _, i := range e.BingCategory {
		if bigCategory == i {
			flag += 1
		}
	}
	for _, i := range e.SmallCategory {
		if smallCategory == i {
			flag += 1
		}
	}

	if flag != 2 {
		code = e.BAD_REQUEST
		data = map[string]interface{}{"error": "bad request"}
		msg = "bad requests"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	} else {
		books, err := book.QueryBookByScore(util.GetPage(c), setting.PAGE_SIZE, bigCategory, smallCategory)
		row, _ := book.SmallCategoryRow(bigCategory, smallCategory)
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
}
