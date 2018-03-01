package queue

type worker struct {
	job   IJob
	queue chan IJob
}

func newWorker(queue chan IJob) *worker {
	return &worker{
		queue: queue}
}

func (w *worker) start() {
	for {
		job := <-w.queue
		if job == nil {
			return
		}

		w.job = job
		w.job.Run()
		w.job = nil
	}
}
