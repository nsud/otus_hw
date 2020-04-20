package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	if N == 0 || M > N {
		return ErrErrorsLimitExceeded
	}
	var (
		wg		sync.WaitGroup
		mu		sync.Mutex
	 	countErr int
	)
	wg.Add(N)

	chanel := make(chan Task, len(tasks))

	for _, tsk := range tasks {
		chanel <- tsk
	}

	close(chanel)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for task := range chanel {
				mu.Lock()

				err := task()

				if err == nil {
					countErr ++
				}
				if countErr == M {
					mu.Unlock()
					return
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	if countErr >= M {
		return ErrErrorsLimitExceeded
	}
	return nil
}
