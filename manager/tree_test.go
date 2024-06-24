package manager

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var count = atomic.Int64{}

func TestStateManagerNode(t *testing.T) {

	root := NewNode(nil, start)

	for i := 0; i < 10; i++ {
		node := NewNode(root, print)
		root.RegisterNode(node)
		for i := 0; i < 10000; i++ {
			nodeTwo := NewNode(node, print)
			node.RegisterNode(nodeTwo)
		}
	}

	root.Do()

	fmt.Println(count.Load())

	t.Log(count.Load())

}

func print() {
	count.Add(1)
}

func start() {
	fmt.Println("start")
}
