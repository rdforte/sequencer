package atomicSequence

import (
	"sync/atomic"
)

func CreateAtomicSequencer() Sequencer {
	return &sequencer{0}
}

type Sequencer interface {
	Increment() int64
}

type sequencer struct {
	count int64
}

func (s *sequencer) Increment() int64 {
	return atomic.AddInt64(&s.count, 1)
}
