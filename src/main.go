package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/go-vgo/robotgo"
)

var commands = []string{
	"left",
	"up",
	"right",
	"down",
	"up_left",
	"up_right",
	"bottom_left",
	"bottom_right",
}

func issueCommand(cmd string) {
	switch cmd {
	case "left":
		robotgo.MoveRelative(-10, 0)
	case "up":
		robotgo.MoveRelative(0, -10)
	case "right":
		robotgo.MoveRelative(10, 0)
	case "down":
		robotgo.MoveRelative(0, 10)
	case "up_left":
		robotgo.MoveRelative(-10, -10)
	case "up_right":
		robotgo.MoveRelative(10, -10)
	case "bottom_left":
		robotgo.MoveRelative(-10, 10)
	case "bottom_right":
		robotgo.MoveRelative(-10, 10)
	case "click":
		robotgo.MouseClick("left", false)
	default:
		fmt.Println("no command found")
	}
}

func handler(_ http.ResponseWriter, r *http.Request) {
		issueCommand(r.URL.Path[1:])
}

func main() {
		fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    for _, cmd := range commands {
    	http.HandleFunc(fmt.Sprintf("/%s", cmd), handler)
    }

    log.Fatal(http.ListenAndServe(":8080", nil))
}
