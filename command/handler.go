package command

type Handler interface {
	Handle()
}

var HANDLERS = map[string]Handler{
	"mousemoveleft": Mouse{Command: "moveleft"},
	"mousemoveup": Mouse{Command: "moveup"},
	"mousemoveright": Mouse{Command: "moveright"},
	"mousemovedown": Mouse{Command: "movedown"},
	"click": Mouse{Command: "click"},
	"volume+": Volume{Command: "+"},
	"volume-": Volume{Command: "-"},
	"volumemute": Volume{Command: "mute"},
	"volumeunmute": Volume{Command: "unmute"},
}

func Handle(cmd string) {
	handler := HANDLERS[cmd]
	handler.Handle()
}
