package internal

import (
	"errors"
	"github.com/busy-cloud/boat/mqtt"
	"sync"
	"sync/atomic"
	"time"
)

var increment atomic.Uint64

type Command map[string]any

//var commands lib.Map[chan Command]

var commands sync.Map

func Execute(noob string, cmd Command) {
	mqtt.Publish("noob/"+noob+"/command", cmd)
}

func Request(noob string, cmd Command) (Command, error) {
	id := increment.Add(1)
	cmd["id"] = id
	mqtt.Publish("noob/"+noob+"/command", cmd)

	ch := make(chan Command)
	commands.Store(id, ch)

	//等待结果
	select {
	case c := <-ch:
		return c, nil
	case <-time.After(time.Second):
		commands.Delete(id)
		return nil, errors.New("timeout")
	}
}

func HandleResponse(id string, cmd Command) {
	v, has := commands.LoadAndDelete(id)
	if has {
		ch := v.(chan Command)
		ch <- cmd
	}
}
