package task_3

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type WebCounter struct {
	visits sync.Map
}

func (wc *WebCounter) AddWebSite(site string) {
	wc.visits.Store(site, int64(0))
}

func (wc *WebCounter) Increment(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	if countVisit, ok := wc.visits.Load(url); ok {
		if count, ok := countVisit.(int64); ok {
			atomic.AddInt64(&count, 1)
			wc.visits.Store(url, count)
		} else {
			fmt.Println("Unexpected type for countVisit")
		}
	}
}

func (wc *WebCounter) GetVisits(url string) int {
	if countVisit, ok := wc.visits.Load(url); ok {
		return int(countVisit.(int64))
	}
	fmt.Println("Site not found")
	return 0
}

func StartTask3() {
	var wg sync.WaitGroup
	web := WebCounter{}
	web.AddWebSite("yandex.ru")
	web.AddWebSite("drive2.ru")
	web.AddWebSite("google.com")

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go web.Increment("yandex.ru", &wg)
	}

	for i := 1; i < 50; i++ {
		wg.Add(1)
		go web.Increment("drive2.ru", &wg)
	}

	for i := 1; i < 82; i++ {
		wg.Add(1)
		go web.Increment("google.com", &wg)
	}

	wg.Wait()

	web.visits.Range(func(site, countVisit interface{}) bool {
		fmt.Println(site, countVisit)
		return true
	})
}
