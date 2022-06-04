package main

import (
	"github.com/caddyserver/caddy"
	"github.com/chuccp/api/vmess"
	"github.com/gin-gonic/gin"
)

func main() {

	caddy.Start()

	caddy.Start()

	r := gin.Default()

	r.GET("/d3MuY2oyMDIw.md", vmess.Api)
	r.Run(":8082")
}
