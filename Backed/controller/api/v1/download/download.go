package download

import (
	"ITBat/model/book"
	"ITBat/model/downlink"
	"ITBat/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// api/v1/book/:bid/download:bid
func DownBookCtr(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	bid := com.StrTo(c.Param("bid")).MustInt()
	rows, _ := book.BooksRow()
	if bid > rows {
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
	download, err := downlink.QueryByBId(bid)
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
	data = map[string]interface{}{"success": "request success", "downlink": download}
	msg = "success"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
