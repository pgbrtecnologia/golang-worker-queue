package queue

type worker struct {
	job   IJob
	queue chan IJob
	quit  chan bool
}

func newWorker(queue chan IJob, quit chan bool) *worker {
	return &worker{
		queue: queue,
		quit:  quit}
}

func (w *worker) getJob() {
	select {
	case job := <-w.queue:
		if job != nil {
			w.job = job
			w.job.Run()
			w.job = nil
		}
	default:
	}

	return
}

func (w *worker) start() {
	for {
		select {
		case <-w.quit:
			return
		default:
			w.getJob()
		}
	}
}
