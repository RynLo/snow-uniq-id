package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	sequenceMax int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	sequence  int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		sequence:  0,
	}, nil
}

func (w *Worker) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *Worker) GetID() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := w.getMilliSeconds()
	if w.timestamp == now {
		w.sequence = (w.sequence + 1) & sequenceMax
		if w.sequence == 0 {
			// 序号满了，要等下一毫秒
			for now <= w.timestamp {
				now = w.getMilliSeconds()
			}
			// 需更新timestamp, 否则可能出现重复
			w.timestamp = now
		}
	} else {
		w.sequence = 0
		w.timestamp = now
	}
	id := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.sequence)

	return id
}
func main() {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetID())
	}
}
