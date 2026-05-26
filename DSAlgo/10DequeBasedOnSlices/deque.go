package main

// Deque (double-ended queue) implementation using slices in Go
// Credits: https://medium.com/@letsgoprogramming/simple-yet-effective-generic-deque-in-go-a4588e62c0b9

type Deque[T any] struct {
	items []T	// slice to store the elements
}

func (d *Deque[T]) Len() int {
	return len(d.items)
}

func (d *Deque[T]) Front() (T, bool) {
	if len(d.items) == 0 {
		var noop T	// zero-value of T
		return noop, false
	}

	return d.items[0], true
}

func (d *Deque[T]) Back() (T, bool) {
	if len(d.items) == 0 {
		var noop T	// zero-value of T
		return noop, false
	}

	return d.items[len(d.items)-1], true
}

func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var noop T	// zero-value of T
		return noop, false
	}

	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]

	return item, true
}

func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var noop T	// Zero-value of T
		return noop, false
	}

	item := d.items[0]
	d.items = d.items[1:]

	return item, true
}