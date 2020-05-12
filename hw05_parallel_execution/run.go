package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	if len(tasks) < 1 {
		return errors.New("the task list is empty")
	}

	if n == 0 || m > len(tasks) {
		return ErrErrorsLimitExceeded
	}

	wg := &sync.WaitGroup{}
	chanel := make(chan Task, len(tasks))
	var counter int32

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for task := range chanel {
				if atomic.LoadInt32(&counter) > int32(m) {
					continue
				}
				err := task()
				if err != nil {
					atomic.AddInt32(&counter, 1)
				}
			}
		}()
	}

	for _, tsk := range tasks {
		chanel <- tsk
	}

	close(chanel)
	wg.Wait()

	if atomic.LoadInt32(&counter) > int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
