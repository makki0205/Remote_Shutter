package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/makki0205/remote_shutter/websocket"
)

type SocketController struct {}

func NewSocketController() SocketController{
	return SocketController{}
}

func (self *SocketController)SocketHandle(c *gin.Context){
	websocket.GetHandle()(c.Writer, c.Request)
}

