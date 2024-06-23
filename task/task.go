package task

/*
TASK = Thing a user wants to run on our cluster
TASK EVENT = Needed when a user wants to stop a task (state change)
*/

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	Name          string
	ID            uuid.UUID
	State         State
	Image         string
	Memory        int               // Amount of memory requested
	Disk          int               // Amount of disk requested
	ExposedPorts  nat.PortSet       // Exposed ports
	PortBindings  map[string]string // Port bindings
	RestartPolicy string            // What to do if the task stops or dies unexpectedly
	StartTime     time.Time
	EndTime       time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State     // Desired state for the task to transition into
	Task      Task      // The task whose state needs to be changed
	Timestamp time.Time // When the state change was requested
}
