package node

/*
NODE = Any machine in the cluster. This means a node
can be a control plane (i.e. manager) or worker.
*/

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Role            string
	TaskCount       int
}
