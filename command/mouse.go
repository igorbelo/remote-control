package command

import (
	"github.com/go-vgo/robotgo"
)

type Mouse struct {
	Command string
}

func (m Mouse) handleMove() {
	var x, y int

	switch m.Command {
	case "moveleft":
		x, y = -5, 0
	case "moveup":
		x, y = 0, -5
	case "moveright":
		x, y = 5, 0
	case "movedown":
		x, y = 0, 5
	}

	robotgo.MoveRelative(x, y)
}

func (m Mouse) handleClick() {
	robotgo.MouseClick("left", false)
}

func (m Mouse) Handle() {
	if m.Command[:4] == "move" {
		m.handleMove()
		return
	}

	m.handleClick()
}
