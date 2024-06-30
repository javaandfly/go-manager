package manager

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count = atomic.Int64{}

func TestStateManagerNode(t *testing.T) {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	now := time.Now()

	for i := 0; i < 100; i++ {
		root := NewNode(nil, nil)

		for i := 0; i < 2; i++ {
			node := NewNode(root, nil)
			root.RegisterNode(node)
			for i := 0; i < 4; i++ {
				nodeTwo := NewNode(node, print)
				node.RegisterNode(nodeTwo)
			}
		}

		root.Do()
	}

	fmt.Println("执行协程个数为 ", count.Load(), "消耗时间为", time.Since(now))

	t.Log(count.Load())

}

func TestGroupWait(t *testing.T) {
	f, err := os.Create("cpu-group.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()

	now := time.Now()
	for i := 0; i < 100; i++ {

		wg := sync.WaitGroup{}
		for i := 0; i < 2; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				wg1 := sync.WaitGroup{}
				for i := 0; i < 4; i++ {
					wg1.Add(1)
					go func() {
						defer wg1.Done()
						print()
					}()
				}
				wg1.Wait()
			}()
		}
		wg.Wait()
	}

	fmt.Println("执行协程个数为 ", count.Load(), "消耗时间为", time.Since(now))

	t.Log(count.Load())
}

func print() {
	for i := 0; i < 10000000; i++ {

	}
	count.Add(1)
}

func start() {
	fmt.Println("start")
}
