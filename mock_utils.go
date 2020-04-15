package mtp

import (
	"fmt"
	"math/rand"
	"time"
)

// Waiter is an interface that can be implemented by any struct
// that sleeps during a controlled duration
// useful for mock implementation of calls to external services that might have network latency for instance
// Examples of implementation: wait a constant amount of time, randomly wait for a long time
type Waiter interface {
	Wait()
}

// WaitFixed waits for some constant time
type WaitFixed struct {
	WaitDuration time.Duration
}

// Wait waits for some time
func (w WaitFixed) Wait() {
	time.Sleep(w.WaitDuration)
}

// Failer is an interface that can be implemented by any struct
// that returns an error at a controlled rate
// useful for mock implementations to test for behavior when errors happen
// Examples of implementation: fail every 5 times, fail randomly 10% of the time, never fail
type Failer interface {
	Fails() error
}

// FailRandom randomly fails with a fixed FailRate
type FailRandom struct {
	FailRate float32
}

// Fails returns an error if it randomly failed
func (f FailRandom) Fails() error {
	if rand.Float32() < f.FailRate {
		return fmt.Errorf("Randomly failed")
	}
	return nil
}
