package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleShow( c *gin.Context )  {
	c.HTML(http.StatusOK, "publish.html", gin.H{
	})
}
