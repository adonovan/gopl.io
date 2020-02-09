package memo

import (
	"fmt"
	"sync"
	"testing"
)

func TestCancel(t *testing.T) {
	cancelled := fmt.Errorf("cancelled")

	finish := make(chan string)
	f := func(key string, done <-chan struct{}) (interface{}, error) {
		if done == nil {
			res := <-finish
			return res, nil
		} else {
			<-done
			return nil, cancelled
		}
	}

	m := New(Func(f))
	key := "key"

	done := make(chan struct{})
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		v, err := m.Get(key, done)
		wg1.Done()
		if v != nil || err != cancelled {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, cancelled)
		}
	}()
	close(done)

	wg1.Wait()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		v, err := m.Get(key, nil)
		s := v.(string)
		if s != "ok" || err != nil {
			t.Errorf("got %v, %v; want %v, %v", s, err, "ok", nil)
		}
		wg2.Done()
	}()
	finish <- "ok"
	wg2.Wait()
}
