package manager

import "sync"

type GoStatus uint8
type GoFunction func()

const (
	Createing GoStatus = iota + 1
	Running
	Waiting
	Stop
)

// 协程状态树
type StateManagerNode struct {
	nextNode      []*StateManagerNode
	mutex         sync.Mutex
	childNodesLen int
	status        GoStatus
	exitChan      chan struct{}
	doFunc        GoFunction
}

func NewNode(doFunc GoFunction) *StateManagerNode {
	return &StateManagerNode{
		nextNode:      make([]*StateManagerNode, 0),
		childNodesLen: 0,
		mutex:         sync.Mutex{},
		status:        Createing,
		exitChan:      make(chan struct{}),
		doFunc:        doFunc,
	}
}

func (node *StateManagerNode) Do() {
	go func() {
		defer func() {
			node.status = Stop
		}()
		node.doFunc()
	}()
}

func (node *StateManagerNode) RegisterNode(newNode ...*StateManagerNode) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	node.nextNode = append(node.nextNode, newNode...)

	node.childNodesLen = len(node.nextNode)

}
