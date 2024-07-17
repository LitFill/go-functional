package itx

import (
	"iter"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
)

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in the `it` package.
	Iterator[V any] iter.Seq[V]

	// Iterator2 is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in the `it` package.
	Iterator2[V, W any] iter.Seq2[V, W]
)

// From converts an iterator in an [Iterator] to support method chaining.
func From[V any](iterator func(func(V) bool)) Iterator[V] {
	return Iterator[V](iterator)
}

// From2 converts an iterator in an [Iterator2] to support method chaining.
func From2[V, W any](iterator func(func(V, W) bool)) Iterator2[V, W] {
	return Iterator2[V, W](iterator)
}

// Seq converts an [Iterator] to an [iter.Seq].
func (iterator Iterator[V]) Seq() iter.Seq[V] {
	return iter.Seq[V](iterator)
}

// Seq converts an [Iterator2] to an [iter.Seq2].
func (iterator Iterator2[V, W]) Seq() iter.Seq2[V, W] {
	return iter.Seq2[V, W](iterator)
}

// Collect is a convenience method for chaining [slices.Collect] on
// [Iterator]s.
func (iterator Iterator[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](iterator))
}

// ForEach is a convenience method for chaining [it.ForEach] on [Iterator]s.
func (iterator Iterator[V]) ForEach(fn func(V)) {
	it.ForEach(iterator, fn)
}

// ForEach is a convenience method for chaining [it.ForEach2] on [Iterator2]s.
func (iterator Iterator2[V, W]) ForEach(fn func(V, W)) {
	it.ForEach2(iterator, fn)
}

// Find is a convenience method for chaining [it.Find] on [Iterator]s.
func (iterator Iterator[V]) Find(predicate func(V) bool) (V, bool) {
	return it.Find(iterator, predicate)
}

// Find is a convenience method for chaining [it.Find2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Find(predicate func(V, W) bool) (V, W, bool) {
	return it.Find2(iterator, predicate)
}