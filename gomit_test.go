package gomit_test

import (
	"testing"

	"github.com/mateusfdl/gomit"
)

func TestListenersAdded(t *testing.T) {
	type Foo struct{}
	type Bar struct{}

	gomit.AddListener(func(s Foo) error { return nil })
	gomit.AddListener(func(s Bar) error { return nil })

	if gomit.ActiveListeners() != 2 {
		t.Errorf("Expected 2 listeners, got %d", len(gomit.Listeners()))
	}
}

func TestEmit(t *testing.T) {
	var calls int
	type FirstEvent struct{}
	type SecondEvent struct{}
	type AssertEvent struct{ Calls *int }
	gomit.AddListener(func(s FirstEvent) error {
		calls += 1
		return nil
	})
	gomit.AddListener(func(s SecondEvent) error {
		calls += 1
		return nil
	})

	gomit.Emit(FirstEvent{})
	gomit.Emit(SecondEvent{})
	gomit.AddListener(func(s AssertEvent) error {
		if *s.Calls != 2 {
			t.Errorf("Expected 2 calls, got %d", calls)
		}
		return nil
	})

	gomit.Emit(AssertEvent{
		Calls: &calls,
	})
}
