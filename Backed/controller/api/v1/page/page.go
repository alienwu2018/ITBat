package page

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"ITBat/pkg/setting"
	"ITBat/pkg/util"
	"github.com/gin-gonic/gin"
)

//  大类别接口
//GET  api/v1/book/category/:bigCategory
func BigCategoryCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bigCategory := c.Param("bigCategory")
	flag := 0
	for _, i := range e.BingCategory {
		if bigCategory == i {
			flag = 1
		}
	}
	if flag != 1 {
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
		books, err := book.QueryBookByBigCategory(util.GetPage(c), setting.AppCfg.PAGE_SIZE, bigCategory)
		row, _ := book.BigCategoryRow(bigCategory)
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

//小类别接口
//GET api/v1/book/category/:bigCategory/:smallCategory
func SmallCategoryCtr(c *gin.Context) {
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
		books, err := book.QueryBookBySmallCategory(util.GetPage(c), setting.AppCfg.PAGE_SIZE, bigCategory, smallCategory)
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
