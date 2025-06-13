package chans

import "sync"

// Merge merges any number of channels to a new single non-nil, possibly empty channel.
// It preserves the relative order of elements that were in the same channel initially,
// but it gives no guarantees for the order of the elements from the different channels
func Merge[T any](channels ...<-chan T) <-chan T {
	if len(channels) == 0 {
		ch := make(chan T, 0)
		close(ch)
		return ch
	}
	if len(channels) == 1 {
		ch := make(chan T)
		go func() {
			defer close(ch)
			for elem := range channels[0] {
				ch <- elem
			}
		}()
		return ch
	}

	ret := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for elem := range ch {
				ret <- elem
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ret)
	}()
	return ret
}

func ProcessAndPipe[T any](channel <-chan T, processor func(T) T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for value := range channel {
			out <- processor(value)
		}
	}()

	return out
}
