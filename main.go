package main

import (
	"github.com/gin-gonic/gin"
	"github.com/makki0205/remote_shutter/controller"
	"github.com/makki0205/remote_shutter/websocket"
	"strconv"
)

//type hoge struct {
//	Name string `json:"name"`
//	Msg  string `json:"foo"`
//}

// HTTP サーバー起動
func main() {
	r := gin.Default()
	soctr := controller.SocketController{}
	r.GET("/ws", soctr.SocketHandle)
	r.GET("/@:id", func(c *gin.Context) {
		idstr := c.Param("id")
		id,_ :=strconv.Atoi(idstr)
		conn := websocket.GetManager().Get(id)
		conn.Emit("shutter", "on!!")
	})
	r.Run("0.0.0.0:2020")
}
//
//// ルーターの作成とルーティング設定
//func createRouter() *golem.Router {
//	router := golem.NewRouter()
//	router.OnConnect(foo)
//	router.On("hoge", echo)
//	return router
//}
//
//// echo イベントでハンドリングされる関数
//// 受け取ったメッセージをそのまま返す
//func echo(conn *golem.Connection, data *echoMessage) {
//	fmt.Printf("%v", data.Msg)
//	conn.Emit("hoge", hoge{Name:"junya",Msg:"junya no aho"})
//}
//func foo(conn *golem.Connection, http *http.Request){
//	fmt.Printf(" %v\n", http.RemoteAddr)
//	conn.Emit("hoge","hello!!")
//
//}
//// メッセージクラス
//type echoMessage struct {
//	Msg string `json:"msg"`
//}