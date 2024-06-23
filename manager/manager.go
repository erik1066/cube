package manager

import (
	"cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue // tasks will be placed here on being submitted. Allows the Manager to handle tasks per FIFO
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
	Workers       []string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Select appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send out work")
}
