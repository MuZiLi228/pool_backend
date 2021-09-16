package ws

import (
	"errors"
	"fmt"
	"pool_backend/src/global"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	data []byte
	err  error
)

//Connection 客户端连接
type Connection struct {
	wsConnect *websocket.Conn // 底层websocket
	inChan    chan *WsMessage // 读队列
	outChan   chan *WsMessage // 写队列
	closeChan chan byte       // 关闭通知
	mutex     sync.Mutex      // 避免重复关闭管道
	isClosed  bool            // 防止closeChan被关闭多次
}

//WsMessage 客户端读写消息
type WsMessage struct {
	MessageType int
	Data        []byte
}

//InitConnection 初始化
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan *WsMessage, 1000),
		outChan:   make(chan *WsMessage, 1000),

		closeChan: make(chan byte, 1),
	}
	// 处理器
	// go conn.ProcLoop()
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

//ReadMessage 读取websocket消息
func (conn *Connection) ReadMessage() (msg *WsMessage, err error) {

	select {
	case msg = <-conn.inChan:
		return msg, nil
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

//WriteMessage 发送消息到websocket
func (conn *Connection) WriteMessage(messageType int, data []byte) (err error) {

	select {
	case conn.outChan <- &WsMessage{messageType, data}:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

//Close 关闭连接
func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	for {
		messageType, data, err := conn.wsConnect.ReadMessage()
		if err != nil {
			goto ERR
		}
		req := &WsMessage{
			MessageType: messageType,
			Data:        data,
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- req:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}

	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-conn.outChan:
			if err := conn.wsConnect.WriteMessage(msg.MessageType, msg.Data); err != nil {
				goto ERR
			}
		case <-conn.closeChan:
			goto ERR
		}

	}

ERR:
	conn.Close()

}

//ProcLoop 发送心跳
func (conn *Connection) ProcLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := conn.WriteMessage(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				conn.Close()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		err = conn.WriteMessage(msg.MessageType, msg.Data)
		if err != nil {
			fmt.Println("write fail")
			break
		}
	}
}

//Send 发送消息给客户端  data: `{"name":"block","id":"12346"}`
func Send(conn *Connection, msg *WsMessage) {
	// 启动线程，不断发消息
	sendData := []byte(msg.Data)
	go func() {
		for {
			if err = conn.WriteMessage(msg.MessageType, sendData); err != nil {
				global.Logger.Error("ws写数据 发生错误:", err)
				return
			}
			//休眠一秒
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if msg, err = conn.ReadMessage(); err != nil {
			global.Logger.Error("ws读数据发生错误:", err)
			goto ERR
		}
		if err = conn.WriteMessage(msg.MessageType, msg.Data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
