package worker

/*
What is a worker and what does it do?
- Run "Tasks" as OCI containers
- Accept tasks to run from the Manager
- Give stats to the Manager so Manager can better schedule
- Keep track of the Worker's assigned tasks
*/

import (
	"cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Db        map[uuid.UUID]*task.Task
	TaskCount int
	Queue     queue.Queue
}

func (w *Worker) RunTask( /*task *task.Task*/ ) {
	//w.Db[task.ID] = task
	fmt.Println("Run task")
}

func (w *Worker) CollectStats() {
	fmt.Println("Collect statistics")
}

func (w *Worker) StartTask() {
	fmt.Println("Start task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stop task")
}
