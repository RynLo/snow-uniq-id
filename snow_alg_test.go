package main

import (
	"github.com/imroc/biu"
	"sync"
	"testing"
)

func Test_Bit(t *testing.T) {

	s1 := biu.ToBinaryString(-1)
	s2 := biu.ToBinaryString(-32)
	t.Log(s1)
	t.Log(s2)
}

func TestWorker_GetID(t *testing.T) {
	w := &Worker{
		mu:        sync.Mutex{},
		timestamp: 0,
		workerId:  1,
		sequence:  0,
	}
	id1 := w.GetID()
	t.Log(id1)
	w.sequence = sequenceMax
	id2 := w.GetID()
	t.Log(id2)
	id3 := w.GetID()
	t.Log(id3)
}