package ws

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type SocketHandler struct{}

//RegisterSocket 注册sokcet路由
func (sc *SocketHandler) RegisterSocket(router *gin.Engine, server *socketio.Server) {

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	// server.OnConnect("/", sc.onConnect)
	// server.OnError("/", sc.onError)
	// server.OnDisconnect("/", sc.onDisconnect)

	// server.OnEvent("/", "SINGLE_CHAT", sc.onSingleChat)
	// server.OnEvent("/", "BYE", sc.onBye)

}

// func (sc *SocketHandler) onConnect(s socketio.Conn) error {
// 	s.SetContext("")
// 	s.Emit("CONNECT", "欢迎连接 ~ ")
// 	return nil
// }

// func (sc *SocketHandler) onError(s socketio.Conn, e error) {
// 	fmt.Println("meet error:", e)
// }

// func (sc *SocketHandler) onDisconnect(s socketio.Conn, reason string) {
// 	fmt.Println("closed", reason)
// }

// func (sc *SocketHandler) onSingleChat(s socketio.Conn, message map[string]string) {
// 	fmt.Printf("%v \n", message)
// }

// func (sc *SocketHandler) onBye(s socketio.Conn) string {
// 	last := s.Context().(string)
// 	s.Emit("BYE", last)
// 	s.Close()
// 	return last
// }
