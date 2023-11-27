package atomicSequence_test

import (
	"github.com/rdforte/sequencer/internal/atomicSequence"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestAtomicSequence(t *testing.T) {
	t.Run("it should return all numbers in order when increment sequencer concurrently", func(t *testing.T) {

		sequencer := atomicSequence.CreateAtomicSequencer()

		numSequences := 1000
		var wg sync.WaitGroup
		wg.Add(numSequences)

		for i := 0; i < numSequences; i++ {
			go func() {
				defer wg.Done()
				sequencer.Increment()
			}()
		}

		wg.Wait()

		wantFinalValue := int64(numSequences + 1)
		gotFinalValue := sequencer.Increment()

		assert.Equal(t, wantFinalValue, gotFinalValue)
	})
}
