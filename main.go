package main

type LRUCache struct {
	capacity int
	cache: map[string]*Node
}


type Node struct {

}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache: map[string]*Node,
	}
}

func (lrucCache LRUCache) Get(key string) interface{} {
	//to do

	return lrucCache.capacity + 1
}

func (lrucCache LRUCache) Set(key string) interface{} {
	return lrucCache.capacity + 1
}

func main() {
	// oi := NewLRUCache(5)

	// fmt.Println(oi.Get())

}
