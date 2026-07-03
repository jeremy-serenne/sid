package sid

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOf_IsInitialized(t *testing.T) {
	t.Run("ShouldBeFalseWhenValueNotCalled", func(t *testing.T) {
		t.Parallel()

		// Given
		lazy := Of(func() int { return 42 })

		// When
		// Then
		require.False(t, lazy.IsInitialized())
	})
	t.Run("ShouldBeTrueWhenValueCalled", func(t *testing.T) {
		t.Parallel()

		// Given
		lazy := Of(func() int { return 42 })
		lazy.Value()

		// When
		// Then
		require.True(t, lazy.IsInitialized())
	})
}

func TestOf_Value(t *testing.T) {
	t.Run("ShouldReturnZeroValueWhenInitializerIsNil", func(t *testing.T) {
		t.Parallel()

		// Given
		lazy := Of[string](nil)

		// When
		// Then
		require.Zero(t, lazy.Value())
	})
	t.Run("ShouldBeTrueWhenValueCalled", func(t *testing.T) {
		t.Parallel()

		// Given
		lazy := Of(func() string { return "Bazinga!" })

		// When
		// Then
		require.Equal(t, "Bazinga!", lazy.Value())
	})
	t.Run("ShouldCallInitializerOnlyOnce", func(t *testing.T) {
		t.Parallel()

		// Given
		counter := 0
		lazy := Of(func() string {
			counter++
			return "Bazinga!"
		})

		// When
		lazy.Value()
		lazy.Value()
		lazy.Value()

		// Then
		require.Equal(t, 1, counter)
	})
	t.Run("ShouldCallInitializerOnlyOnceWithMultipleGoroutines", func(t *testing.T) {
		t.Parallel()

		// Given
		var counter int32
		wg := sync.WaitGroup{}

		// When
		lazy := Of(func() string {
			atomic.AddInt32(&counter, 1)
			return "Bazinga!"
		})

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				lazy.Value()
			}()
		}
		wg.Wait()

		// Then
		require.Equal(t, int32(1), counter)
	})
}
