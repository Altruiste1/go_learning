package scheduler

import "github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/engine"

// 每个worker公用一个chan
type SimpleScheduler struct {
	workerChan chan engine.Request
}
//type Scheduler interface {
//	Submit(engine.Request)
//	//ConfigureMasterWorkerChan(chan Request)
//	WorkerChan()engine.Request
//	WorkerReady(chan engine.Request)
//	Run()
//}
func (s *SimpleScheduler)Submit(
	r engine.Request){
		go func() {
			s.workerChan <- r
		}()
	}

// 分配公用的channel
func (s *SimpleScheduler)WorkerChan()chan engine.Request{
	return s.workerChan
}

func (s *SimpleScheduler)WorkerReady(chan engine.Request){

}
func (s *SimpleScheduler)Run(){
	s.workerChan=make(chan engine.Request)
}


