package task_2

import (
	"fmt"
	"sync"
)

type LogBuffer struct {
	buffer []string
	mute   sync.Mutex
}

func (lb *LogBuffer) LogBuffer(message string) {
	lb.mute.Lock()
	lb.buffer = append(lb.buffer, message)
	lb.mute.Unlock()
}

func StartTask2() {
	logs := LogBuffer{}
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			message := fmt.Sprintf("Message number: %d", num)
			logs.LogBuffer(message)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(logs.buffer)
	fmt.Println("Количество добавленых сообщений: ", len(logs.buffer))
}
