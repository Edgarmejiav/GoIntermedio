package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkPool  chan chan Job
	MaxWorkes int
	JobQueue  chan Job
}

func NewWorker(id int, workPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("worker with id %d started \n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("worker with id %d finished with resul %d \n", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("worker with id %d stopped \n", w.Id)
			}

		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.QuitChan <- true
	}()
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:  jobQueue,
		MaxWorkes: maxWorkers,
		WorkPool:  worker,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQuee := <-d.WorkPool
				workerJobQuee <- job
			}()

		}
	}

}
func (d *Dispatcher) run() {
	for i := 0; i < d.MaxWorkes; i++ {
		worker := NewWorker(i, d.WorkPool)
		worker.start()

	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	delay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)

}
func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8079"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatecher := NewDispatcher(jobQueue, maxWorkers)
	dispatecher.run()

	http.HandleFunc("/fib", func(writer http.ResponseWriter, request *http.Request) {
		RequestHandler(writer, request, jobQueue)
	})
	log.Fatalln(http.ListenAndServe(port, nil))
}
