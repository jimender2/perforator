package ccache

type Configuration[T any] struct {
	maxSize        int64
	buckets        int
	itemsToPrune   int
	deleteBuffer   int
	promoteBuffer  int
	getsPerPromote int32
	tracking       bool
	onDelete       func(item *Item[T])
}

// Creates a configuration object with sensible defaults
// Use this as the start of the fluent configuration:
// e.g.: ccache.New(ccache.Configure().MaxSize(10000))
func Configure[T any]() *Configuration[T] {
	return &Configuration[T]{
		buckets:        16,
		itemsToPrune:   500,
		deleteBuffer:   1024,
		getsPerPromote: 3,
		promoteBuffer:  1024,
		maxSize:        5000,
		tracking:       false,
	}
}

// The max size for the cache
// [5000]
func (c *Configuration[T]) MaxSize(max int64) *Configuration[T] {
	c.maxSize = max
	return c
}

// Keys are hashed into % bucket count to provide greater concurrency (every set
// requires a write lock on the bucket). Must be a power of 2 (1, 2, 4, 8, 16, ...)
// [16]
func (c *Configuration[T]) Buckets(count uint32) *Configuration[T] {
	if count == 0 || ((count&(^count+1)) == count) == false {
		count = 16
	}
	c.buckets = int(count)
	return c
}

// The number of items to prune when memory is low
// [500]
func (c *Configuration[T]) ItemsToPrune(count uint32) *Configuration[T] {
	c.itemsToPrune = int(count)
	return c
}

// The size of the queue for items which should be promoted. If the queue fills
// up, promotions are skipped
// [1024]
func (c *Configuration[T]) PromoteBuffer(size uint32) *Configuration[T] {
	c.promoteBuffer = int(size)
	return c
}

// The size of the queue for items which should be deleted. If the queue fills
// up, calls to Delete() will block
func (c *Configuration[T]) DeleteBuffer(size uint32) *Configuration[T] {
	c.deleteBuffer = int(size)
	return c
}

// Give a large cache with a high read / write ratio, it's usually unnecessary
// to promote an item on every Get. GetsPerPromote specifies the number of Gets
// a key must have before being promoted
// [3]
func (c *Configuration[T]) GetsPerPromote(count int32) *Configuration[T] {
	c.getsPerPromote = count
	return c
}

// Typically, a cache is agnostic about how cached values are use. This is fine
// for a typical cache usage, where you fetch an item from the cache, do something
// (write it out) and nothing else.

// However, if callers are going to keep a reference to a cached item for a long
// time, things get messy. Specifically, the cache can evict the item, while
// references still exist. Technically, this isn't an issue. However, if you reload
// the item back into the cache, you end up with 2 objects representing the same
// data. This is a waste of space and could lead to weird behavior (the type an
// identity map is meant to solve).

// By turning tracking on and using the cache's TrackingGet, the cache
// won't evict items which you haven't called Release() on. It's a simple reference
// counter.
func (c *Configuration[T]) Track() *Configuration[T] {
	c.tracking = true
	return c
}

// OnDelete allows setting a callback function to react to ideam deletion.
// This typically allows to do a cleanup of resources, such as calling a Close() on
// cached object that require some kind of tear-down.
func (c *Configuration[T]) OnDelete(callback func(item *Item[T])) *Configuration[T] {
	c.onDelete = callback
	return c
}
