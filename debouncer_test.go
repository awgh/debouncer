package debouncer

import (
	"testing"
	"time"
)

func TestFires(t *testing.T) {
	fired := false
	duration := 10 * time.Millisecond
	debouncer := New(duration, func() {
		fired = true
	})
	debouncer.Trigger()
	time.Sleep(2 * duration)
	if !fired {
		t.Error("did not fire")
	}
}

func TestFiresTwice(t *testing.T) {
	count := 0
	duration := 10 * time.Millisecond
	debouncer := New(duration, func() {
		count++
	})

	debouncer.Trigger()
	time.Sleep(duration * 2)
	debouncer.Trigger()
	time.Sleep(duration * 2)
	if count != 2 {
		t.Errorf("didn't fire twice: %d", count)
	}
}

func TestFiresRateLimit(t *testing.T) {
	duration := 10 * time.Millisecond
	lastFired := time.Unix(0, 0)
	count := 0
	debouncer := New(duration, func() {
		period := time.Now().Sub(lastFired)
		if period < duration {
			t.Errorf("period less than duration; count=%d", count)
		}
		lastFired = time.Now()
		count++
	})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Microsecond)
		debouncer.Trigger()
	}
	time.Sleep(duration * 2)
	if count < 1 {
		t.Error("never fired")
	}
}
