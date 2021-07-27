package queue

import "testing"

func TestInitQueue_LoopQueue(t *testing.T) {
	var lq LoopQueue
	lq.InitQueue(5)
	t.Logf("%#v", lq.Data)
}

func TestEnterQueue_LoopQueue(t *testing.T) {
	var lq LoopQueue
	lq.InitQueue(5)
	for i := 1; i <= 6; i++ {
		lq.EnterQueue(i)
		t.Logf("%#v", lq.Data)
	}
}

func TestDeleteQueue_LoopQueue(t *testing.T) {
	var lq LoopQueue
	lq.InitQueue(5)
	for i := 1; i <= 6; i++ {
		lq.EnterQueue(i)
	}
	for i := 1; i <= 6; i++ {
		lq.DeleteQueue()
		t.Logf("%#v", lq.Data)
	}
}
