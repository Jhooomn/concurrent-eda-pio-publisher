package infrastructure

import (
	"github.com/Jhooomn/concurrent-eda-pio/publisher/infrastructure/server"
	"github.com/Jhooomn/concurrent-eda-pio/publisher/shared/task"
)

func SetUp() {
	go task.Flood()
	server.SetUp() // last line of code
}
