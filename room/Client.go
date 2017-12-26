package room

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:12345", "http service address")

type Chat struct {
}

func (c *Chat) ChatRoomStart() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	var dialer *websocket.Dialer

	// 通过聊天室地址(u.String())创建一个聊天室连接
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// go c.timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
	}
}

func (c *Chat) timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 60)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}

func (c *Chat) printTime(manager *ClientManager) {
	// 当有人说话时，如果离上一次说话时间超过五分钟，则打印一次当前时间
	now := time.Now().Unix()
	nowUnix := time.Now()
	if (now - manager.chatTime) > 300 {
		time := []byte(nowUnix.Format("15:04:05"))
		for conn := range manager.clients {
			conn.send <- time
		}
	}
}
