package main

import (
	"github.com/chuccp/api/vmess"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/d3MuY2oyMDIw.md", vmess.Api2)
	//r.GET("/put", vmess.Api2)
	r.Run(":8082")
}
