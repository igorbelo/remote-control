package server

import (
		"fmt"
    "log"
    "net/http"

		"github.com/gorilla/websocket"
		"github.com/igorbelo/remote-control/command"
)

func handler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		command.Handle(string(message))
		err = c.WriteMessage(mt, message)
	}
}

func Serve(addr string) {
	log.Println(fmt.Sprintf("started %s", addr))
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/cmd", handler)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(addr, nil))
}
