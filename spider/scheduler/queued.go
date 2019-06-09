package scheduler

import "spider/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//把request送到requestChan通道
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//把request的通道送到chan
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request) //通道的通道
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan: //submit往里送数据，这里取数据
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: //WorkerReady往里送数据，这里取数据
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //request 的chan加到active
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]

			}
		}
	}()
}
