package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

//
// Будьте внимательны - это пример использования некоторых примитивов синхронизации.
// Данный код лучше не копировать себе в проект, а разобраться, как он работает,
// что можно было бы улучшить и как бы такую же задачу вы решили сами.
//

type job struct {
	name      string
	executed  atomic.Bool
	executeAt atomic.Int64
	fn        func() error
}

type result struct {
	name    string
	outcome bool
	err     error
}

func worker(ch chan *job, res chan result, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range ch {
		if !j.executed.Load() && time.Now().UnixMilli() >= j.executeAt.Load() {
			err := j.fn()
			j.executed.Store(true)
			res <- result{
				name:    j.name,
				outcome: err == nil,
				err:     err,
			}
		}
	}
}

func generateJobs() map[string]*job {
	jobs := make(map[string]*job)

	for i := 0; i < 9; i++ {
		name := uuid.NewString()
		j := &job{
			name: name,
			fn: func() error {
				st := fmt.Sprintf("execution time: %d\n", time.Now().UnixMilli())

				fmt.Println(st)
				return nil
			},
		}

		j.executeAt.Store(time.Now().UnixMilli() + 1*time.Second.Milliseconds())

		jobs[name] = j
	}

	return jobs
}

func main() {
	ch := make(chan *job)
	resCh := make(chan result)

	wg := new(sync.WaitGroup)

	jobs := generateJobs()

	var m sync.Mutex

	count := 3
	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker(ch, resCh, wg)
	}

	go func() {
		for res := range resCh {
			go func() {
				m.Lock()
				delete(jobs, res.name)
				m.Unlock()
			}()
		}
	}()

	for {
		m.Lock()
		if len(jobs) == 0 {
			close(ch)
			break
		}
		m.Unlock()

		m.Lock()
		for _, j := range jobs {
			ch <- j
		}
		m.Unlock()
	}

	wg.Wait()

	close(resCh)
}
