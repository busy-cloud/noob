package internal

import (
	"github.com/busy-cloud/boat/boot"
)

func init() {
	boot.Register("noob", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"log", "mqtt", "database"},
	})
}

func Startup() error {

	//订阅通知
	subscribe()

	return nil
}

func Shutdown() error {

	return nil
}
