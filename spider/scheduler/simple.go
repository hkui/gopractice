package scheduler

import "spider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan=make(chan engine.Request)
}


//把request送给workerchan
func (s *SimpleScheduler) Submit( r engine.Request) {
	go func() {
		s.workerChan<-r
	}()

}

