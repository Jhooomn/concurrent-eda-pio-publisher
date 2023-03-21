package task

import (
	"github.com/Jhooomn/concurrent-eda-pio/publisher/infrastructure/broker"
	"github.com/google/uuid"
	"sync"
	"time"
)

func Flood() {
	for {
		Do()
		time.Sleep(time.Duration(5 * time.Second))
	}
}

func Do() {
	go func() {
		var wg sync.WaitGroup
		l := 100
		wg.Add(l)
		for i := 0; i < l; i++ {
			go func() {
				defer wg.Done()
				broker.Send(uuid.NewString())
			}()
		}
		wg.Wait()
	}()
}
