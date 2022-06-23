package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	if len(tasks) == 0 || n < 1 {
		return nil
	}
	if len(tasks) < n {
		n = len(tasks)
	}
	workChan := make(chan Task)
	var errorsCount, maxErrorsCount int32 = 0, int32(m)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range workChan {
				err := task()
				if err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) >= maxErrorsCount {
			break
		}
		workChan <- task
	}
	close(workChan)
	wg.Wait()
	if errorsCount >= maxErrorsCount {
		return ErrErrorsLimitExceeded
	}
	return nil
}
