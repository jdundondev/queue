package queue

import (
	"testing"
	"sync/atomic"
)

func TestQueue (t *testing.T) {

	var count int32 = 0
	handler := func(val interface{}) {

		atomic.AddInt32(&count, int32(val.(int)))

	}
	q := NewQueue(handler, 5)

	for i := 0; i < 200; i++ {

		q.Push(i)
	}

	q.Wait()
	if count != 19900 {

		t.Fail()
	}

}

// In the original implementation, if a Push is followed by a Wait,
// it is possible that the WaitGroup has not yet been incremented
// before the Wait executes (in particular if a large amount of data
// is pushed). So the wait returns immediately rather than wait for
// the work to be done. This was fixed by moving the wg.Add(1) into
// the Push method. This test verifies this fix.
func TestQueue2(t *testing.T) {
	var count int32

	handler := func(val interface{}) {
		count = 123
	}

	q := NewQueue(handler, 1)
	bigPush := make([]int64, 3000000)
	q.Push(bigPush)
	q.Wait()

	if count != 123 {
		t.Fail()
	}
}
