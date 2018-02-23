package queue

// IJob defines the interface that has to
// be implemented by all jobs
type IJob interface {
	Run()
}
