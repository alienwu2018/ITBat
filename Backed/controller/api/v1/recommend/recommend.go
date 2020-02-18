package recommend

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//Get api/v1/book/recommend/:bid
func RecommendsCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	id := c.Param("bid")
	bid, err := com.StrTo(id).Int()
	if err != nil {
		code = e.BAD_REQUEST
		data = map[string]interface{}{"err": "bad request"}
		msg = "请求参数有误"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	books, err := book.RecommendBooks(bid)
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
	var res []struct {
		ID       int64
		BookName string
		PngUrl   string
	}
	for _, i := range books {
		tmp := struct {
			ID       int64
			BookName string
			PngUrl   string
		}{i.ID, i.BookName, i.PngUrl}
		res = append(res, tmp)
	}

	code = e.OK
	data = map[string]interface{}{"success": "request success", "recommend": res}
	msg = "request success"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
