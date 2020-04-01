package command

import (
	"github.com/itchyny/volume-go"
	"log"
)

type Volume struct {
	Command string
}

func (v Volume) Handle() {
	var err error

	switch v.Command {
	case "+":
		err = volume.IncreaseVolume(1)
	case "-":
		err = volume.IncreaseVolume(-1)
	case "mute":
		err = volume.Mute()
	case "unmute":
		err = volume.Unmute()
	}

	if err != nil {
		log.Println(err)
	}
}
