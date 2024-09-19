package gomit_test

import (
	"testing"
	"time"

	"github.com/mateusfdl/gomit"
)

func TestListenersAdded(t *testing.T) {
	firstCallback := func(s ...interface{}) error {
		return nil
	}
	secondCallback := func(s ...interface{}) error {
		return nil
	}

	gomit.AddListener("test", firstCallback)
	gomit.AddListener("test", secondCallback)

	listeners := gomit.Listeners()
	if len(listeners["test"]) != 2 {
		t.Errorf("Expected 2 listeners, got %d", len(gomit.Listeners()))
	}
}

func TestEmit(t *testing.T) {
	var calls int
	firstCallback := func(s ...interface{}) error {
		calls += 1
		return nil
	}
	secondCallback := func(s ...interface{}) error {
		calls += 1
		return nil
	}

	gomit.AddListener("test", firstCallback)
	gomit.AddListener("test", secondCallback)

	gomit.Emit("test")

	// Wait for 200ms since its async
	<-time.After(time.Millisecond * 200)
	if calls != 2 {
		t.Errorf("Expected calls to be 2, got %d", calls)
	}
}
