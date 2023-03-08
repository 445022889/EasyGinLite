package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"body": nil})
}
