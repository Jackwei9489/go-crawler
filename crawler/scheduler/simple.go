package scheduler

import "go-crawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker channel
	go func() {
		s.workerChan <- request
	}()
}
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
