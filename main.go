package main

import (
	"chatroom/room"
	"fmt"
	"net/http"
)

func main() {
	go new(room.Server).ServerStart()
	fmt.Println("start server")
	go new(room.Chat).ChatRoomStart()
	fmt.Println("start client")
	http.HandleFunc("/ws", new(room.Server).WsPage)
	http.ListenAndServe(":12345", nil)
}
