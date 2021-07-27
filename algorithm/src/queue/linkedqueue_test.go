package queue

import (
	"strings"
	"testing"
)

func TestInitQueue_LinkedQueue(t *testing.T) {
	lq := LinkedQueue{}
	lq.InitQueue()
	t.Logf("%+v", lq)
}

func TestEnterQueue_LinkedQueue(t *testing.T) {
	lq := LinkedQueue{}
	lq.InitQueue()
	for i := 0; i < 5; i++ {
		lq.EnterQueue(i)
		lq.Print()
	}
}

func TestDeleteQueue_LinkedQueue(t *testing.T) {
	lq := LinkedQueue{}
	lq.InitQueue()
	for i := 0; i < 5; i++ {
		lq.EnterQueue(i)
	}
	for i := 0; i < 6; i++ {
		lq.DeleteQueue()
		lq.Print()
	}

	s := "abc"
	strings.Index(s, "abc")
}
