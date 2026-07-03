package sid

import (
	"sync"
)

type Lazy[T any] struct {
	initializer func() T
	once        sync.Once

	isInitialized bool
	value         T
}

func (o *Lazy[T]) Value() T {
	if o.initializer != nil {
		o.once.Do(func() {
			o.value = o.initializer()
			o.isInitialized = true
			o.initializer = nil // Release it for the GC to collect it
		})
	}
	return o.value
}

func (o *Lazy[T]) IsInitialized() bool {
	return o.isInitialized
}

func Of[T any](initializer func() T) *Lazy[T] {
	return &Lazy[T]{initializer: initializer}
}
